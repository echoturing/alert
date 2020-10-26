package dals

import (
	"context"
	"fmt"
	"time"

	"github.com/echoturing/alert/datasources"
)

const (
	tableDatasource        = "datasource"
	DatasourceColumnID     = "id"
	DatasourceColumnName   = "name"
	DatasourceColumnType   = "type"
	DatasourceColumnDetail = "detail"

	DatasourceColumnCreatedAt = "createdAt"
	DatasourceColumnUpdatedAt = "updatedAt"
)

var dataSourceInsertColumns = []string{
	DatasourceColumnName,
	DatasourceColumnType,
	DatasourceColumnDetail,
}

var datasourceAllColumns = append(
	append([]string{DatasourceColumnID}, dataSourceInsertColumns...),
	AlertColumnCreatedAt,
	AlertColumnUpdatedAt,
)

func (i *impl) CreateDatasource(ctx context.Context, datasource *datasources.Datasource) (*datasources.Datasource, error) {
	statement := fmt.Sprintf("insert into %s (%s) values (%s)", tableDatasource, listToStrWithQuotes(dataSourceInsertColumns), generatePlaceholders(len(dataSourceInsertColumns)))
	res, err := i.db.ExecContext(ctx, statement, datasource.Name, datasource.Type, datasource.Detail)
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, statement, datasource)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, statement, datasource)
	}
	datasource.ID = id
	datasource.CreatedAt = time.Now()
	datasource.UpdatedAt = time.Now()
	return datasource, nil
}

func (i *impl) GetDatasourceByID(ctx context.Context, id int64) (*datasources.Datasource, error) {
	panic("implement me")
}

func (i *impl) ListDatasource(ctx context.Context) ([]*datasources.Datasource, error) {
	query := fmt.Sprintf("select %s from %s where 1=1", listToStrWithQuotes(datasourceAllColumns), tableDatasource)
	var values []interface{}
	rows, err := i.db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, query, values)
	}
	results := make([]*datasources.Datasource, 0)
	for rows.Next() {
		a := &datasources.Datasource{}
		err := rows.Scan(&a.ID, &a.Name, &a.Type, &a.Detail, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%w:%s	%#v", err, query, values)
		}
		results = append(results, a)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, query, values)
	}
	return results, nil
}

func (i *impl) UpdateDatasource(ctx context.Context, id int64, kvs map[string]interface{}) (int64, error) {
	panic("implement me")
}
