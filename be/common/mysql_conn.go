package common

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlConnection(
	user, password string,
	host string,
	port int,
	dbName string,
) (*sql.DB, error) {
	config := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", host, port),
		DBName:               dbName,
		Params:               nil,
		Collation:            "utf8mb4_general_ci",
		Loc:                  DefaultLoc,
		MaxAllowedPacket:     4 << 20, // 4 MiB
		Timeout:              time.Second * 30,
		ReadTimeout:          time.Minute,
		WriteTimeout:         time.Minute,
		InterpolateParams:    true,
		ParseTime:            true,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 30)
	return db, nil
}
