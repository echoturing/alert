package sub

import (
	"context"
	"testing"

	"github.com/echoturing/log"
)

func TestPrometheusConfig_Evaluates(t *testing.T) {
	cfg := PrometheusConfig{Endpoint: "http://127.0.0.1:9090"}
	res, err := cfg.Evaluates(context.Background(), "third_part_api_elapsed_seconds_count{source_type=\"1\"}")
	if err != nil {
		t.Error(err.Error())
		return
	}
	log.Debug("eval", "res", res)
}
