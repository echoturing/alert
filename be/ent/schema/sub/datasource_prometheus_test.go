package sub

import (
	"context"
	"testing"
)

func TestPrometheusConfig_Evaluates(t *testing.T) {
	cfg := PrometheusConfig{Endpoint: "http://127.0.0.1:9090"}
	err := cfg.Evaluates(context.Background(), "sum(third_part_api_elapsed_seconds_count{source_type=\"1\"})")
	if err != nil {
		t.Error(err.Error())
		return
	}

}
