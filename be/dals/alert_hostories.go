package dals

import (
	"context"

	"github.com/echoturing/alert/ent"
)

func (i *impl) CreateAlertHistory(ctx context.Context, history *ent.AlertHistory) (*ent.AlertHistory, error) {
	return i.client.AlertHistory.Create().
		SetAlertID(history.AlertID).
		SetAlertName(history.AlertName).
		SetDetail(history.Detail).Save(ctx)

}

func (i *impl) ListAlertHistory(ctx context.Context, limit, offset int) ([]*ent.AlertHistory, error) {
	return i.client.AlertHistory.Query().
		Limit(limit).
		Offset(offset).All(ctx)
}
