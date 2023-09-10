// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"irg1008/pals/ent/userdata"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserDataCreate is the builder for creating a UserData entity.
type UserDataCreate struct {
	config
	mutation *UserDataMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (udc *UserDataCreate) SetUserID(s string) *UserDataCreate {
	udc.mutation.SetUserID(s)
	return udc
}

// SetName sets the "name" field.
func (udc *UserDataCreate) SetName(s string) *UserDataCreate {
	udc.mutation.SetName(s)
	return udc
}

// SetEmail sets the "email" field.
func (udc *UserDataCreate) SetEmail(s string) *UserDataCreate {
	udc.mutation.SetEmail(s)
	return udc
}

// SetRole sets the "role" field.
func (udc *UserDataCreate) SetRole(u userdata.Role) *UserDataCreate {
	udc.mutation.SetRole(u)
	return udc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (udc *UserDataCreate) SetNillableRole(u *userdata.Role) *UserDataCreate {
	if u != nil {
		udc.SetRole(*u)
	}
	return udc
}

// SetCreatedAt sets the "created_at" field.
func (udc *UserDataCreate) SetCreatedAt(t time.Time) *UserDataCreate {
	udc.mutation.SetCreatedAt(t)
	return udc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (udc *UserDataCreate) SetNillableCreatedAt(t *time.Time) *UserDataCreate {
	if t != nil {
		udc.SetCreatedAt(*t)
	}
	return udc
}

// Mutation returns the UserDataMutation object of the builder.
func (udc *UserDataCreate) Mutation() *UserDataMutation {
	return udc.mutation
}

// Save creates the UserData in the database.
func (udc *UserDataCreate) Save(ctx context.Context) (*UserData, error) {
	udc.defaults()
	return withHooks(ctx, udc.sqlSave, udc.mutation, udc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (udc *UserDataCreate) SaveX(ctx context.Context) *UserData {
	v, err := udc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udc *UserDataCreate) Exec(ctx context.Context) error {
	_, err := udc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udc *UserDataCreate) ExecX(ctx context.Context) {
	if err := udc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (udc *UserDataCreate) defaults() {
	if _, ok := udc.mutation.Role(); !ok {
		v := userdata.DefaultRole
		udc.mutation.SetRole(v)
	}
	if _, ok := udc.mutation.CreatedAt(); !ok {
		v := userdata.DefaultCreatedAt()
		udc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (udc *UserDataCreate) check() error {
	if _, ok := udc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserData.user_id"`)}
	}
	if _, ok := udc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserData.name"`)}
	}
	if _, ok := udc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "UserData.email"`)}
	}
	if _, ok := udc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "UserData.role"`)}
	}
	if v, ok := udc.mutation.Role(); ok {
		if err := userdata.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "UserData.role": %w`, err)}
		}
	}
	if _, ok := udc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserData.created_at"`)}
	}
	return nil
}

func (udc *UserDataCreate) sqlSave(ctx context.Context) (*UserData, error) {
	if err := udc.check(); err != nil {
		return nil, err
	}
	_node, _spec := udc.createSpec()
	if err := sqlgraph.CreateNode(ctx, udc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	udc.mutation.id = &_node.ID
	udc.mutation.done = true
	return _node, nil
}

func (udc *UserDataCreate) createSpec() (*UserData, *sqlgraph.CreateSpec) {
	var (
		_node = &UserData{config: udc.config}
		_spec = sqlgraph.NewCreateSpec(userdata.Table, sqlgraph.NewFieldSpec(userdata.FieldID, field.TypeInt))
	)
	if value, ok := udc.mutation.UserID(); ok {
		_spec.SetField(userdata.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := udc.mutation.Name(); ok {
		_spec.SetField(userdata.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := udc.mutation.Email(); ok {
		_spec.SetField(userdata.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := udc.mutation.Role(); ok {
		_spec.SetField(userdata.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := udc.mutation.CreatedAt(); ok {
		_spec.SetField(userdata.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// UserDataCreateBulk is the builder for creating many UserData entities in bulk.
type UserDataCreateBulk struct {
	config
	builders []*UserDataCreate
}

// Save creates the UserData entities in the database.
func (udcb *UserDataCreateBulk) Save(ctx context.Context) ([]*UserData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(udcb.builders))
	nodes := make([]*UserData, len(udcb.builders))
	mutators := make([]Mutator, len(udcb.builders))
	for i := range udcb.builders {
		func(i int, root context.Context) {
			builder := udcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserDataMutation)
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
					_, err = mutators[i+1].Mutate(root, udcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, udcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, udcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (udcb *UserDataCreateBulk) SaveX(ctx context.Context) []*UserData {
	v, err := udcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udcb *UserDataCreateBulk) Exec(ctx context.Context) error {
	_, err := udcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udcb *UserDataCreateBulk) ExecX(ctx context.Context) {
	if err := udcb.Exec(ctx); err != nil {
		panic(err)
	}
}
