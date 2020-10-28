package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"

	"github.com/echoturing/alert/ent/schema/sub"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

const (
	tableChannel = "channel"

	ChannelColumnID     = "id"
	ChannelColumnName   = "name"
	ChannelColumnType   = "type"
	ChannelColumnDetail = "detail"
)

func (Channel) Config() ent.Config {
	return ent.Config{
		Table: tableChannel,
	}
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64(ChannelColumnID),
		field.String(ChannelColumnName),
		field.Int8(ChannelColumnType).GoType(sub.ChannelType(0)),
		field.String(ChannelColumnDetail).GoType(&sub.ChannelDetail{}),
		field.Time(ColumnCreatedAt).Default(time.Now).Immutable().Optional(),
		field.Time(ColumnUpdatedAt).Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return nil
}
