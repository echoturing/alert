package dals

import (
	"context"
	"database/sql"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/ent/schema"
)

type Interface interface {
	CreateDatasource(ctx context.Context, datasource *ent.Datasource) (*ent.Datasource, error)
	GetDatasourceByID(ctx context.Context, id int64) (*ent.Datasource, error)
	ListDatasource(ctx context.Context) ([]*ent.Datasource, error)
	UpdateDatasource(ctx context.Context, id int64, datasource *ent.Datasource) (*ent.Datasource, error)

	ListAlerts(ctx context.Context, status schema.AlertStatus, alertStatus schema.AlertState) ([]*ent.Alert, error)
	CreateAlert(ctx context.Context, alert *ent.Alert) (*ent.Alert, error)
	UpdateAlert(ctx context.Context, id int64, alert *ent.Alert) (*ent.Alert, error)
	GetAlertByID(ctx context.Context, id int64) (*ent.Alert, error)

	GetChannelByID(ctx context.Context, id int64) (*ent.Channel, error)
	ListChannels(ctx context.Context) ([]*ent.Channel, error)
	UpdateChannel(ctx context.Context, id int64, channel *ent.Channel) (*ent.Channel, error)
	CreateChannel(ctx context.Context, channel *ent.Channel) (*ent.Channel, error)
}

type impl struct {
	db     *sql.DB
	client *ent.Client
}

var _ Interface = (*impl)(nil)

func NewDALInterface(db *sql.DB, entClient *ent.Client) Interface {
	return &impl{db: db, client: entClient}
}
