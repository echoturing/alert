package services

import (
	"context"

	"github.com/echoturing/alert/ent"
)

func (i *impl) ListAlertHistories(ctx context.Context, limit, offset int) ([]*ent.AlertHistory, error) {
	return i.dal.ListAlertHistory(ctx, limit, offset)
}
