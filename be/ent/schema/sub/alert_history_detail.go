package sub

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type AlertHistoryDetail struct {
	Rule       *Rule       `json:"rule"`
	RuleResult *RuleResult `json:"ruleResult"`
}

func (r AlertHistoryDetail) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *AlertHistoryDetail) Scan(src interface{}) error {
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
