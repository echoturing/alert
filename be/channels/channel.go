package channels

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ChannelType int

const (
	ChannelTypeUndefined ChannelType = 0
	ChannelTypeWebhook   ChannelType = 1
)

type Webhook struct {
	URL string `json:"url"`
}

type Channel struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	Type      ChannelType `json:"type"`
	Detail    *Detail     `json:"detail"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type Detail struct {
	Webhook *Webhook `json:"webhook"`
}

func (r Detail) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Detail) Scan(src interface{}) error {
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
