package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("user_id"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owned_by", User.Type).
			Ref("cards").
			Unique().
			Required().
			Field("user_id"),
		edge.From("faved_users", User.Type).
			Ref("favorite_cards").
			Through("player_favorite_cards", PlayerFavoriteCards.Type),
	}
}
