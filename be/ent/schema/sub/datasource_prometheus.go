package sub

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type PrometheusConfig struct {
	Endpoint string
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

func (p *PrometheusConfig) Evaluates(ctx context.Context, script string) error {
	client, err := newAPI(p.Endpoint)
	if err != nil {
		return err
	}
	v1api := v1.NewAPI(client)
	value, _, err := v1api.QueryRange(ctx, script, v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Hour,
	})
	if err != nil {
		return err
	}
	fmt.Println(value.Type())
	metric := value.(model.Matrix)
	fmt.Println(metric)
	return nil

}
