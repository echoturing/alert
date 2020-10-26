package schema

import "github.com/facebook/ent"

// Datasource holds the schema definition for the Datasource entity.
type Datasource struct {
	ent.Schema
}

// Fields of the Datasource.
func (Datasource) Fields() []ent.Field {
	return nil
}

// Edges of the Datasource.
func (Datasource) Edges() []ent.Edge {
	return nil
}
