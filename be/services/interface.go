package services

import (
	"context"
	"sync"
	"time"

	"github.com/echoturing/alert/alerts"
	"github.com/echoturing/alert/dals"
	"github.com/echoturing/alert/datasources"
)

type Interface interface {
	CreateAlert(ctx context.Context, alert *alerts.Alert) (*alerts.Alert, error)
	ListAlerts(ctx context.Context, status alerts.Status, alertStatus alerts.AlertStatus) ([]*alerts.Alert, error)
	UpdateAlert(ctx context.Context, id int64, update *UpdateAlertRequest) (int64, error)

	CreateDatasource(ctx context.Context, datasource *datasources.Datasource) (*datasources.Datasource, error)
	ListDatasource(ctx context.Context) ([]*datasources.Datasource, error)
	UpdateDatasource(ctx context.Context, id int64, update *UpdateDatasourceRequest) (int64, error)
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
