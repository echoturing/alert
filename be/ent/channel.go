// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/echoturing/alert/ent/channel"
	"github.com/echoturing/alert/ent/schema/sub"
	"github.com/facebook/ent/dialect/sql"
)

// Channel is the model entity for the Channel schema.
type Channel struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type sub.ChannelType `json:"type,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail sub.ChannelDetail `json:"detail,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Channel) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},     // id
		&sql.NullString{},    // name
		&sql.NullInt64{},     // type
		&sub.ChannelDetail{}, // detail
		&sql.NullTime{},      // createdAt
		&sql.NullTime{},      // updatedAt
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Channel fields.
func (c *Channel) assignValues(values ...interface{}) error {
	if m, n := len(values), len(channel.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		c.Name = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[1])
	} else if value.Valid {
		c.Type = sub.ChannelType(value.Int64)
	}
	if value, ok := values[2].(*sub.ChannelDetail); !ok {
		return fmt.Errorf("unexpected type %T for field detail", values[2])
	} else if value != nil {
		c.Detail = *value
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[3])
	} else if value.Valid {
		c.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[4])
	} else if value.Valid {
		c.UpdatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this Channel.
// Note that, you need to call Channel.Unwrap() before calling this method, if this Channel
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Channel) Update() *ChannelUpdateOne {
	return (&ChannelClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Channel) Unwrap() *Channel {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Channel is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Channel) String() string {
	var builder strings.Builder
	builder.WriteString("Channel(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", detail=")
	builder.WriteString(fmt.Sprintf("%v", c.Detail))
	builder.WriteString(", createdAt=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Channels is a parsable slice of Channel.
type Channels []*Channel

func (c Channels) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
