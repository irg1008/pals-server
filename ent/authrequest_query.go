// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"irg1008/pals/ent/authrequest"
	"irg1008/pals/ent/predicate"
	"irg1008/pals/ent/user"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthRequestQuery is the builder for querying AuthRequest entities.
type AuthRequestQuery struct {
	config
	ctx        *QueryContext
	order      []authrequest.OrderOption
	inters     []Interceptor
	predicates []predicate.AuthRequest
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthRequestQuery builder.
func (arq *AuthRequestQuery) Where(ps ...predicate.AuthRequest) *AuthRequestQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit the number of records to be returned by this query.
func (arq *AuthRequestQuery) Limit(limit int) *AuthRequestQuery {
	arq.ctx.Limit = &limit
	return arq
}

// Offset to start from.
func (arq *AuthRequestQuery) Offset(offset int) *AuthRequestQuery {
	arq.ctx.Offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AuthRequestQuery) Unique(unique bool) *AuthRequestQuery {
	arq.ctx.Unique = &unique
	return arq
}

// Order specifies how the records should be ordered.
func (arq *AuthRequestQuery) Order(o ...authrequest.OrderOption) *AuthRequestQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryUser chains the current query on the "user" edge.
func (arq *AuthRequestQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: arq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authrequest.Table, authrequest.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authrequest.UserTable, authrequest.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthRequest entity from the query.
// Returns a *NotFoundError when no AuthRequest was found.
func (arq *AuthRequestQuery) First(ctx context.Context) (*AuthRequest, error) {
	nodes, err := arq.Limit(1).All(setContextOp(ctx, arq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authrequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AuthRequestQuery) FirstX(ctx context.Context) *AuthRequest {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthRequest ID from the query.
// Returns a *NotFoundError when no AuthRequest ID was found.
func (arq *AuthRequestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(1).IDs(setContextOp(ctx, arq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authrequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AuthRequestQuery) FirstIDX(ctx context.Context) int {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthRequest entity is found.
// Returns a *NotFoundError when no AuthRequest entities are found.
func (arq *AuthRequestQuery) Only(ctx context.Context) (*AuthRequest, error) {
	nodes, err := arq.Limit(2).All(setContextOp(ctx, arq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authrequest.Label}
	default:
		return nil, &NotSingularError{authrequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AuthRequestQuery) OnlyX(ctx context.Context) *AuthRequest {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthRequest ID in the query.
// Returns a *NotSingularError when more than one AuthRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AuthRequestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(2).IDs(setContextOp(ctx, arq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authrequest.Label}
	default:
		err = &NotSingularError{authrequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AuthRequestQuery) OnlyIDX(ctx context.Context) int {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthRequests.
func (arq *AuthRequestQuery) All(ctx context.Context) ([]*AuthRequest, error) {
	ctx = setContextOp(ctx, arq.ctx, "All")
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AuthRequest, *AuthRequestQuery]()
	return withInterceptors[[]*AuthRequest](ctx, arq, qr, arq.inters)
}

// AllX is like All, but panics if an error occurs.
func (arq *AuthRequestQuery) AllX(ctx context.Context) []*AuthRequest {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthRequest IDs.
func (arq *AuthRequestQuery) IDs(ctx context.Context) (ids []int, err error) {
	if arq.ctx.Unique == nil && arq.path != nil {
		arq.Unique(true)
	}
	ctx = setContextOp(ctx, arq.ctx, "IDs")
	if err = arq.Select(authrequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AuthRequestQuery) IDsX(ctx context.Context) []int {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AuthRequestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, arq.ctx, "Count")
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, arq, querierCount[*AuthRequestQuery](), arq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AuthRequestQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AuthRequestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, arq.ctx, "Exist")
	switch _, err := arq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AuthRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AuthRequestQuery) Clone() *AuthRequestQuery {
	if arq == nil {
		return nil
	}
	return &AuthRequestQuery{
		config:     arq.config,
		ctx:        arq.ctx.Clone(),
		order:      append([]authrequest.OrderOption{}, arq.order...),
		inters:     append([]Interceptor{}, arq.inters...),
		predicates: append([]predicate.AuthRequest{}, arq.predicates...),
		withUser:   arq.withUser.Clone(),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AuthRequestQuery) WithUser(opts ...func(*UserQuery)) *AuthRequestQuery {
	query := (&UserClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	arq.withUser = query
	return arq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Active bool `json:"active,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthRequest.Query().
//		GroupBy(authrequest.FieldActive).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (arq *AuthRequestQuery) GroupBy(field string, fields ...string) *AuthRequestGroupBy {
	arq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AuthRequestGroupBy{build: arq}
	grbuild.flds = &arq.ctx.Fields
	grbuild.label = authrequest.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Active bool `json:"active,omitempty"`
//	}
//
//	client.AuthRequest.Query().
//		Select(authrequest.FieldActive).
//		Scan(ctx, &v)
func (arq *AuthRequestQuery) Select(fields ...string) *AuthRequestSelect {
	arq.ctx.Fields = append(arq.ctx.Fields, fields...)
	sbuild := &AuthRequestSelect{AuthRequestQuery: arq}
	sbuild.label = authrequest.Label
	sbuild.flds, sbuild.scan = &arq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AuthRequestSelect configured with the given aggregations.
func (arq *AuthRequestQuery) Aggregate(fns ...AggregateFunc) *AuthRequestSelect {
	return arq.Select().Aggregate(fns...)
}

func (arq *AuthRequestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range arq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, arq); err != nil {
				return err
			}
		}
	}
	for _, f := range arq.ctx.Fields {
		if !authrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AuthRequestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthRequest, error) {
	var (
		nodes       = []*AuthRequest{}
		withFKs     = arq.withFKs
		_spec       = arq.querySpec()
		loadedTypes = [1]bool{
			arq.withUser != nil,
		}
	)
	if arq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, authrequest.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AuthRequest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AuthRequest{config: arq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := arq.withUser; query != nil {
		if err := arq.loadUser(ctx, query, nodes, nil,
			func(n *AuthRequest, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (arq *AuthRequestQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*AuthRequest, init func(*AuthRequest), assign func(*AuthRequest, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AuthRequest)
	for i := range nodes {
		if nodes[i].user_requests == nil {
			continue
		}
		fk := *nodes[i].user_requests
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_requests" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (arq *AuthRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	_spec.Node.Columns = arq.ctx.Fields
	if len(arq.ctx.Fields) > 0 {
		_spec.Unique = arq.ctx.Unique != nil && *arq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AuthRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(authrequest.Table, authrequest.Columns, sqlgraph.NewFieldSpec(authrequest.FieldID, field.TypeInt))
	_spec.From = arq.sql
	if unique := arq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if arq.path != nil {
		_spec.Unique = true
	}
	if fields := arq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authrequest.FieldID)
		for i := range fields {
			if fields[i] != authrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AuthRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(authrequest.Table)
	columns := arq.ctx.Fields
	if len(columns) == 0 {
		columns = authrequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.ctx.Unique != nil && *arq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthRequestGroupBy is the group-by builder for AuthRequest entities.
type AuthRequestGroupBy struct {
	selector
	build *AuthRequestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AuthRequestGroupBy) Aggregate(fns ...AggregateFunc) *AuthRequestGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the selector query and scans the result into the given value.
func (argb *AuthRequestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, argb.build.ctx, "GroupBy")
	if err := argb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthRequestQuery, *AuthRequestGroupBy](ctx, argb.build, argb, argb.build.inters, v)
}

func (argb *AuthRequestGroupBy) sqlScan(ctx context.Context, root *AuthRequestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*argb.flds)+len(argb.fns))
		for _, f := range *argb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*argb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AuthRequestSelect is the builder for selecting fields of AuthRequest entities.
type AuthRequestSelect struct {
	*AuthRequestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ars *AuthRequestSelect) Aggregate(fns ...AggregateFunc) *AuthRequestSelect {
	ars.fns = append(ars.fns, fns...)
	return ars
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AuthRequestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ars.ctx, "Select")
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthRequestQuery, *AuthRequestSelect](ctx, ars.AuthRequestQuery, ars, ars.inters, v)
}

func (ars *AuthRequestSelect) sqlScan(ctx context.Context, root *AuthRequestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ars.fns))
	for _, fn := range ars.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ars.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
