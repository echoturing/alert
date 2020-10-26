package services

import (
	"context"

	"github.com/echoturing/alert/dals"
	"github.com/echoturing/alert/datasources"
)

func (i *impl) CreateDatasource(ctx context.Context, datasource *datasources.Datasource) (*datasources.Datasource, error) {
	err := datasource.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return i.dal.CreateDatasource(ctx, datasource)
}

func (i *impl) ListDatasource(ctx context.Context) ([]*datasources.Datasource, error) {
	return i.dal.ListDatasource(ctx)
}

type UpdateDatasourceRequest struct {
	Type   datasources.DatasourceType `json:"type,omitempty"`
	Detail *datasources.Detail        `json:"detail"`
}

func (i *impl) UpdateDatasource(ctx context.Context, id int64, update *UpdateDatasourceRequest) (int64, error) {
	ds := &datasources.Datasource{
		Type:   update.Type,
		Detail: update.Detail,
	}
	err := ds.Connect(ctx)
	if err != nil {
		return 0, err
	}
	return i.dal.UpdateDatasource(ctx, id, map[string]interface{}{
		dals.DatasourceColumnType:   update.Type,
		dals.DatasourceColumnDetail: update.Type,
	})
}
