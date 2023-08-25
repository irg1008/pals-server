// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"irg1008/pals/ent/authrequest"
	"irg1008/pals/ent/predicate"
	"irg1008/pals/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthRequestUpdate is the builder for updating AuthRequest entities.
type AuthRequestUpdate struct {
	config
	hooks    []Hook
	mutation *AuthRequestMutation
}

// Where appends a list predicates to the AuthRequestUpdate builder.
func (aru *AuthRequestUpdate) Where(ps ...predicate.AuthRequest) *AuthRequestUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetActive sets the "active" field.
func (aru *AuthRequestUpdate) SetActive(b bool) *AuthRequestUpdate {
	aru.mutation.SetActive(b)
	return aru
}

// SetExpiresAt sets the "expires_at" field.
func (aru *AuthRequestUpdate) SetExpiresAt(t time.Time) *AuthRequestUpdate {
	aru.mutation.SetExpiresAt(t)
	return aru
}

// SetType sets the "type" field.
func (aru *AuthRequestUpdate) SetType(a authrequest.Type) *AuthRequestUpdate {
	aru.mutation.SetType(a)
	return aru
}

// SetUserID sets the "user" edge to the User entity by ID.
func (aru *AuthRequestUpdate) SetUserID(id int) *AuthRequestUpdate {
	aru.mutation.SetUserID(id)
	return aru
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (aru *AuthRequestUpdate) SetNillableUserID(id *int) *AuthRequestUpdate {
	if id != nil {
		aru = aru.SetUserID(*id)
	}
	return aru
}

// SetUser sets the "user" edge to the User entity.
func (aru *AuthRequestUpdate) SetUser(u *User) *AuthRequestUpdate {
	return aru.SetUserID(u.ID)
}

// Mutation returns the AuthRequestMutation object of the builder.
func (aru *AuthRequestUpdate) Mutation() *AuthRequestMutation {
	return aru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (aru *AuthRequestUpdate) ClearUser() *AuthRequestUpdate {
	aru.mutation.ClearUser()
	return aru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AuthRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AuthRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AuthRequestUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AuthRequestUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aru *AuthRequestUpdate) check() error {
	if v, ok := aru.mutation.GetType(); ok {
		if err := authrequest.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AuthRequest.type": %w`, err)}
		}
	}
	return nil
}

func (aru *AuthRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(authrequest.Table, authrequest.Columns, sqlgraph.NewFieldSpec(authrequest.FieldID, field.TypeInt))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.Active(); ok {
		_spec.SetField(authrequest.FieldActive, field.TypeBool, value)
	}
	if value, ok := aru.mutation.ExpiresAt(); ok {
		_spec.SetField(authrequest.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := aru.mutation.GetType(); ok {
		_spec.SetField(authrequest.FieldType, field.TypeEnum, value)
	}
	if aru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authrequest.UserTable,
			Columns: []string{authrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authrequest.UserTable,
			Columns: []string{authrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// AuthRequestUpdateOne is the builder for updating a single AuthRequest entity.
type AuthRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthRequestMutation
}

// SetActive sets the "active" field.
func (aruo *AuthRequestUpdateOne) SetActive(b bool) *AuthRequestUpdateOne {
	aruo.mutation.SetActive(b)
	return aruo
}

// SetExpiresAt sets the "expires_at" field.
func (aruo *AuthRequestUpdateOne) SetExpiresAt(t time.Time) *AuthRequestUpdateOne {
	aruo.mutation.SetExpiresAt(t)
	return aruo
}

// SetType sets the "type" field.
func (aruo *AuthRequestUpdateOne) SetType(a authrequest.Type) *AuthRequestUpdateOne {
	aruo.mutation.SetType(a)
	return aruo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (aruo *AuthRequestUpdateOne) SetUserID(id int) *AuthRequestUpdateOne {
	aruo.mutation.SetUserID(id)
	return aruo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (aruo *AuthRequestUpdateOne) SetNillableUserID(id *int) *AuthRequestUpdateOne {
	if id != nil {
		aruo = aruo.SetUserID(*id)
	}
	return aruo
}

// SetUser sets the "user" edge to the User entity.
func (aruo *AuthRequestUpdateOne) SetUser(u *User) *AuthRequestUpdateOne {
	return aruo.SetUserID(u.ID)
}

// Mutation returns the AuthRequestMutation object of the builder.
func (aruo *AuthRequestUpdateOne) Mutation() *AuthRequestMutation {
	return aruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (aruo *AuthRequestUpdateOne) ClearUser() *AuthRequestUpdateOne {
	aruo.mutation.ClearUser()
	return aruo
}

// Where appends a list predicates to the AuthRequestUpdate builder.
func (aruo *AuthRequestUpdateOne) Where(ps ...predicate.AuthRequest) *AuthRequestUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AuthRequestUpdateOne) Select(field string, fields ...string) *AuthRequestUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AuthRequest entity.
func (aruo *AuthRequestUpdateOne) Save(ctx context.Context) (*AuthRequest, error) {
	return withHooks(ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AuthRequestUpdateOne) SaveX(ctx context.Context) *AuthRequest {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AuthRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AuthRequestUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aruo *AuthRequestUpdateOne) check() error {
	if v, ok := aruo.mutation.GetType(); ok {
		if err := authrequest.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AuthRequest.type": %w`, err)}
		}
	}
	return nil
}

func (aruo *AuthRequestUpdateOne) sqlSave(ctx context.Context) (_node *AuthRequest, err error) {
	if err := aruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(authrequest.Table, authrequest.Columns, sqlgraph.NewFieldSpec(authrequest.FieldID, field.TypeInt))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authrequest.FieldID)
		for _, f := range fields {
			if !authrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.Active(); ok {
		_spec.SetField(authrequest.FieldActive, field.TypeBool, value)
	}
	if value, ok := aruo.mutation.ExpiresAt(); ok {
		_spec.SetField(authrequest.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := aruo.mutation.GetType(); ok {
		_spec.SetField(authrequest.FieldType, field.TypeEnum, value)
	}
	if aruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authrequest.UserTable,
			Columns: []string{authrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authrequest.UserTable,
			Columns: []string{authrequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AuthRequest{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
