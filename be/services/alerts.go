package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/echoturing/log"
	"github.com/labstack/echo/v4"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/ent/schema"
)

func (i *impl) CreateAlert(ctx context.Context, alert *ent.Alert) (*ent.Alert, error) {
	// TODO:validate data field
	alert, err := i.dal.CreateAlert(ctx, alert)
	if err != nil {
		return nil, err
	}
	err = i.StartAlert(ctx, alert)
	if err != nil {
		log.ErrorWithContext(ctx, "start alert error", "err", err.Error())
		// just log an error,do not return
	}
	return alert, nil
}

func (i *impl) ListAlerts(ctx context.Context, status schema.AlertStatus, alertStatus schema.AlertState) ([]*ent.Alert, error) {
	return i.dal.ListAlerts(ctx, status, alertStatus)
}

type UpdateAlertRequest struct {
	Name     string             `json:"name"`
	Channels []int64            `json:"channels"`
	Rule     *schema.Rule       `json:"rule"`
	Status   schema.AlertStatus `json:"status"`
}

func (i *impl) UpdateAlert(ctx context.Context, id int64, update *UpdateAlertRequest) (*ent.Alert, error) {
	count, err := i.dal.UpdateAlert(ctx, id, &ent.Alert{
		Name:     update.Name,
		Channels: update.Channels,
		Rule:     *update.Rule,
		Status:   update.Status,
	})
	if err != nil {
		return nil, err
	}
	alert, err := i.dal.GetAlertByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get alert by id err:%w", err)
	}
	switch update.Status {
	case schema.StatusOpen:
		err = i.StartAlert(ctx, alert)
		if err != nil {
			return nil, fmt.Errorf("start alert error:%w", err)
		}
	case schema.StatusClose:
		err = i.stopAlert(ctx, alert)
		if err != nil {
			return nil, fmt.Errorf("stop alert error:%w", err)
		}
	}
	return count, nil
}

func (i *impl) StartAllAlert(ctx context.Context) error {
	// get all online alerts
	// start one
	als, err := i.dal.ListAlerts(ctx, schema.StatusOpen, 0)
	if err != nil {
		return err
	}
	for _, alert := range als {
		err := i.StartAlert(ctx, alert)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *impl) evaluatesRule(ctx context.Context, rule schema.Rule) (*schema.RuleResult, error) {
	var final schema.RuleResult
	for _, condition := range rule.Conditions {
		conditionResults, err := i.evaluatesCondition(ctx, condition.Condition)
		if err != nil {
			return nil, err
		}
		ruleResult := mergeConditionResultToRuleResult(ctx, conditionResults)
		switch condition.Type {
		default:
			log.ErrorWithContext(ctx, "unknown condition type", "condition", condition)
		case schema.ConditionRelationTypeOr, schema.ConditionRelationTypeUndefined:
			final.Qualified = final.Qualified || ruleResult.Qualified
		case schema.ConditionRelationTypeAnd:
			final.Qualified = final.Qualified && ruleResult.Qualified
		}
		final.Detail = append(final.Detail, ruleResult.Detail...)
	}
	return &final, nil
}

func (i *impl) evaluatesCondition(ctx context.Context, condition *schema.Condition) ([]*schema.ConditionResult, error) {
	// get datasource
	datasource, err := i.dal.GetDatasourceByID(ctx, condition.DatasourceID)
	if err != nil {
		return nil, err
	}
	datasourceResults, err := i.evaluatesDatasource(ctx, datasource, condition.Script)
	if err != nil {
		return nil, err
	}

	results := make([]*schema.ConditionResult, 0, len(datasourceResults))
	for _, dr := range datasourceResults {
		results = append(results, &schema.ConditionResult{
			Name:      dr.Name,
			Value:     dr.Value,
			Valid:     condition.Benchmark.Valid(dr.Value),
			Condition: condition,
		})

	}
	return results, nil
}

func mergeConditionResultToRuleResult(tx context.Context, results []*schema.ConditionResult) *schema.RuleResult {
	rr := &schema.RuleResult{
		Qualified: true,
		Detail:    make([]string, 0),
	}
	for _, result := range results {
		if !result.Valid {
			rr.Qualified = false
			rr.Detail = append(rr.Detail, result.String())
		}
	}
	return rr
}

func (i *impl) StartAlert(ctx context.Context, alert *ent.Alert) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if _, ok := i.alerts[alert.ID]; ok {
		//
		return nil
	}
	ticker := time.NewTicker(time.Second * time.Duration(alert.Rule.Interval))
	i.alerts[alert.ID] = ticker
	go func() {
		for {
			ctx := log.NewDefaultContext()
			for range ticker.C {
				ruleResult, err := i.evaluatesRule(ctx, alert.Rule)
				if err != nil {
					log.ErrorWithContext(ctx, "evaluate error", "err", err.Error())
					continue
				}
				for _, channelID := range alert.Channels {
					channel, err := i.dal.GetChannelByID(ctx, channelID)
					if err != nil {
						log.ErrorWithContext(ctx, "get channel by id error", "err", err.Error())
						continue
					}
					if ruleResult.Qualified {
						//
					} else {
						// send alert to channels and update alert status
						err := i.sendAlert(ctx, alert, channel, ruleResult)
						if err != nil {
							log.ErrorWithContext(ctx, "send alert error", "err", err.Error())
							continue
						}
					}
					alert.State = ruleResultToAlertState(ruleResult)
					_, err = i.dal.UpdateAlert(ctx, alert.ID, alert)
					if err != nil {
						log.ErrorWithContext(ctx, "update alert error", "err", err.Error())
						continue
					}
				}
			}
		}
	}()
	return nil
}

func ruleResultToAlertState(result *schema.RuleResult) schema.AlertState {
	switch result.Qualified {
	default:
		// TODO: status may be pending...but now just has two status
		return schema.AlertStateOK
	case true:
		return schema.AlertStateOK
	case false:
		return schema.AlertStateAlerting
	}
}

func (i *impl) stopAlert(ctx context.Context, alert *ent.Alert) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if ticker, ok := i.alerts[alert.ID]; ok {
		ticker.Stop()
		delete(i.alerts, alert.ID)
	}
	return nil
}

func (i *impl) sendAlert(ctx context.Context, alert *ent.Alert, channel *ent.Channel, result *schema.RuleResult) error {
	switch channel.Type {
	default:
		return fmt.Errorf("unknown channel type")
	case schema.ChannelTypeWebhook:
		return i.sendWebhookAlert(ctx, alert, channel.Detail.Webhook, result)
	}
}

func (i *impl) sendWebhookAlert(ctx context.Context, alert *ent.Alert, webhook *schema.Webhook, result *schema.RuleResult) error {
	postData := map[string]string{
		"msg":         result.String(),
		"alert title": alert.Name,
	}
	data, err := json.Marshal(postData)
	if err != nil {
		return err
	}
	// TODO: temp ignore response
	_, err = http.DefaultClient.Post(webhook.URL, echo.MIMEApplicationJSON, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	return nil
}
