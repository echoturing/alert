package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/echoturing/alert/alerts"
	"github.com/echoturing/alert/alerts/rules"
	"github.com/echoturing/alert/services"
)

type CreateAlertRequest struct {
	Name    string      `json:"name"`
	Channel []int64     `json:"channel"`
	Rule    *rules.Rule `json:"rule"`
}

func (i *impl) CreateAlert(c echo.Context) error {
	ctx := c.Request().Context()
	req := &CreateAlertRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	_, err := i.service.CreateAlert(ctx, &alerts.Alert{
		Name:     req.Name,
		Channels: req.Channel,
		Rule:     req.Rule,
		Status:   alerts.Status(1),
	})
	if err != nil {
		return err
	}
	return nil
}

type ListAlertsRequest struct {
	Status      alerts.Status      `query:"status"`
	AlertStatus alerts.AlertStatus `query:"alertStatus"`
}

type ListAlertsReply struct {
	Code    int             `json:"code"`
	List    []*alerts.Alert `json:"list"`
	Message string          `json:"message"`
}

func (i *impl) ListAlerts(c echo.Context) error {
	ctx := c.Request().Context()
	req := ListAlertsRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	alertList, err := i.service.ListAlerts(ctx, req.Status, req.AlertStatus)
	if err != nil {
		return err
	}
	resp := &ListAlertsReply{
		Code: 0,
		List: alertList,
	}
	return c.JSON(200, resp)
}

func (i *impl) UpdateAlert(c echo.Context) error {
	ctx := c.Request().Context()
	req := &services.UpdateAlertRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		return err
	}
	_, err = i.service.UpdateAlert(ctx, id, req)
	if err != nil {
		return err
	}
	return nil
}
