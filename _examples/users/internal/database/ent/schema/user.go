package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/yyewolf/entvis"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Annotations(
				entvis.Visibility("self", "admin"),
			),
		field.String("selected_card_id").
			Annotations(
				entvis.Visibility("self", "admin", "public"),
			).
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
		edge.To("favorite_cards", Card.Type).
			Through("player_favorite_cards", PlayerFavoriteCards.Type),
		edge.To("selected_card", Card.Type).
			Unique().
			Field("selected_card_id"),
	}
}
