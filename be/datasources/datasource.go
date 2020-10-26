package datasources

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DatasourceType uint

const (
	DatasourceTypeUndefined  DatasourceType = 0
	DatasourceTypeMySQL      DatasourceType = 1
	DatasourceTypePrometheus DatasourceType = 2
)

type Detail struct {
	Mysql      *MySQLConfig      `json:"mysql,omitempty"`
	Prometheus *PrometheusConfig `json:"prometheus,omitempty"`
}

func (r Detail) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Detail) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch v := src.(type) {
	case []byte:
		if len(v) == 0 {
			return nil
		}
		return json.Unmarshal(v, r)
	case string:
		if v == "" {
			return nil
		}
		return json.NewDecoder(strings.NewReader(v)).Decode(r)
	default:
		return fmt.Errorf("cannot unmarshal %T:%v ", v, v)
	}
}

type Datasource struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Type      DatasourceType `json:"type,omitempty"`
	Detail    *Detail        `json:"detail"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

// Connect test the datasource is valid
func (d *Datasource) Connect(ctx context.Context) error {
	switch d.Type {
	default:
		return fmt.Errorf("unknown datasource type")
	case DatasourceTypeMySQL:
		return d.Detail.Mysql.Connect(ctx)
	case DatasourceTypePrometheus:
		// TODO(xiangxu)
	}
	return nil
}

type Result struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (d *Datasource) EvalScript(ctx context.Context, script string) ([]*Result, error) {
	switch d.Type {
	default:
		return nil, fmt.Errorf("unknow type:%d", d.Type)
	case DatasourceTypeMySQL:
		return d.Detail.Mysql.EvalScript(ctx, script)
	}
}
