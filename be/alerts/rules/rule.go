package rules

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/echoturing/log"

	"github.com/echoturing/alert/alerts/conditions"
)

type Rule struct {
	ID         int64                `json:"id"`
	Name       string               `json:"name"`
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
	Condition conditions.Condition  `json:"condition"`
}

func (a ConditionRelation) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type RuleResult struct {
	Qualified bool     `json:"qualified"`
	Detail    []string `json:"detail"`
}

func (rr *RuleResult) String() string {
	res := fmt.Sprintf("alerting:%t\n", rr.Qualified)
	for _, line := range rr.Detail {
		res += line + ";\n"
	}
	return res
}

// Evaluates evaluate every condition in sequence
// merge all condition result into a summary
func (r *Rule) Evaluates(ctx context.Context, datasourceGetter conditions.DatasourceGetter) (*RuleResult, error) {
	var final RuleResult
	for _, condition := range r.Conditions {
		conditionResults, err := condition.Condition.Evaluates(ctx, datasourceGetter)
		if err != nil {
			return nil, err
		}
		ruleResult := mergeToRuleResult(ctx, conditionResults)
		switch condition.Type {
		default:
			log.ErrorWithContext(ctx, "unknown condition type", "condition", condition)
		case ConditionRelationTypeOr, ConditionRelationTypeUndefined:
			final.Qualified = final.Qualified || ruleResult.Qualified
		case ConditionRelationTypeAnd:
			final.Qualified = final.Qualified && ruleResult.Qualified
		}
		final.Detail = append(final.Detail, ruleResult.Detail...)
	}
	return &final, nil
}

func mergeToRuleResult(tx context.Context, results []*conditions.Result) *RuleResult {
	rr := &RuleResult{
		Qualified: true,
		Detail:    make([]string, 0),
	}
	for _, result := range results {
		if !result.Valid {
			rr.Qualified = false
			rr.Detail = append(rr.Detail, result.String())
		}
	}
	return rr
}
