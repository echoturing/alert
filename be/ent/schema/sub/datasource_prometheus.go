package sub

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"

	"github.com/echoturing/alert/common"
)

type PrometheusConfig struct {
	Endpoint string `json:"endpoint"`
}

func newAPI(endpoint string) (api.Client, error) {
	return api.NewClient(api.Config{
		Address: endpoint,
	})
}

func (p *PrometheusConfig) Connect(ctx context.Context) error {
	_, err := newAPI(p.Endpoint)
	return err
}

// calcAccumulator values sorted in time order
// and Prometheus Counter be reset to zero when service restart
// so we calc the accumulator to find every restart
func calcAccumulator(values []model.SamplePair) float64 {
	var total float64
	prev := values[0]
	for i := 1; i < len(values); i++ {
		current := values[i]
		accumulator := float64(current.Value - prev.Value)
		// real accumulated
		if accumulator > 0 {
			total += accumulator
			prev = current
		} else {
			// service restart,so we should add all init value
			total += float64(current.Value)
			prev = current
		}
	}
	return total
}

// wrapMetricsWithSumByInstance ..
// third_part_api_elapsed_seconds_count -> sum(third_part_api_elapsed_seconds_count) by (instance)
// third_part_api_elapsed_seconds_count{} -> sum(third_part_api_elapsed_seconds_count{}) by (instance)
func wrapMetricsWithSumByInstance(script string) string {
	return fmt.Sprintf("sum(%s) by (instance)", script)
}

// Evaluates temporary only care about absolute current days data.
// so QueryRange with
func (p *PrometheusConfig) Evaluates(ctx context.Context, script string) ([]*DatasourceResult, error) {
	client, err := newAPI(p.Endpoint)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	startOfDay := time.Date(now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0, common.DefaultLoc)
	v1api := v1.NewAPI(client)
	script = wrapMetricsWithSumByInstance(script)
	value, _, err := v1api.QueryRange(ctx, wrapMetricsWithSumByInstance(script), v1.Range{
		Start: startOfDay,
		End:   time.Now(),
		Step:  time.Minute,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*DatasourceResult, 0)
	switch value.Type() {
	default:
		return nil, fmt.Errorf("unsupported value type:%s", value.Type())
	case model.ValMatrix:
		matrix := value.(model.Matrix)
		dsr := &DatasourceResult{
			Name:      script,
			Kind:      reflect.Float64,
			IsMetrics: true,
		}
		for _, instance := range matrix {
			length := len(instance.Values)
			if length >= 2 { // we use first.val - last.val,so we should only care about length>=2
				dsr.ValueNumeric += calcAccumulator(instance.Values)
			}
		}
		result = append(result, dsr)
	}
	return result, nil

}
