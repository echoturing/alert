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

	"github.com/echoturing/alert/alerts"
	"github.com/echoturing/alert/alerts/rules"
	"github.com/echoturing/alert/channels"
	"github.com/echoturing/alert/dals"
)

func (i *impl) CreateAlert(ctx context.Context, alert *alerts.Alert) (*alerts.Alert, error) {
	// TODO:
	return i.dal.CreateAlert(ctx, alert)
}

func (i *impl) ListAlerts(ctx context.Context, status alerts.Status, alertStatus alerts.AlertStatus) ([]*alerts.Alert, error) {
	return i.dal.ListAlerts(ctx, status, alertStatus)
}

type UpdateAlertRequest struct {
	Name     string        `json:"name"`
	Channels []int64       `json:"channels"`
	Rule     *rules.Rule   `json:"rule"`
	Status   alerts.Status `json:"status"`
}

func (i *impl) UpdateAlert(ctx context.Context, id int64, update *UpdateAlertRequest) (int64, error) {
	count, err := i.dal.UpdateAlert(ctx, id, map[string]interface{}{
		dals.AlertColumnName:     update.Name,
		dals.AlertColumnChannels: update.Channels,
		dals.AlertColumnRule:     update.Rule,
		dals.AlertColumnStatus:   update.Status,
	})
	if err != nil {
		return 0, err
	}
	alert, err := i.dal.GetAlertByID(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("get alert by id err:%w", err)
	}
	switch update.Status {
	case alerts.StatusOpen:
		err = i.StartAlert(ctx, alert)
		if err != nil {
			return 0, fmt.Errorf("start alert error:%w", err)
		}
	case alerts.StatusClose:
		err = i.StopAlert(ctx, alert)
		if err != nil {
			return 0, fmt.Errorf("stop alert error:%w", err)
		}
	}
	return count, nil
}

func (i *impl) StartAllAlert(ctx context.Context) error {
	// get all online alerts
	// start one
	als, err := i.dal.ListAlerts(ctx, alerts.StatusOpen, 0)
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

func (i *impl) StartAlert(ctx context.Context, alert *alerts.Alert) error {
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
				ruleResult, err := alert.Rule.Evaluates(ctx, i.dal.GetDatasourceByID)
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
					_, err = i.dal.UpdateAlert(ctx, alert.ID, map[string]interface{}{dals.AlertColumnAlertStatus: ruleResultToAlertStatus(ruleResult)})
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
func ruleResultToAlertStatus(result *rules.RuleResult) alerts.AlertStatus {
	switch result.Qualified {
	default:
		// TODO: status may be pending...but now just has two status
		return alerts.AlertStatusOK
	case true:
		return alerts.AlertStatusOK
	case false:
		return alerts.AlertStatusAlerting
	}
}

func (i *impl) StopAlert(ctx context.Context, alert *alerts.Alert) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if ticker, ok := i.alerts[alert.ID]; ok {
		ticker.Stop()
		delete(i.alerts, alert.ID)
	}
	return nil
}

func (i *impl) sendAlert(ctx context.Context, alert *alerts.Alert, channel *channels.Channel, result *rules.RuleResult) error {
	switch channel.Type {
	default:
		return fmt.Errorf("unknown channel type")
	case channels.ChannelTypeWebhook:
		return i.sendWebhookAlert(ctx, alert, channel.Detail.Webhook, result)
	}
}

func (i *impl) sendWebhookAlert(ctx context.Context, alert *alerts.Alert, webhook *channels.Webhook, result *rules.RuleResult) error {
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
