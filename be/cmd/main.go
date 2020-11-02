package main

import (
	"context"
	"fmt"
	"os"

	"github.com/echoturing/log"
	"github.com/facebook/ent/dialect"
	entSql "github.com/facebook/ent/dialect/sql"
	flag "github.com/spf13/pflag"

	"github.com/echoturing/alert/common"
	"github.com/echoturing/alert/dals"
	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/handlers"
	"github.com/echoturing/alert/routers"
	"github.com/echoturing/alert/services"
)

var (
	host string
	port int

	mysqlHost     string
	mysqlPort     int
	mysqlUser     string
	mysqlPassword string
	mysqlDB       string
)

func parseFlag() {
	flag.StringVarP(&host, "hose", "", "0.0.0.0", "server bind host")
	flag.IntVarP(&port, "port", "", 8899, "server bind port")
	flag.StringVarP(&mysqlHost, "mysqlHost", "", "mysql", "the mysql that this service connect to")
	flag.IntVarP(&mysqlPort, "mysqlPort", "", 3306, "")
	flag.StringVarP(&mysqlUser, "mysqlUser", "", "alert", "")
	flag.StringVarP(&mysqlPassword, "mysqlPassword", "", "123456", "")
	flag.StringVarP(&mysqlDB, "mysqlDB", "", "alert", "")
	flag.Parse()
}

func main() {
	parseFlag()
	dbConn, err := common.NewMysqlConnection(
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDB,
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
	_ = entClient.Debug().Schema.WriteTo(context.Background(), os.Stdout)
	err = entClient.Schema.Create(context.Background())
	if err != nil {
		log.Error(err.Error())
	}
	err = service.StartAllAlert(log.NewDefaultContext())
	if err != nil {
		log.Error("start all alerts error", "err", err.Error())
	}
	routers.Route(host, port, handler, "alert")
}
