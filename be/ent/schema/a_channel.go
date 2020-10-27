package schema

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type ChannelType int8

const (
	ChannelTypeUndefined ChannelType = 0
	ChannelTypeWebhook   ChannelType = 1
)

type Webhook struct {
	URL string `json:"url"`
}

type ChannelDetail struct {
	Webhook *Webhook `json:"webhook"`
}

func (r ChannelDetail) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ChannelDetail) Scan(src interface{}) error {
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
