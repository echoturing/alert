package handlers

import (
	"github.com/labstack/echo"

	"github.com/echoturing/alert/services"
)

type Interface interface {
	ListAlerts(context echo.Context) (interface{}, error)
	CreateAlert(context echo.Context) (interface{}, error)
	UpdateAlert(context echo.Context) (interface{}, error)
	GetAlertResult(context echo.Context) (interface{}, error)

	GetAlertHistories(context echo.Context) (interface{}, error)

	ListDatasource(context echo.Context) (interface{}, error)
	CreateDatasource(context echo.Context) (interface{}, error)
	UpdateDatasource(context echo.Context) (interface{}, error)

	ListChannels(context echo.Context) (interface{}, error)
	CreateChannel(context echo.Context) (interface{}, error)
	UpdateChannel(context echo.Context) (interface{}, error)
}

type impl struct {
	service services.Interface
}

var _ Interface = (*impl)(nil)

func NewHandlerInterface(service services.Interface) Interface {
	return &impl{service: service}
}
