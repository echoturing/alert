// Code generated by entc, DO NOT EDIT.

package alerthistory

import (
	"time"
)

const (
	// Label holds the string label denoting the alerthistory type in the database.
	Label = "alert_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAlertID holds the string denoting the alert_id field in the database.
	FieldAlertID = "alert_id"
	// FieldAlertName holds the string denoting the alert_name field in the database.
	FieldAlertName = "alert_name"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the alerthistory in the database.
	Table = "alert_history"
)

// Columns holds all SQL columns for alerthistory fields.
var Columns = []string{
	FieldID,
	FieldAlertID,
	FieldAlertName,
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
