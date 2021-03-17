package schema

import "entgo.io/ent"

// OS holds the schema definition for the OS entity.
type OS struct {
	ent.Schema
}

// Fields of the OS.
func (OS) Fields() []ent.Field {
	return nil
}

// Edges of the OS.
func (OS) Edges() []ent.Edge {
	return nil
}
