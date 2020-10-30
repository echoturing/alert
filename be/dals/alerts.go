package dals

import (
	"context"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/ent/alert"
	"github.com/echoturing/alert/ent/predicate"
	"github.com/echoturing/alert/ent/schema"
)

func (i *impl) ListAlerts(ctx context.Context, status schema.AlertStatus, state schema.AlertState) ([]*ent.Alert, error) {
	var where []predicate.Alert
	if status != schema.StatusUndefined {
		where = append(where, alert.StatusEQ(status))
	}
	if state != schema.AlertStateUndefined {
		where = append(where, alert.StateEQ(state))
	}
	return i.client.Alert.Query().Where(where...).All(ctx)
}

func (i *impl) CreateAlert(ctx context.Context, alert *ent.Alert) (*ent.Alert, error) {
	return i.client.Alert.Create().
		SetName(alert.Name).
		SetChannels(alert.Channels).
		SetRule(alert.Rule).
		SetState(alert.State).
		SetStatus(alert.Status).Save(ctx)
}

func (i *impl) UpdateAlert(ctx context.Context, id int64, a *ent.Alert, updateKeys []string) (*ent.Alert, error) {
	updater := i.client.Alert.UpdateOneID(id)
	for _, key := range updateKeys {
		switch key {
		case alert.FieldName:
			updater = updater.SetName(a.Name)
		case alert.FieldChannels:
			updater = updater.SetChannels(a.Channels)
		case alert.FieldRule:
			updater = updater.SetRule(a.Rule)
		case alert.FieldState:
			updater = updater.SetState(a.State)
		case alert.FieldStatus:
			updater = updater.SetStatus(a.Status)
		}
	}
	return updater.Save(ctx)
}

func (i *impl) GetAlertByID(ctx context.Context, id int64) (*ent.Alert, error) {
	return i.client.Alert.Get(ctx, id)
}
