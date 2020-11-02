package services

import (
	"context"
	"fmt"
	"time"

	"github.com/echoturing/log"

	"github.com/echoturing/alert/ent"
	alertFields "github.com/echoturing/alert/ent/alert"
	"github.com/echoturing/alert/ent/schema"
	"github.com/echoturing/alert/ent/schema/sub"
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
	Rule     *sub.Rule          `json:"rule"`
	Status   schema.AlertStatus `json:"status"`
}

func (i *impl) UpdateAlert(ctx context.Context, id int64, update *UpdateAlertRequest) (*ent.Alert, error) {
	count, err := i.dal.UpdateAlert(ctx, id, &ent.Alert{
		Name:     update.Name,
		Channels: update.Channels,
		Rule:     *update.Rule,
		Status:   update.Status,
	}, []string{alertFields.FieldName,
		alertFields.FieldChannels,
		alertFields.FieldRule,
		alertFields.FieldStatus,
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
		err = i.StopAlert(ctx, alert)
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

func (i *impl) evaluatesRule(ctx context.Context, rule sub.Rule) (*sub.RuleResult, error) {
	var final sub.RuleResult
	for _, condition := range rule.Conditions {
		conditionResults, err := i.evaluatesCondition(ctx, condition.Condition)
		if err != nil {
			return nil, err
		}
		ruleResult := mergeConditionResultToRuleResult(ctx, conditionResults)
		switch condition.Type {
		default:
			log.ErrorWithContext(ctx, "unknown condition type", "condition", condition)
		case sub.ConditionRelationTypeOr, sub.ConditionRelationTypeUndefined:
			final.Qualified = final.Qualified || ruleResult.Qualified
		case sub.ConditionRelationTypeAnd:
			final.Qualified = final.Qualified && ruleResult.Qualified
		}
		final.Detail = append(final.Detail, ruleResult.Detail...)
	}
	return &final, nil
}

func (i *impl) evaluatesCondition(ctx context.Context, condition *sub.Condition) ([]*sub.ConditionResult, error) {
	// get datasource
	datasource, err := i.dal.GetDatasourceByID(ctx, condition.DatasourceID)
	if err != nil {
		return nil, err
	}
	datasourceResults, err := i.evaluatesDatasource(ctx, datasource, condition.Script)
	if err != nil {
		return nil, err
	}

	results := make([]*sub.ConditionResult, 0, len(datasourceResults))
	for _, dr := range datasourceResults {
		results = append(results, &sub.ConditionResult{
			Name:      dr.Name,
			Value:     dr.Value,
			Valid:     !condition.Benchmark.NotValid(dr.Value),
			Condition: condition,
		})

	}
	return results, nil
}

func mergeConditionResultToRuleResult(tx context.Context, results []*sub.ConditionResult) *sub.RuleResult {
	rr := &sub.RuleResult{
		Qualified: true,
	}
	for _, result := range results {
		if !result.Valid {
			rr.Qualified = false
		}
		rr.Detail = append(rr.Detail, result)
	}
	return rr
}

func (i *impl) StartAlert(ctx context.Context, alert *ent.Alert) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if ticker, ok := i.alerts[alert.ID]; ok {
		ticker.Stop()
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
				current := ruleResultToAlertState(ruleResult, alert)
				prev := alert.State
				if current != prev {
					alert.State = current
					alert, err = i.dal.UpdateAlert(ctx, alert.ID, alert, []string{alertFields.FieldState})
					if err != nil {
						log.ErrorWithContext(ctx, "update alert error", "err", err.Error())
						continue
					}
					if alert.State == schema.AlertStateAlerting {
						for _, channelID := range alert.Channels {
							channel, err := i.dal.GetChannelByID(ctx, channelID)
							if err != nil {
								log.ErrorWithContext(ctx, "get channel by id error", "err", err.Error())
								continue
							}
							err = i.sendAlert(ctx, alert, channel, ruleResult)
							if err != nil {
								log.ErrorWithContext(ctx, "send alert error", "err", err.Error())
								continue
							}
						}
					}
				}
			}
		}
	}()
	return nil
}

func ruleResultToAlertState(result *sub.RuleResult, alert *ent.Alert) schema.AlertState {
	switch result.Qualified {
	default:
		return schema.AlertStateOK
	case true:
		return schema.AlertStateOK
	case false:
		if alert.Rule.For == 0 {
			return schema.AlertStateAlerting
		}
		if alert.State == schema.AlertStateOK {
			return schema.AlertStatusPending
		}
		if alert.State == schema.AlertStatusPending {
			if time.Now().Unix()-alert.UpdatedAt.Unix() < alert.Rule.For {
				return schema.AlertStatusPending
			}
		}
		return schema.AlertStateAlerting
	}
}

func (i *impl) StopAlert(ctx context.Context, alert *ent.Alert) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if ticker, ok := i.alerts[alert.ID]; ok {
		ticker.Stop()
		delete(i.alerts, alert.ID)
	}
	return nil
}

func (i *impl) sendAlert(ctx context.Context, alert *ent.Alert, channel *ent.Channel, result *sub.RuleResult) error {
	switch channel.Type {
	default:
		return fmt.Errorf("unknown channel type")
	case sub.ChannelTypeWebhook:
		return channel.Detail.Webhook.SendWebhookAlert(ctx, &sub.WebhookMessage{
			Title:   alert.Name,
			Message: result.String(),
		})
	}
}
