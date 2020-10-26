package main

import (
	"flag"
	"fmt"

	"github.com/echoturing/log"

	"github.com/echoturing/alert/common"
	"github.com/echoturing/alert/config"
	"github.com/echoturing/alert/dals"
	"github.com/echoturing/alert/handlers"
	"github.com/echoturing/alert/routers"
	"github.com/echoturing/alert/services"
)

var (
	configFile string
)

func parseFlag() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "config path")
}

func main() {
	parseFlag()
	cfg := config.LoadConfig(configFile)
	mysqlConfig := cfg.Mysql
	dbConn, err := common.NewMysqlConnection(
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName,
	)
	if err != nil {
		fmt.Println("db conn error")
		return
	}

	dal := dals.NewDALInterface(dbConn)
	service := services.NewServiceInterface(dal)
	handler := handlers.NewHandlerInterface(service)

	err = service.StartAllAlert(log.NewDefaultContext())
	if err != nil {
		log.Error("start all alerts error", "err", err.Error())
	}
	routers.Route(cfg.Host, cfg.Port, handler)
}
