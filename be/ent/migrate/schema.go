// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AlertColumns holds the columns for the "alert" table.
	AlertColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "channels", Type: field.TypeString, Size: 2147483647},
		{Name: "rule", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeInt8},
		{Name: "state", Type: field.TypeInt8},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AlertTable holds the schema information for the "alert" table.
	AlertTable = &schema.Table{
		Name:        "alert",
		Columns:     AlertColumns,
		PrimaryKey:  []*schema.Column{AlertColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ChannelColumns holds the columns for the "channel" table.
	ChannelColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeInt8},
		{Name: "detail", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ChannelTable holds the schema information for the "channel" table.
	ChannelTable = &schema.Table{
		Name:        "channel",
		Columns:     ChannelColumns,
		PrimaryKey:  []*schema.Column{ChannelColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DatasourceColumns holds the columns for the "datasource" table.
	DatasourceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeInt8},
		{Name: "detail", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// DatasourceTable holds the schema information for the "datasource" table.
	DatasourceTable = &schema.Table{
		Name:        "datasource",
		Columns:     DatasourceColumns,
		PrimaryKey:  []*schema.Column{DatasourceColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlertTable,
		ChannelTable,
		DatasourceTable,
	}
)

func init() {
}
