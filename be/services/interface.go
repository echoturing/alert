package services

import (
	"context"
	"sync"
	"time"

	"github.com/echoturing/alert/dals"
	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/ent/schema"
	"github.com/echoturing/alert/ent/schema/sub"
)

type Interface interface {
	CreateAlert(ctx context.Context, alert *ent.Alert) (*ent.Alert, error)
	ListAlerts(ctx context.Context, status schema.AlertStatus, alertStatus schema.AlertState) ([]*ent.Alert, error)
	UpdateAlert(ctx context.Context, id int64, update *UpdateAlertRequest) (*ent.Alert, error)
	GetAlertResult(ctx context.Context, id int64) (*ent.Alert, *sub.RuleResult, error)
	TestRule(ctx context.Context, alert *ent.Alert) (*sub.RuleResult, error)

	ListChannels(ctx context.Context) ([]*ent.Channel, error)
	CreateChannel(ctx context.Context, channel *ent.Channel) (*ent.Channel, error)
	UpdateChannel(ctx context.Context, id int64, update *UpdateChannelRequest) (*ent.Channel, error)

	ListAlertHistories(ctx context.Context, limit, offset int) ([]*ent.AlertHistory, error)

	CreateDatasource(ctx context.Context, datasource *ent.Datasource) (*ent.Datasource, error)
	ListDatasource(ctx context.Context) ([]*ent.Datasource, error)
	UpdateDatasource(ctx context.Context, id int64, update *UpdateDatasourceRequest) (*ent.Datasource, error)
	StartAllAlert(ctx context.Context) error
}

type impl struct {
	dal dals.Interface

	alerts map[int64]*time.Ticker
	mutex  *sync.Mutex
}

var _ Interface = (*impl)(nil)

func NewServiceInterface(dal dals.Interface) Interface {

	return &impl{
		dal:    dal,
		alerts: make(map[int64]*time.Ticker),
		mutex:  &sync.Mutex{},
	}
}
