package sub

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type Rule struct {
	Interval   int64                `json:"interval"`
	For        int64                `json:"for"`
	Conditions []*ConditionRelation `json:"conditions"`
}

func (r Rule) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Rule) Scan(src interface{}) error {
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

type ConditionRelationType uint

const (
	ConditionRelationTypeUndefined ConditionRelationType = 0
	ConditionRelationTypeAnd       ConditionRelationType = 1
	ConditionRelationTypeOr        ConditionRelationType = 2
)

type ConditionRelation struct {
	Type      ConditionRelationType `json:"type"`
	Condition *Condition            `json:"condition"`
}

func (a ConditionRelation) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type RuleResult struct {
	Alerting bool               `json:"alerting"`
	Detail   []*ConditionResult `json:"detail"`
}

func (rr *RuleResult) String() string {
	res := fmt.Sprintf("alerting:%t\n", rr.Alerting)
	for _, c := range rr.Detail {
		res += c.String() + ";\n"
	}
	return res
}

func (rr *RuleResult) AlertMessage() string {
	res := fmt.Sprintf("alerting:%t\n", rr.Alerting)
	for _, c := range rr.Detail {
		if c.Alerting {
			res += c.String() + ";\n"
		}
	}
	return res
}
