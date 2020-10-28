// Code generated by entc, DO NOT EDIT.

package channel

import (
	"time"
)

const (
	// Label holds the string label denoting the channel type in the database.
	Label = "channel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the channel in the database.
	Table = "channel"
)

// Columns holds all SQL columns for channel fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldDetail,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
)
