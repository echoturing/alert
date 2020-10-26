package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/echoturing/alert/services"
)

type Interface interface {
	ListAlerts(context echo.Context) error
	CreateAlert(context echo.Context) error
	UpdateAlert(context echo.Context) error

	ListDatasource(context echo.Context) error
	CreateDatasource(context echo.Context) error
	UpdateDatasource(context echo.Context) error

	ListChannels(context echo.Context) error
	CreateChannel(context echo.Context) error
	UpdateChannel(context echo.Context) error
}

type impl struct {
	service services.Interface
}

var _ Interface = (*impl)(nil)

func NewHandlerInterface(service services.Interface) Interface {
	return &impl{service: service}
}
