package sub

import (
	"fmt"
)

type BenchmarkType uint

const (
	BenchmarkTypeTypeUndefined BenchmarkType = 0
	BenchmarkTypeEqual         BenchmarkType = 1
	BenchmarkTypeGT            BenchmarkType = 2
	BenchmarkTypeGTE           BenchmarkType = 3
	BenchmarkTypeLT            BenchmarkType = 4
	BenchmarkTypeLTE           BenchmarkType = 5
	BenchmarkTypeRange         BenchmarkType = 6
	BenchmarkTypeOutOfRange    BenchmarkType = 7
	BenchmarkTypeNoValue       BenchmarkType = 8 // TODO:to be implemented
	BenchmarkTypeHasValue      BenchmarkType = 9 // TODO:to be implemented
)

type Condition struct {
	DatasourceID int64     `json:"datasourceId"`
	Script       string    `json:"script"`
	Benchmark    Benchmark `json:"benchmark"`
}

type Benchmark struct {
	Type        BenchmarkType `json:"type"`
	SingleValue float64       `json:"singleValue"`
	Range       struct {
		Lower float64 `json:"lower"`
		Upper float64 `json:"upper"`
	}
}

func (b *Benchmark) String() string {
	switch b.Type {
	default:
		return fmt.Sprintf("unsupported benchmark:%d", b.Type)
	case BenchmarkTypeEqual:
		return fmt.Sprintf("!=%.2f", b.SingleValue)
	case BenchmarkTypeGT:
		return fmt.Sprintf("<=%.2f", b.SingleValue)
	case BenchmarkTypeGTE:
		return fmt.Sprintf("<%.2f", b.SingleValue)
	case BenchmarkTypeLT:
		return fmt.Sprintf(">=%.2f", b.SingleValue)
	case BenchmarkTypeLTE:
		return fmt.Sprintf(">%.2f", b.SingleValue)
	case BenchmarkTypeRange:
		return fmt.Sprintf("not in (%.2f,%.2f)", b.Range.Lower, b.Range.Upper)
	case BenchmarkTypeOutOfRange:
		return fmt.Sprintf("in (%.2f,%.2f)", b.Range.Lower, b.Range.Upper)
	}
}

type ConditionResult struct {
	Name             string            `json:"name"`
	Value            float64           `json:"value"`
	DatasourceResult *DatasourceResult `json:"datasourceResult"`
	Alerting         bool              `json:"alerting"`
	Condition        *Condition        `json:"condition"`
}

func (r *ConditionResult) String() string {
	if r.Alerting {
		return fmt.Sprintf("%s should %s, but is %.4f", r.Name, r.Condition.Benchmark.String(), r.Value)
	}
	return fmt.Sprintf("%s should %s, and is %.2f", r.Name, r.Condition.Benchmark.String(), r.Value)
}

// NotValid test the benchmark result is not valid
func (b *Benchmark) NotValid(value float64) bool {
	switch b.Type {
	default:
		return false
	case BenchmarkTypeEqual:
		return value == b.SingleValue
	case BenchmarkTypeGT:
		return value > b.SingleValue
	case BenchmarkTypeGTE:
		return value >= b.SingleValue
	case BenchmarkTypeLT:
		return value < b.SingleValue
	case BenchmarkTypeLTE:
		return value <= b.SingleValue
	case BenchmarkTypeRange:
		return value >= b.Range.Lower && value <= b.Range.Upper
	case BenchmarkTypeOutOfRange:
		return !(value >= b.Range.Lower && value <= b.Range.Upper)
	}
}
