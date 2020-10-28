package schema

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"

	"github.com/echoturing/alert/ent/schema/sub"
)

type AlertState int8

const (
	AlertStateUndefined AlertState = 0
	AlertStateOK        AlertState = 1
	AlertStatusPending  AlertState = 2
	AlertStateAlerting  AlertState = 3
)

type AlertStatus int8

const (
	StatusUndefined AlertStatus = 0
	StatusOpen      AlertStatus = 1
	StatusClose     AlertStatus = 2
)

// Alert holds the schema definition for the Alert entity.
type Alert struct {
	ent.Schema
}

const (
	tableAlert             = "alert"
	AlertColumnID          = "id"
	AlertColumnName        = "name"
	AlertColumnChannels    = "channels"
	AlertColumnRule        = "rule"
	AlertColumnAlertState  = "state"
	AlertColumnAlertStatus = "status"
	ColumnCreatedAt        = "createdAt"
	ColumnUpdatedAt        = "updatedAt"
)

func (Alert) Config() ent.Config {
	return ent.Config{
		Table: tableAlert,
	}
}

type ChannelIDS []int64

func (r ChannelIDS) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ChannelIDS) Scan(src interface{}) error {
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

// Fields of the Alert.
func (Alert) Fields() []ent.Field {
	return []ent.Field{
		field.Int64(AlertColumnID),
		field.String(AlertColumnName),
		field.Text(AlertColumnChannels).GoType(&ChannelIDS{}),
		field.Text(AlertColumnRule).GoType(&sub.Rule{}),
		field.Int8(AlertColumnAlertStatus).GoType(AlertStatus(0)),
		field.Int8(AlertColumnAlertState).GoType(AlertState(0)),
		field.Time(ColumnCreatedAt).Default(time.Now).Immutable().Optional(),
		field.Time(ColumnUpdatedAt).Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Alert.
func (Alert) Edges() []ent.Edge {
	return nil
}
