package dals

import (
	"context"

	"github.com/echoturing/alert/ent"
)

func (i *impl) CreateDatasource(ctx context.Context, datasource *ent.Datasource) (*ent.Datasource, error) {
	return i.client.Datasource.Create().SetName(datasource.Name).
		SetType(datasource.Type).
		SetDetail(datasource.Detail).Save(ctx)
}

func (i *impl) GetDatasourceByID(ctx context.Context, id int64) (*ent.Datasource, error) {
	return i.client.Datasource.Get(ctx, id)
}

func (i *impl) ListDatasource(ctx context.Context) ([]*ent.Datasource, error) {
	return i.client.Datasource.Query().All(ctx)
}

func (i *impl) UpdateDatasource(ctx context.Context, id int64, datasource *ent.Datasource) (*ent.Datasource, error) {
	return i.client.Datasource.UpdateOneID(id).SetName(datasource.Name).
		SetType(datasource.Type).
		SetDetail(datasource.Detail).Save(ctx)
}
