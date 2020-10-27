package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Datasource holds the schema definition for the Datasource entity.
type Datasource struct {
	ent.Schema
}

const (
	tableDatasource        = "datasource"
	DatasourceColumnID     = "id"
	DatasourceColumnName   = "name"
	DatasourceColumnType   = "type"
	DatasourceColumnDetail = "detail"
)

func (Datasource) Config() ent.Config {
	return ent.Config{
		Table: tableDatasource,
	}
}

// Fields of the Datasource.
func (Datasource) Fields() []ent.Field {
	return []ent.Field{
		field.Int64(DatasourceColumnID),
		field.String(DatasourceColumnName),
		field.Int8(DatasourceColumnType).GoType(DatasourceType(0)),
		field.String(DatasourceColumnDetail).GoType(&DatasourceDetail{}),
		field.Time(ColumnCreatedAt).Default(time.Now).Immutable().Optional(),
		field.Time(ColumnUpdatedAt).Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Datasource.
func (Datasource) Edges() []ent.Edge {
	return nil
}
