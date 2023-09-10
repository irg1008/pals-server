// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"irg1008/pals/ent/predicate"
	"irg1008/pals/ent/userdata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserDataUpdate is the builder for updating UserData entities.
type UserDataUpdate struct {
	config
	hooks    []Hook
	mutation *UserDataMutation
}

// Where appends a list predicates to the UserDataUpdate builder.
func (udu *UserDataUpdate) Where(ps ...predicate.UserData) *UserDataUpdate {
	udu.mutation.Where(ps...)
	return udu
}

// SetName sets the "name" field.
func (udu *UserDataUpdate) SetName(s string) *UserDataUpdate {
	udu.mutation.SetName(s)
	return udu
}

// SetEmail sets the "email" field.
func (udu *UserDataUpdate) SetEmail(s string) *UserDataUpdate {
	udu.mutation.SetEmail(s)
	return udu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (udu *UserDataUpdate) SetNillableEmail(s *string) *UserDataUpdate {
	if s != nil {
		udu.SetEmail(*s)
	}
	return udu
}

// ClearEmail clears the value of the "email" field.
func (udu *UserDataUpdate) ClearEmail() *UserDataUpdate {
	udu.mutation.ClearEmail()
	return udu
}

// SetPicture sets the "picture" field.
func (udu *UserDataUpdate) SetPicture(s string) *UserDataUpdate {
	udu.mutation.SetPicture(s)
	return udu
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (udu *UserDataUpdate) SetNillablePicture(s *string) *UserDataUpdate {
	if s != nil {
		udu.SetPicture(*s)
	}
	return udu
}

// ClearPicture clears the value of the "picture" field.
func (udu *UserDataUpdate) ClearPicture() *UserDataUpdate {
	udu.mutation.ClearPicture()
	return udu
}

// SetRole sets the "role" field.
func (udu *UserDataUpdate) SetRole(u userdata.Role) *UserDataUpdate {
	udu.mutation.SetRole(u)
	return udu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (udu *UserDataUpdate) SetNillableRole(u *userdata.Role) *UserDataUpdate {
	if u != nil {
		udu.SetRole(*u)
	}
	return udu
}

// Mutation returns the UserDataMutation object of the builder.
func (udu *UserDataUpdate) Mutation() *UserDataMutation {
	return udu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (udu *UserDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, udu.sqlSave, udu.mutation, udu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (udu *UserDataUpdate) SaveX(ctx context.Context) int {
	affected, err := udu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (udu *UserDataUpdate) Exec(ctx context.Context) error {
	_, err := udu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udu *UserDataUpdate) ExecX(ctx context.Context) {
	if err := udu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (udu *UserDataUpdate) check() error {
	if v, ok := udu.mutation.Role(); ok {
		if err := userdata.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "UserData.role": %w`, err)}
		}
	}
	return nil
}

func (udu *UserDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := udu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userdata.Table, userdata.Columns, sqlgraph.NewFieldSpec(userdata.FieldID, field.TypeInt))
	if ps := udu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := udu.mutation.Name(); ok {
		_spec.SetField(userdata.FieldName, field.TypeString, value)
	}
	if value, ok := udu.mutation.Email(); ok {
		_spec.SetField(userdata.FieldEmail, field.TypeString, value)
	}
	if udu.mutation.EmailCleared() {
		_spec.ClearField(userdata.FieldEmail, field.TypeString)
	}
	if value, ok := udu.mutation.Picture(); ok {
		_spec.SetField(userdata.FieldPicture, field.TypeString, value)
	}
	if udu.mutation.PictureCleared() {
		_spec.ClearField(userdata.FieldPicture, field.TypeString)
	}
	if value, ok := udu.mutation.Role(); ok {
		_spec.SetField(userdata.FieldRole, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, udu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	udu.mutation.done = true
	return n, nil
}

// UserDataUpdateOne is the builder for updating a single UserData entity.
type UserDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserDataMutation
}

// SetName sets the "name" field.
func (uduo *UserDataUpdateOne) SetName(s string) *UserDataUpdateOne {
	uduo.mutation.SetName(s)
	return uduo
}

// SetEmail sets the "email" field.
func (uduo *UserDataUpdateOne) SetEmail(s string) *UserDataUpdateOne {
	uduo.mutation.SetEmail(s)
	return uduo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uduo *UserDataUpdateOne) SetNillableEmail(s *string) *UserDataUpdateOne {
	if s != nil {
		uduo.SetEmail(*s)
	}
	return uduo
}

// ClearEmail clears the value of the "email" field.
func (uduo *UserDataUpdateOne) ClearEmail() *UserDataUpdateOne {
	uduo.mutation.ClearEmail()
	return uduo
}

// SetPicture sets the "picture" field.
func (uduo *UserDataUpdateOne) SetPicture(s string) *UserDataUpdateOne {
	uduo.mutation.SetPicture(s)
	return uduo
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (uduo *UserDataUpdateOne) SetNillablePicture(s *string) *UserDataUpdateOne {
	if s != nil {
		uduo.SetPicture(*s)
	}
	return uduo
}

// ClearPicture clears the value of the "picture" field.
func (uduo *UserDataUpdateOne) ClearPicture() *UserDataUpdateOne {
	uduo.mutation.ClearPicture()
	return uduo
}

// SetRole sets the "role" field.
func (uduo *UserDataUpdateOne) SetRole(u userdata.Role) *UserDataUpdateOne {
	uduo.mutation.SetRole(u)
	return uduo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uduo *UserDataUpdateOne) SetNillableRole(u *userdata.Role) *UserDataUpdateOne {
	if u != nil {
		uduo.SetRole(*u)
	}
	return uduo
}

// Mutation returns the UserDataMutation object of the builder.
func (uduo *UserDataUpdateOne) Mutation() *UserDataMutation {
	return uduo.mutation
}

// Where appends a list predicates to the UserDataUpdate builder.
func (uduo *UserDataUpdateOne) Where(ps ...predicate.UserData) *UserDataUpdateOne {
	uduo.mutation.Where(ps...)
	return uduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uduo *UserDataUpdateOne) Select(field string, fields ...string) *UserDataUpdateOne {
	uduo.fields = append([]string{field}, fields...)
	return uduo
}

// Save executes the query and returns the updated UserData entity.
func (uduo *UserDataUpdateOne) Save(ctx context.Context) (*UserData, error) {
	return withHooks(ctx, uduo.sqlSave, uduo.mutation, uduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uduo *UserDataUpdateOne) SaveX(ctx context.Context) *UserData {
	node, err := uduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uduo *UserDataUpdateOne) Exec(ctx context.Context) error {
	_, err := uduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uduo *UserDataUpdateOne) ExecX(ctx context.Context) {
	if err := uduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uduo *UserDataUpdateOne) check() error {
	if v, ok := uduo.mutation.Role(); ok {
		if err := userdata.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "UserData.role": %w`, err)}
		}
	}
	return nil
}

func (uduo *UserDataUpdateOne) sqlSave(ctx context.Context) (_node *UserData, err error) {
	if err := uduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userdata.Table, userdata.Columns, sqlgraph.NewFieldSpec(userdata.FieldID, field.TypeInt))
	id, ok := uduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userdata.FieldID)
		for _, f := range fields {
			if !userdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uduo.mutation.Name(); ok {
		_spec.SetField(userdata.FieldName, field.TypeString, value)
	}
	if value, ok := uduo.mutation.Email(); ok {
		_spec.SetField(userdata.FieldEmail, field.TypeString, value)
	}
	if uduo.mutation.EmailCleared() {
		_spec.ClearField(userdata.FieldEmail, field.TypeString)
	}
	if value, ok := uduo.mutation.Picture(); ok {
		_spec.SetField(userdata.FieldPicture, field.TypeString, value)
	}
	if uduo.mutation.PictureCleared() {
		_spec.ClearField(userdata.FieldPicture, field.TypeString)
	}
	if value, ok := uduo.mutation.Role(); ok {
		_spec.SetField(userdata.FieldRole, field.TypeEnum, value)
	}
	_node = &UserData{config: uduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uduo.mutation.done = true
	return _node, nil
}