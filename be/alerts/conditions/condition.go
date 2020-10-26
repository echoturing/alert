package conditions

import (
	"context"
	"fmt"

	"github.com/echoturing/alert/datasources"
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

type DatasourceGetter = func(_ context.Context, _ int64) (*datasources.Datasource, error)

type Condition struct {
	ID           int64     `json:"id"`
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
		return fmt.Sprintf("=%f", b.SingleValue)
	case BenchmarkTypeGT:
		return fmt.Sprintf(">%f", b.SingleValue)
	case BenchmarkTypeGTE:
		return fmt.Sprintf(">=%f", b.SingleValue)
	case BenchmarkTypeLT:
		return fmt.Sprintf("<%f", b.SingleValue)
	case BenchmarkTypeLTE:
		return fmt.Sprintf("<=%f", b.SingleValue)
	case BenchmarkTypeRange:
		return fmt.Sprintf("in (%f,%f)", b.Range.Lower, b.Range.Upper)
	case BenchmarkTypeOutOfRange:
		return fmt.Sprintf("exclude (%f,%f)", b.Range.Lower, b.Range.Upper)
	}
}

type Result struct {
	Name      string     `json:"name"`
	Value     float64    `json:"value"`
	Valid     bool       `json:"valid"` // the result is satisfy the condition?
	Condition *Condition `json:"condition"`
}

func (r *Result) String() string {
	return fmt.Sprintf("%s should %s, but is %f", r.Name, r.Condition.Benchmark.String(), r.Value)
}

// Evaluates  ...
// one condition may evaluates more results.
// like `select x,y from z limit 1`,we should return 2 result:named x,y field.
func (c *Condition) Evaluates(ctx context.Context, datasourceGetter DatasourceGetter) ([]*Result, error) {
	// get datasource
	datasource, err := datasourceGetter(ctx, c.DatasourceID)
	if err != nil {
		return nil, err
	}
	datasourceResults, err := datasource.EvalScript(ctx, c.Script)
	if err != nil {
		return nil, err
	}

	results := make([]*Result, 0, len(datasourceResults))
	for _, dr := range datasourceResults {
		results = append(results, &Result{
			Name:      dr.Name,
			Value:     dr.Value,
			Valid:     c.Benchmark.valid(dr.Value),
			Condition: c,
		})

	}
	return results, nil
}

func (b *Benchmark) valid(value float64) bool {
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
