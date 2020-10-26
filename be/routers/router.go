package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/echoturing/alert/handlers"
)

func Route(host string, port int, handler handlers.Interface) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	apiV1 := e.Group("/api/v1")

	apiV1.GET("/alerts", handler.ListAlerts)
	apiV1.POST("/alerts", handler.CreateAlert)
	apiV1.PUT("/alerts/:id", handler.UpdateAlert)

	apiV1.GET("/datasource", handler.ListDatasource)
	apiV1.POST("/datasource", handler.CreateDatasource)
	apiV1.PUT("/datasource/:id", handler.UpdateDatasource)

	apiV1.GET("/channels", handler.ListChannels)
	apiV1.POST("/channels", handler.CreateChannel)
	apiV1.PUT("/channels/:id", handler.UpdateChannel)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}
