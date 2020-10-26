package dals

import (
	"context"
	"database/sql"

	"github.com/echoturing/alert/alerts"
	"github.com/echoturing/alert/channels"
	"github.com/echoturing/alert/datasources"
)

type Interface interface {
	CreateDatasource(ctx context.Context, datasource *datasources.Datasource) (*datasources.Datasource, error)
	GetDatasourceByID(ctx context.Context, id int64) (*datasources.Datasource, error)
	ListDatasource(ctx context.Context) ([]*datasources.Datasource, error)
	UpdateDatasource(ctx context.Context, id int64, kvs map[string]interface{}) (int64, error)

	ListAlerts(ctx context.Context, status alerts.Status, alertStatus alerts.AlertStatus) ([]*alerts.Alert, error)
	CreateAlert(ctx context.Context, alert *alerts.Alert) (*alerts.Alert, error)
	UpdateAlert(ctx context.Context, id int64, kvs map[string]interface{}) (int64, error)
	GetAlertByID(ctx context.Context, id int64) (*alerts.Alert, error)

	GetChannelByID(ctx context.Context, id int64) (*channels.Channel, error)
}

type impl struct {
	db *sql.DB
}

var _ Interface = (*impl)(nil)

func NewDALInterface(db *sql.DB) Interface {
	return &impl{db: db}
}
