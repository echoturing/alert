package schema

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

type MySQLConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbName"`
}

func newMysqlConnection(
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
		Loc:                  loc,
		MaxAllowedPacket:     4 << 20, // 4 MiB
		Timeout:              time.Second * 3,
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

func (d *MySQLConfig) Connect(ctx context.Context) error {
	db, err := newMysqlConnection(d.User, d.Password, d.Host, d.Port, d.DBName)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Ping()
}

type DatasourceResult struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func initResults(columns []string) []*DatasourceResult {
	result := make([]*DatasourceResult, 0, len(columns))
	for i := 0; i < len(columns); i++ {
		result = append(result, &DatasourceResult{
			Name:  columns[i],
			Value: 0,
		})
	}
	return result
}

func resultsToValueInterfacePointer(results []*DatasourceResult) []interface{} {
	rt := make([]interface{}, 0, len(results))
	for _, result := range results {
		rt = append(rt, &result.Value)
	}
	return rt
}

func (d *MySQLConfig) EvalScript(ctx context.Context, script string) ([]*DatasourceResult, error) {
	db, err := newMysqlConnection(d.User, d.Password, d.Host, d.Port, d.DBName)
	if err != nil {
		return nil, err
	}
	// TODO: optimize to long connections
	defer db.Close()

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	results := initResults(columns)

	for rows.Next() {
		err = rows.Scan(resultsToValueInterfacePointer(results)...)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
