// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yyewolf/entvis/_examples/users/internal/database/ent/card"
	"github.com/yyewolf/entvis/_examples/users/internal/database/ent/predicate"
	"github.com/yyewolf/entvis/_examples/users/internal/database/ent/user"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CardUpdate) SetUserID(s string) *CardUpdate {
	cu.mutation.SetUserID(s)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CardUpdate) SetNillableUserID(s *string) *CardUpdate {
	if s != nil {
		cu.SetUserID(*s)
	}
	return cu
}

// SetOwnedByID sets the "owned_by" edge to the User entity by ID.
func (cu *CardUpdate) SetOwnedByID(id string) *CardUpdate {
	cu.mutation.SetOwnedByID(id)
	return cu
}

// SetOwnedBy sets the "owned_by" edge to the User entity.
func (cu *CardUpdate) SetOwnedBy(u *User) *CardUpdate {
	return cu.SetOwnedByID(u.ID)
}

// AddFavedUserIDs adds the "faved_users" edge to the User entity by IDs.
func (cu *CardUpdate) AddFavedUserIDs(ids ...string) *CardUpdate {
	cu.mutation.AddFavedUserIDs(ids...)
	return cu
}

// AddFavedUsers adds the "faved_users" edges to the User entity.
func (cu *CardUpdate) AddFavedUsers(u ...*User) *CardUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddFavedUserIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cu *CardUpdate) Mutation() *CardMutation {
	return cu.mutation
}

// ClearOwnedBy clears the "owned_by" edge to the User entity.
func (cu *CardUpdate) ClearOwnedBy() *CardUpdate {
	cu.mutation.ClearOwnedBy()
	return cu
}

// ClearFavedUsers clears all "faved_users" edges to the User entity.
func (cu *CardUpdate) ClearFavedUsers() *CardUpdate {
	cu.mutation.ClearFavedUsers()
	return cu
}

// RemoveFavedUserIDs removes the "faved_users" edge to User entities by IDs.
func (cu *CardUpdate) RemoveFavedUserIDs(ids ...string) *CardUpdate {
	cu.mutation.RemoveFavedUserIDs(ids...)
	return cu
}

// RemoveFavedUsers removes "faved_users" edges to User entities.
func (cu *CardUpdate) RemoveFavedUsers(u ...*User) *CardUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveFavedUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CardUpdate) check() error {
	if cu.mutation.OwnedByCleared() && len(cu.mutation.OwnedByIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Card.owned_by"`)
	}
	return nil
}

func (cu *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.OwnedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnedByTable,
			Columns: []string{card.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnedByTable,
			Columns: []string{card.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.FavedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.FavedUsersTable,
			Columns: card.FavedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedFavedUsersIDs(); len(nodes) > 0 && !cu.mutation.FavedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.FavedUsersTable,
			Columns: card.FavedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FavedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.FavedUsersTable,
			Columns: card.FavedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetUserID sets the "user_id" field.
func (cuo *CardUpdateOne) SetUserID(s string) *CardUpdateOne {
	cuo.mutation.SetUserID(s)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableUserID(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetUserID(*s)
	}
	return cuo
}

// SetOwnedByID sets the "owned_by" edge to the User entity by ID.
func (cuo *CardUpdateOne) SetOwnedByID(id string) *CardUpdateOne {
	cuo.mutation.SetOwnedByID(id)
	return cuo
}

// SetOwnedBy sets the "owned_by" edge to the User entity.
func (cuo *CardUpdateOne) SetOwnedBy(u *User) *CardUpdateOne {
	return cuo.SetOwnedByID(u.ID)
}

// AddFavedUserIDs adds the "faved_users" edge to the User entity by IDs.
func (cuo *CardUpdateOne) AddFavedUserIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.AddFavedUserIDs(ids...)
	return cuo
}

// AddFavedUsers adds the "faved_users" edges to the User entity.
func (cuo *CardUpdateOne) AddFavedUsers(u ...*User) *CardUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddFavedUserIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cuo *CardUpdateOne) Mutation() *CardMutation {
	return cuo.mutation
}

// ClearOwnedBy clears the "owned_by" edge to the User entity.
func (cuo *CardUpdateOne) ClearOwnedBy() *CardUpdateOne {
	cuo.mutation.ClearOwnedBy()
	return cuo
}

// ClearFavedUsers clears all "faved_users" edges to the User entity.
func (cuo *CardUpdateOne) ClearFavedUsers() *CardUpdateOne {
	cuo.mutation.ClearFavedUsers()
	return cuo
}

// RemoveFavedUserIDs removes the "faved_users" edge to User entities by IDs.
func (cuo *CardUpdateOne) RemoveFavedUserIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.RemoveFavedUserIDs(ids...)
	return cuo
}

// RemoveFavedUsers removes "faved_users" edges to User entities.
func (cuo *CardUpdateOne) RemoveFavedUsers(u ...*User) *CardUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveFavedUserIDs(ids...)
}

// Where appends a list predicates to the CardUpdate builder.
func (cuo *CardUpdateOne) Where(ps ...predicate.Card) *CardUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Card entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CardUpdateOne) check() error {
	if cuo.mutation.OwnedByCleared() && len(cuo.mutation.OwnedByIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Card.owned_by"`)
	}
	return nil
}

func (cuo *CardUpdateOne) sqlSave(ctx context.Context) (_node *Card, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.OwnedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnedByTable,
			Columns: []string{card.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnedByTable,
			Columns: []string{card.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.FavedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.FavedUsersTable,
			Columns: card.FavedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedFavedUsersIDs(); len(nodes) > 0 && !cuo.mutation.FavedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.FavedUsersTable,
			Columns: card.FavedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FavedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.FavedUsersTable,
			Columns: card.FavedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Card{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
