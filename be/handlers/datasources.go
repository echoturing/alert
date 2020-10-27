package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/services"
)

type ListDatasourceRequest struct {
}

type ListDatasourceReply struct {
	Code    int               `json:"code"`
	List    []*ent.Datasource `json:"list"`
	Message string            `json:"message"`
}

func (i *impl) ListDatasource(c echo.Context) error {
	ctx := c.Request().Context()
	list, err := i.service.ListDatasource(ctx)
	if err != nil {
		return err
	}
	return c.JSON(200, ListDatasourceReply{
		Code:    0,
		List:    list,
		Message: "",
	})
}

func (i *impl) CreateDatasource(c echo.Context) error {
	ctx := c.Request().Context()
	ds := &ent.Datasource{}
	if err := c.Bind(ds); err != nil {
		return err
	}
	_, err := i.service.CreateDatasource(ctx, ds)
	return err
}

func (i *impl) UpdateDatasource(c echo.Context) error {
	ctx := c.Request().Context()
	req := &services.UpdateDatasourceRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		return err
	}
	_, err = i.service.UpdateDatasource(ctx, id, req)
	if err != nil {
		return err
	}
	return nil
}
