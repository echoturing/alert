package schema

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type DatasourceType int8

const (
	DatasourceTypeUndefined  DatasourceType = 0
	DatasourceTypeMySQL      DatasourceType = 1
	DatasourceTypePrometheus DatasourceType = 2
)

type DatasourceDetail struct {
	Mysql      *MySQLConfig      `json:"mysql,omitempty"`
	Prometheus *PrometheusConfig `json:"prometheus,omitempty"`
}

func (r DatasourceDetail) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *DatasourceDetail) Scan(src interface{}) error {
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
