// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"irg1008/next-go/ent/authrequest"
	"irg1008/next-go/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuthRequestCreate is the builder for creating a AuthRequest entity.
type AuthRequestCreate struct {
	config
	mutation *AuthRequestMutation
	hooks    []Hook
}

// SetActive sets the "active" field.
func (arc *AuthRequestCreate) SetActive(b bool) *AuthRequestCreate {
	arc.mutation.SetActive(b)
	return arc
}

// SetExpiresAt sets the "expires_at" field.
func (arc *AuthRequestCreate) SetExpiresAt(t time.Time) *AuthRequestCreate {
	arc.mutation.SetExpiresAt(t)
	return arc
}

// SetToken sets the "token" field.
func (arc *AuthRequestCreate) SetToken(u uuid.UUID) *AuthRequestCreate {
	arc.mutation.SetToken(u)
	return arc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (arc *AuthRequestCreate) SetNillableToken(u *uuid.UUID) *AuthRequestCreate {
	if u != nil {
		arc.SetToken(*u)
	}
	return arc
}

// SetType sets the "type" field.
func (arc *AuthRequestCreate) SetType(a authrequest.Type) *AuthRequestCreate {
	arc.mutation.SetType(a)
	return arc
}

// SetCreatedAt sets the "created_at" field.
func (arc *AuthRequestCreate) SetCreatedAt(t time.Time) *AuthRequestCreate {
	arc.mutation.SetCreatedAt(t)
	return arc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arc *AuthRequestCreate) SetNillableCreatedAt(t *time.Time) *AuthRequestCreate {
	if t != nil {
		arc.SetCreatedAt(*t)
	}
	return arc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (arc *AuthRequestCreate) SetUserID(id int) *AuthRequestCreate {
	arc.mutation.SetUserID(id)
	return arc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (arc *AuthRequestCreate) SetNillableUserID(id *int) *AuthRequestCreate {
	if id != nil {
		arc = arc.SetUserID(*id)
	}
	return arc
}

// SetUser sets the "user" edge to the User entity.
func (arc *AuthRequestCreate) SetUser(u *User) *AuthRequestCreate {
	return arc.SetUserID(u.ID)
}

// Mutation returns the AuthRequestMutation object of the builder.
func (arc *AuthRequestCreate) Mutation() *AuthRequestMutation {
	return arc.mutation
}

// Save creates the AuthRequest in the database.
func (arc *AuthRequestCreate) Save(ctx context.Context) (*AuthRequest, error) {
	arc.defaults()
	return withHooks(ctx, arc.sqlSave, arc.mutation, arc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AuthRequestCreate) SaveX(ctx context.Context) *AuthRequest {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AuthRequestCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AuthRequestCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AuthRequestCreate) defaults() {
	if _, ok := arc.mutation.Token(); !ok {
		v := authrequest.DefaultToken()
		arc.mutation.SetToken(v)
	}
	if _, ok := arc.mutation.CreatedAt(); !ok {
		v := authrequest.DefaultCreatedAt()
		arc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AuthRequestCreate) check() error {
	if _, ok := arc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "AuthRequest.active"`)}
	}
	if _, ok := arc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "AuthRequest.expires_at"`)}
	}
	if _, ok := arc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "AuthRequest.token"`)}
	}
	if _, ok := arc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "AuthRequest.type"`)}
	}
	if v, ok := arc.mutation.GetType(); ok {
		if err := authrequest.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AuthRequest.type": %w`, err)}
		}
	}
	if _, ok := arc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuthRequest.created_at"`)}
	}
	return nil
}

func (arc *AuthRequestCreate) sqlSave(ctx context.Context) (*AuthRequest, error) {
	if err := arc.check(); err != nil {
		return nil, err
	}
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	arc.mutation.id = &_node.ID
	arc.mutation.done = true
	return _node, nil
}

func (arc *AuthRequestCreate) createSpec() (*AuthRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthRequest{config: arc.config}
		_spec = sqlgraph.NewCreateSpec(authrequest.Table, sqlgraph.NewFieldSpec(authrequest.FieldID, field.TypeInt))
	)
	if value, ok := arc.mutation.Active(); ok {
		_spec.SetField(authrequest.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := arc.mutation.ExpiresAt(); ok {
		_spec.SetField(authrequest.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := arc.mutation.Token(); ok {
		_spec.SetField(authrequest.FieldToken, field.TypeUUID, value)
		_node.Token = value
	}
	if value, ok := arc.mutation.GetType(); ok {
		_spec.SetField(authrequest.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.SetField(authrequest.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := arc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_requests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AuthRequestCreateBulk is the builder for creating many AuthRequest entities in bulk.
type AuthRequestCreateBulk struct {
	config
	builders []*AuthRequestCreate
}

// Save creates the AuthRequest entities in the database.
func (arcb *AuthRequestCreateBulk) Save(ctx context.Context) ([]*AuthRequest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AuthRequest, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthRequestMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AuthRequestCreateBulk) SaveX(ctx context.Context) []*AuthRequest {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AuthRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AuthRequestCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}