package sub

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
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

// DatasourceResult only support numeric types
type DatasourceResult struct {
	Name           string       `json:"name"`
	Kind           reflect.Kind `json:"kind"`
	ValueNumeric   float64      `json:"valueNumeric"`
	ValueInterface interface{}  `json:"valueInterface"`
	IsMetrics      bool         `json:"isMetrics"`
	Msg            string       `json:"-"`
}

func (dr *DatasourceResult) String() string {
	return fmt.Sprintf("%s|%s|%f|%s|%t|%s", dr.Name, dr.Kind, dr.ValueNumeric, dr.ValueInterface, dr.IsMetrics, dr.Msg)
}

func (dr *DatasourceResult) TryConvertToFloat() {
	if CanBeNumeric(dr.Kind) {
		if bs, ok := dr.ValueInterface.([]byte); ok {
			dataStr := string(bs)
			floatData, err := strconv.ParseFloat(dataStr, 64)
			if err != nil {
				dr.Msg = err.Error()
				return
			}
			dr.ValueNumeric = floatData
			dr.IsMetrics = true
		}
	}
}
