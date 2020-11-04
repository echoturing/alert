package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"

	"github.com/echoturing/alert/ent/schema/sub"
)

// AlertHistory holds the schema definition for the AlertHistory entity.
type AlertHistory struct {
	ent.Schema
}

const (
	tableAlertHistory           = "alert_history"
	AlertHistoryColumnID        = "id"
	AlertHistoryColumnAlertID   = "alert_id"
	AlertHistoryColumnAlertName = "alert_name"
	AlertHistoryColumnDetail    = "detail"
)

func (AlertHistory) Config() ent.Config {
	return ent.Config{
		Table: tableAlertHistory,
	}
}

// Fields of the AlertHistory.
func (AlertHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64(AlertHistoryColumnID),
		field.Int64(AlertHistoryColumnAlertID),
		field.String(AlertHistoryColumnAlertName),
		field.Text(AlertHistoryColumnDetail).GoType(&sub.AlertHistoryDetail{}),
		field.Time(ColumnCreatedAt).Default(time.Now).Immutable().Optional(),
		field.Time(ColumnUpdatedAt).Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AlertHistory.
func (AlertHistory) Edges() []ent.Edge {
	return nil
}

func (AlertHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(AlertHistoryColumnAlertID, ColumnCreatedAt),
	}
}
