// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yyewolf/entvis/_examples/users/internal/database/ent/card"
	"github.com/yyewolf/entvis/_examples/users/internal/database/ent/playerfavoritecards"
	"github.com/yyewolf/entvis/_examples/users/internal/database/ent/user"
)

// PlayerFavoriteCardsCreate is the builder for creating a PlayerFavoriteCards entity.
type PlayerFavoriteCardsCreate struct {
	config
	mutation *PlayerFavoriteCardsMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (pfcc *PlayerFavoriteCardsCreate) SetUserID(s string) *PlayerFavoriteCardsCreate {
	pfcc.mutation.SetUserID(s)
	return pfcc
}

// SetCardID sets the "card_id" field.
func (pfcc *PlayerFavoriteCardsCreate) SetCardID(s string) *PlayerFavoriteCardsCreate {
	pfcc.mutation.SetCardID(s)
	return pfcc
}

// SetUser sets the "user" edge to the User entity.
func (pfcc *PlayerFavoriteCardsCreate) SetUser(u *User) *PlayerFavoriteCardsCreate {
	return pfcc.SetUserID(u.ID)
}

// SetCard sets the "card" edge to the Card entity.
func (pfcc *PlayerFavoriteCardsCreate) SetCard(c *Card) *PlayerFavoriteCardsCreate {
	return pfcc.SetCardID(c.ID)
}

// Mutation returns the PlayerFavoriteCardsMutation object of the builder.
func (pfcc *PlayerFavoriteCardsCreate) Mutation() *PlayerFavoriteCardsMutation {
	return pfcc.mutation
}

// Save creates the PlayerFavoriteCards in the database.
func (pfcc *PlayerFavoriteCardsCreate) Save(ctx context.Context) (*PlayerFavoriteCards, error) {
	return withHooks(ctx, pfcc.sqlSave, pfcc.mutation, pfcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pfcc *PlayerFavoriteCardsCreate) SaveX(ctx context.Context) *PlayerFavoriteCards {
	v, err := pfcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfcc *PlayerFavoriteCardsCreate) Exec(ctx context.Context) error {
	_, err := pfcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfcc *PlayerFavoriteCardsCreate) ExecX(ctx context.Context) {
	if err := pfcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfcc *PlayerFavoriteCardsCreate) check() error {
	if _, ok := pfcc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "PlayerFavoriteCards.user_id"`)}
	}
	if _, ok := pfcc.mutation.CardID(); !ok {
		return &ValidationError{Name: "card_id", err: errors.New(`ent: missing required field "PlayerFavoriteCards.card_id"`)}
	}
	if len(pfcc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "PlayerFavoriteCards.user"`)}
	}
	if len(pfcc.mutation.CardIDs()) == 0 {
		return &ValidationError{Name: "card", err: errors.New(`ent: missing required edge "PlayerFavoriteCards.card"`)}
	}
	return nil
}

func (pfcc *PlayerFavoriteCardsCreate) sqlSave(ctx context.Context) (*PlayerFavoriteCards, error) {
	if err := pfcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pfcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pfcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (pfcc *PlayerFavoriteCardsCreate) createSpec() (*PlayerFavoriteCards, *sqlgraph.CreateSpec) {
	var (
		_node = &PlayerFavoriteCards{config: pfcc.config}
		_spec = sqlgraph.NewCreateSpec(playerfavoritecards.Table, nil)
	)
	if nodes := pfcc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playerfavoritecards.UserTable,
			Columns: []string{playerfavoritecards.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfcc.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   playerfavoritecards.CardTable,
			Columns: []string{playerfavoritecards.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CardID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlayerFavoriteCardsCreateBulk is the builder for creating many PlayerFavoriteCards entities in bulk.
type PlayerFavoriteCardsCreateBulk struct {
	config
	err      error
	builders []*PlayerFavoriteCardsCreate
}

// Save creates the PlayerFavoriteCards entities in the database.
func (pfccb *PlayerFavoriteCardsCreateBulk) Save(ctx context.Context) ([]*PlayerFavoriteCards, error) {
	if pfccb.err != nil {
		return nil, pfccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pfccb.builders))
	nodes := make([]*PlayerFavoriteCards, len(pfccb.builders))
	mutators := make([]Mutator, len(pfccb.builders))
	for i := range pfccb.builders {
		func(i int, root context.Context) {
			builder := pfccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerFavoriteCardsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pfccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pfccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pfccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pfccb *PlayerFavoriteCardsCreateBulk) SaveX(ctx context.Context) []*PlayerFavoriteCards {
	v, err := pfccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfccb *PlayerFavoriteCardsCreateBulk) Exec(ctx context.Context) error {
	_, err := pfccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfccb *PlayerFavoriteCardsCreateBulk) ExecX(ctx context.Context) {
	if err := pfccb.Exec(ctx); err != nil {
		panic(err)
	}
}
