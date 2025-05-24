package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/yyewolf/entvis"
)

// Light holds the schema definition for the Light entity.
type Light struct {
	ent.Schema
}

// Fields of the Light.
func (Light) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("private_field").
			Annotations(
				entvis.Visibility("self", "admin"),
			).
			Optional(),
	}
}

// Edges of the Light.
func (Light) Edges() []ent.Edge {
	return nil
}
