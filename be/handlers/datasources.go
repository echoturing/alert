package handlers

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/services"
)

type ListDatasourceRequest struct {
}

type ListDatasourceReply struct {
	List []*ent.Datasource `json:"list"`
}

func (i *impl) ListDatasource(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	list, err := i.service.ListDatasource(ctx)
	if err != nil {
		return nil, err
	}
	return ListDatasourceReply{
		List: list,
	}, nil
}

func (i *impl) CreateDatasource(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	ds := &ent.Datasource{}
	if err := c.Bind(ds); err != nil {
		return nil, err
	}
	_, err := i.service.CreateDatasource(ctx, ds)
	return nil, err
}

func (i *impl) UpdateDatasource(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	req := &services.UpdateDatasourceRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		return nil, err
	}
	ds, err := i.service.UpdateDatasource(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return ds, nil
}
