package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PlayerFavoriteCards holds the schema definition for the PlayerFavoriteCards entity.
type PlayerFavoriteCards struct {
	ent.Schema
}

func (PlayerFavoriteCards) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "card_id"),
	}
}

// Fields of the PlayerFavoriteCards.
func (PlayerFavoriteCards) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id"),
		field.String("card_id").Unique(),
	}
}

func (PlayerFavoriteCards) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("card_id").
			Unique(),
	}
}

// Edges of the PlayerFavoriteCards.
func (PlayerFavoriteCards) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("card", Card.Type).
			Unique().
			Required().
			Field("card_id"),
	}
}
