package routers

import (
	"fmt"

	"github.com/echoturing/tools/middlewaretools"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/segmentio/ksuid"

	"github.com/echoturing/alert/handlers"
)

func Route(host string, port int, handler handlers.Interface, serviceName string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{Generator: func() string {
		return ksuid.New().String()
	}})) // request-id
	e.Use(middlewaretools.WrapContextWithUser())          // wrap request-id into context
	e.Use(middlewaretools.PrometheusMetrics(serviceName)) // prometheus
	e.Use(middlewaretools.AccessLog(nil))
	apiV1 := e.Group("/api/v1")

	apiV1.GET("/alerts", middlewaretools.HandlerFuncWrapper(handler.ListAlerts))
	apiV1.POST("/alerts", middlewaretools.HandlerFuncWrapper(handler.CreateAlert))
	apiV1.PUT("/alerts/:id", middlewaretools.HandlerFuncWrapper(handler.UpdateAlert))
	apiV1.GET("/alerts/:id/result", middlewaretools.HandlerFuncWrapper(handler.GetAlertResult))
	apiV1.GET("/alerts/test", middlewaretools.HandlerFuncWrapper(handler.TestAlert))

	apiV1.GET("/alert_histories", middlewaretools.HandlerFuncWrapper(handler.GetAlertHistories))

	apiV1.GET("/datasource", middlewaretools.HandlerFuncWrapper(handler.ListDatasource))
	apiV1.POST("/datasource", middlewaretools.HandlerFuncWrapper(handler.CreateDatasource))
	apiV1.PUT("/datasource/:id", middlewaretools.HandlerFuncWrapper(handler.UpdateDatasource))

	apiV1.GET("/channels", middlewaretools.HandlerFuncWrapper(handler.ListChannels))
	apiV1.POST("/channels", middlewaretools.HandlerFuncWrapper(handler.CreateChannel))
	apiV1.PUT("/channels/:id", middlewaretools.HandlerFuncWrapper(handler.UpdateChannel))

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}
