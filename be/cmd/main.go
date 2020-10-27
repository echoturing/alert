package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/echoturing/log"
	"github.com/facebook/ent/dialect"
	entSql "github.com/facebook/ent/dialect/sql"

	"github.com/echoturing/alert/common"
	"github.com/echoturing/alert/config"
	"github.com/echoturing/alert/dals"
	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/handlers"
	"github.com/echoturing/alert/routers"
	"github.com/echoturing/alert/services"
)

var (
	configFile string
	doMigrate  bool
)

func parseFlag() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "config path")
	flag.BoolVar(&doMigrate, "migrate", false, "migrate")
	flag.Parse()
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

	driver := entSql.OpenDB(dialect.MySQL, dbConn)
	entClient := ent.NewClient(ent.Driver(driver))

	dal := dals.NewDALInterface(dbConn, entClient)
	service := services.NewServiceInterface(dal)
	handler := handlers.NewHandlerInterface(service)
	fmt.Println(doMigrate)
	if doMigrate {
		var err error
		_ = entClient.Debug().Schema.WriteTo(context.Background(), os.Stdout)
		err = entClient.Schema.Create(context.Background())
		if err != nil {
			log.Error(err.Error())
		}
		return
	}

	err = service.StartAllAlert(log.NewDefaultContext())
	if err != nil {
		log.Error("start all alerts error", "err", err.Error())
	}
	routers.Route(cfg.Host, cfg.Port, handler, "alert")
}
