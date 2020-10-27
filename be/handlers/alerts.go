package handlers

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/ent/schema"
	"github.com/echoturing/alert/services"
)

type CreateAlertRequest struct {
	Name    string       `json:"name"`
	Channel []int64      `json:"channel"`
	Rule    *schema.Rule `json:"rule"`
}

func (i *impl) CreateAlert(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	req := &CreateAlertRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	alert, err := i.service.CreateAlert(ctx, &ent.Alert{
		Name:     req.Name,
		Channels: req.Channel,
		Rule:     *req.Rule,
		Status:   schema.AlertStatus(1),
	})
	if err != nil {
		return nil, err
	}
	return alert, nil
}

type ListAlertsRequest struct {
	Status      schema.AlertStatus `query:"status"`
	AlertStatus schema.AlertState  `query:"alertStatus"`
}

type ListAlertsReply struct {
	List []*ent.Alert `json:"list"`
}

func (i *impl) ListAlerts(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	req := ListAlertsRequest{}

	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	alertList, err := i.service.ListAlerts(ctx, req.Status, req.AlertStatus)
	if err != nil {
		return nil, err
	}
	resp := &ListAlertsReply{
		List: alertList,
	}
	return resp, nil
}

func (i *impl) UpdateAlert(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	req := &services.UpdateAlertRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		return nil, err
	}
	alert, err := i.service.UpdateAlert(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return alert, nil
}
