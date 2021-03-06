package sub

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
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
		Timeout:              time.Second * 5,
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

func initResults(columns []*sql.ColumnType) []*DatasourceResult {
	result := make([]*DatasourceResult, 0, len(columns))
	for i := 0; i < len(columns); i++ {
		dr := &DatasourceResult{
			Name: columns[i].Name(),
			Kind: columns[i].ScanType().Kind(),
		}
		result = append(result, dr)
	}
	return result
}

func CanBeNumeric(i reflect.Kind) bool {
	if isNumeric(i) {
		return true
	}
	return i == reflect.Struct || i == reflect.Slice
}

func isNumeric(i reflect.Kind) bool {
	return i >= reflect.Bool && i <= reflect.Float64
}

// resultsToValueInterfacePointer get the result value pointer
func resultsToValueInterfacePointer(results []*DatasourceResult) []interface{} {
	rt := make([]interface{}, 0, len(results))
	for _, result := range results {
		if isNumeric(result.Kind) {
			rt = append(rt, &result.ValueNumeric)
			result.IsMetrics = true
		} else {
			rt = append(rt, &result.ValueInterface)
		}
	}
	return rt
}

var errNoData = fmt.Errorf("no data found")

func (d *MySQLConfig) Evaluates(ctx context.Context, script string) ([]*DatasourceResult, error) {
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
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	var results []*DatasourceResult
	for rows.Next() {
		results = initResults(columnTypes)
		// only care about first row result
		err = rows.Scan(resultsToValueInterfacePointer(results)...)
		if err != nil {
			return nil, err
		}
		break
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errNoData
	}
	for _, i := range results {
		i.TryConvertToFloat()
	}
	return results, nil
}
