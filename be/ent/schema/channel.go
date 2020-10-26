package schema

import "github.com/facebook/ent"

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return nil
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return nil
}
