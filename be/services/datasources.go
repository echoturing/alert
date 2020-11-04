package services

import (
	"context"
	"fmt"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/ent/schema/sub"
)

func (i *impl) CreateDatasource(ctx context.Context, datasource *ent.Datasource) (*ent.Datasource, error) {
	err := i.connectDatasource(ctx, datasource)
	if err != nil {
		return nil, err
	}
	return i.dal.CreateDatasource(ctx, datasource)
}

func (i *impl) ListDatasource(ctx context.Context) ([]*ent.Datasource, error) {
	return i.dal.ListDatasource(ctx)
}

type UpdateDatasourceRequest struct {
	Name   string                `json:"name"`
	Type   sub.DatasourceType    `json:"type"`
	Detail *sub.DatasourceDetail `json:"detail"`
}

func (i *impl) UpdateDatasource(ctx context.Context, id int64, update *UpdateDatasourceRequest) (*ent.Datasource, error) {
	ds := &ent.Datasource{
		Name:   update.Name,
		Type:   update.Type,
		Detail: *update.Detail,
	}
	err := i.connectDatasource(ctx, ds)
	if err != nil {
		return nil, err
	}
	return i.dal.UpdateDatasource(ctx, id, ds)
}

// Connect test the datasource is valid
func (i *impl) connectDatasource(ctx context.Context, datasource *ent.Datasource) error {
	switch datasource.Type {
	default:
		return fmt.Errorf("unknown datasource type")
	case sub.DatasourceTypeMySQL:
		return datasource.Detail.Mysql.Connect(ctx)
	case sub.DatasourceTypePrometheus:
		return datasource.Detail.Prometheus.Connect(ctx)
	}
}

func (_ *impl) evaluatesDatasource(ctx context.Context, datasource *ent.Datasource, script string) ([]*sub.DatasourceResult, error) {
	switch datasource.Type {
	default:
		return nil, fmt.Errorf("unknow type:%d", datasource.Type)
	case sub.DatasourceTypeMySQL:
		return datasource.Detail.Mysql.Evaluates(ctx, script)
	case sub.DatasourceTypePrometheus:
		return datasource.Detail.Prometheus.Evaluates(ctx, script)
	}
}
