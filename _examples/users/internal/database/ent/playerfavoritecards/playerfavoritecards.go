// Code generated by ent, DO NOT EDIT.

package playerfavoritecards

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the playerfavoritecards type in the database.
	Label = "player_favorite_cards"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCardID holds the string denoting the card_id field in the database.
	FieldCardID = "card_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// CardFieldID holds the string denoting the ID field of the Card.
	CardFieldID = "id"
	// Table holds the table name of the playerfavoritecards in the database.
	Table = "player_favorite_cards"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "player_favorite_cards"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CardTable is the table that holds the card relation/edge.
	CardTable = "player_favorite_cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "card_id"
)

// Columns holds all SQL columns for playerfavoritecards fields.
var Columns = []string{
	FieldUserID,
	FieldCardID,
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

// OrderOption defines the ordering options for the PlayerFavoriteCards queries.
type OrderOption func(*sql.Selector)

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCardID orders the results by the card_id field.
func ByCardID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCardField orders the results by card field.
func ByCardField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCardStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, UserColumn),
		sqlgraph.To(UserInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
func newCardStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, CardColumn),
		sqlgraph.To(CardInverseTable, CardFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CardTable, CardColumn),
	)
}
