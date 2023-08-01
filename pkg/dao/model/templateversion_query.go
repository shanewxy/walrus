// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/service"
	"github.com/seal-io/seal/pkg/dao/model/template"
	"github.com/seal-io/seal/pkg/dao/model/templateversion"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// TemplateVersionQuery is the builder for querying TemplateVersion entities.
type TemplateVersionQuery struct {
	config
	ctx          *QueryContext
	order        []templateversion.OrderOption
	inters       []Interceptor
	predicates   []predicate.TemplateVersion
	withTemplate *TemplateQuery
	withServices *ServiceQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TemplateVersionQuery builder.
func (tvq *TemplateVersionQuery) Where(ps ...predicate.TemplateVersion) *TemplateVersionQuery {
	tvq.predicates = append(tvq.predicates, ps...)
	return tvq
}

// Limit the number of records to be returned by this query.
func (tvq *TemplateVersionQuery) Limit(limit int) *TemplateVersionQuery {
	tvq.ctx.Limit = &limit
	return tvq
}

// Offset to start from.
func (tvq *TemplateVersionQuery) Offset(offset int) *TemplateVersionQuery {
	tvq.ctx.Offset = &offset
	return tvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tvq *TemplateVersionQuery) Unique(unique bool) *TemplateVersionQuery {
	tvq.ctx.Unique = &unique
	return tvq
}

// Order specifies how the records should be ordered.
func (tvq *TemplateVersionQuery) Order(o ...templateversion.OrderOption) *TemplateVersionQuery {
	tvq.order = append(tvq.order, o...)
	return tvq
}

// QueryTemplate chains the current query on the "template" edge.
func (tvq *TemplateVersionQuery) QueryTemplate() *TemplateQuery {
	query := (&TemplateClient{config: tvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(templateversion.Table, templateversion.FieldID, selector),
			sqlgraph.To(template.Table, template.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, templateversion.TemplateTable, templateversion.TemplateColumn),
		)
		schemaConfig := tvq.schemaConfig
		step.To.Schema = schemaConfig.Template
		step.Edge.Schema = schemaConfig.TemplateVersion
		fromU = sqlgraph.SetNeighbors(tvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryServices chains the current query on the "services" edge.
func (tvq *TemplateVersionQuery) QueryServices() *ServiceQuery {
	query := (&ServiceClient{config: tvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(templateversion.Table, templateversion.FieldID, selector),
			sqlgraph.To(service.Table, service.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, templateversion.ServicesTable, templateversion.ServicesColumn),
		)
		schemaConfig := tvq.schemaConfig
		step.To.Schema = schemaConfig.Service
		step.Edge.Schema = schemaConfig.Service
		fromU = sqlgraph.SetNeighbors(tvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TemplateVersion entity from the query.
// Returns a *NotFoundError when no TemplateVersion was found.
func (tvq *TemplateVersionQuery) First(ctx context.Context) (*TemplateVersion, error) {
	nodes, err := tvq.Limit(1).All(setContextOp(ctx, tvq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{templateversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tvq *TemplateVersionQuery) FirstX(ctx context.Context) *TemplateVersion {
	node, err := tvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TemplateVersion ID from the query.
// Returns a *NotFoundError when no TemplateVersion ID was found.
func (tvq *TemplateVersionQuery) FirstID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = tvq.Limit(1).IDs(setContextOp(ctx, tvq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{templateversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tvq *TemplateVersionQuery) FirstIDX(ctx context.Context) object.ID {
	id, err := tvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TemplateVersion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TemplateVersion entity is found.
// Returns a *NotFoundError when no TemplateVersion entities are found.
func (tvq *TemplateVersionQuery) Only(ctx context.Context) (*TemplateVersion, error) {
	nodes, err := tvq.Limit(2).All(setContextOp(ctx, tvq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{templateversion.Label}
	default:
		return nil, &NotSingularError{templateversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tvq *TemplateVersionQuery) OnlyX(ctx context.Context) *TemplateVersion {
	node, err := tvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TemplateVersion ID in the query.
// Returns a *NotSingularError when more than one TemplateVersion ID is found.
// Returns a *NotFoundError when no entities are found.
func (tvq *TemplateVersionQuery) OnlyID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = tvq.Limit(2).IDs(setContextOp(ctx, tvq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{templateversion.Label}
	default:
		err = &NotSingularError{templateversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tvq *TemplateVersionQuery) OnlyIDX(ctx context.Context) object.ID {
	id, err := tvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TemplateVersions.
func (tvq *TemplateVersionQuery) All(ctx context.Context) ([]*TemplateVersion, error) {
	ctx = setContextOp(ctx, tvq.ctx, "All")
	if err := tvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TemplateVersion, *TemplateVersionQuery]()
	return withInterceptors[[]*TemplateVersion](ctx, tvq, qr, tvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tvq *TemplateVersionQuery) AllX(ctx context.Context) []*TemplateVersion {
	nodes, err := tvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TemplateVersion IDs.
func (tvq *TemplateVersionQuery) IDs(ctx context.Context) (ids []object.ID, err error) {
	if tvq.ctx.Unique == nil && tvq.path != nil {
		tvq.Unique(true)
	}
	ctx = setContextOp(ctx, tvq.ctx, "IDs")
	if err = tvq.Select(templateversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tvq *TemplateVersionQuery) IDsX(ctx context.Context) []object.ID {
	ids, err := tvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tvq *TemplateVersionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tvq.ctx, "Count")
	if err := tvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tvq, querierCount[*TemplateVersionQuery](), tvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tvq *TemplateVersionQuery) CountX(ctx context.Context) int {
	count, err := tvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tvq *TemplateVersionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tvq.ctx, "Exist")
	switch _, err := tvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tvq *TemplateVersionQuery) ExistX(ctx context.Context) bool {
	exist, err := tvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TemplateVersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tvq *TemplateVersionQuery) Clone() *TemplateVersionQuery {
	if tvq == nil {
		return nil
	}
	return &TemplateVersionQuery{
		config:       tvq.config,
		ctx:          tvq.ctx.Clone(),
		order:        append([]templateversion.OrderOption{}, tvq.order...),
		inters:       append([]Interceptor{}, tvq.inters...),
		predicates:   append([]predicate.TemplateVersion{}, tvq.predicates...),
		withTemplate: tvq.withTemplate.Clone(),
		withServices: tvq.withServices.Clone(),
		// clone intermediate query.
		sql:  tvq.sql.Clone(),
		path: tvq.path,
	}
}

// WithTemplate tells the query-builder to eager-load the nodes that are connected to
// the "template" edge. The optional arguments are used to configure the query builder of the edge.
func (tvq *TemplateVersionQuery) WithTemplate(opts ...func(*TemplateQuery)) *TemplateVersionQuery {
	query := (&TemplateClient{config: tvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tvq.withTemplate = query
	return tvq
}

// WithServices tells the query-builder to eager-load the nodes that are connected to
// the "services" edge. The optional arguments are used to configure the query builder of the edge.
func (tvq *TemplateVersionQuery) WithServices(opts ...func(*ServiceQuery)) *TemplateVersionQuery {
	query := (&ServiceClient{config: tvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tvq.withServices = query
	return tvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TemplateVersion.Query().
//		GroupBy(templateversion.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (tvq *TemplateVersionQuery) GroupBy(field string, fields ...string) *TemplateVersionGroupBy {
	tvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TemplateVersionGroupBy{build: tvq}
	grbuild.flds = &tvq.ctx.Fields
	grbuild.label = templateversion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.TemplateVersion.Query().
//		Select(templateversion.FieldCreateTime).
//		Scan(ctx, &v)
func (tvq *TemplateVersionQuery) Select(fields ...string) *TemplateVersionSelect {
	tvq.ctx.Fields = append(tvq.ctx.Fields, fields...)
	sbuild := &TemplateVersionSelect{TemplateVersionQuery: tvq}
	sbuild.label = templateversion.Label
	sbuild.flds, sbuild.scan = &tvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TemplateVersionSelect configured with the given aggregations.
func (tvq *TemplateVersionQuery) Aggregate(fns ...AggregateFunc) *TemplateVersionSelect {
	return tvq.Select().Aggregate(fns...)
}

func (tvq *TemplateVersionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tvq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tvq); err != nil {
				return err
			}
		}
	}
	for _, f := range tvq.ctx.Fields {
		if !templateversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if tvq.path != nil {
		prev, err := tvq.path(ctx)
		if err != nil {
			return err
		}
		tvq.sql = prev
	}
	return nil
}

func (tvq *TemplateVersionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TemplateVersion, error) {
	var (
		nodes       = []*TemplateVersion{}
		_spec       = tvq.querySpec()
		loadedTypes = [2]bool{
			tvq.withTemplate != nil,
			tvq.withServices != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TemplateVersion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TemplateVersion{config: tvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tvq.schemaConfig.TemplateVersion
	ctx = internal.NewSchemaConfigContext(ctx, tvq.schemaConfig)
	if len(tvq.modifiers) > 0 {
		_spec.Modifiers = tvq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tvq.withTemplate; query != nil {
		if err := tvq.loadTemplate(ctx, query, nodes, nil,
			func(n *TemplateVersion, e *Template) { n.Edges.Template = e }); err != nil {
			return nil, err
		}
	}
	if query := tvq.withServices; query != nil {
		if err := tvq.loadServices(ctx, query, nodes,
			func(n *TemplateVersion) { n.Edges.Services = []*Service{} },
			func(n *TemplateVersion, e *Service) { n.Edges.Services = append(n.Edges.Services, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tvq *TemplateVersionQuery) loadTemplate(ctx context.Context, query *TemplateQuery, nodes []*TemplateVersion, init func(*TemplateVersion), assign func(*TemplateVersion, *Template)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*TemplateVersion)
	for i := range nodes {
		fk := nodes[i].TemplateID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(template.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "template_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tvq *TemplateVersionQuery) loadServices(ctx context.Context, query *ServiceQuery, nodes []*TemplateVersion, init func(*TemplateVersion), assign func(*TemplateVersion, *Service)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[object.ID]*TemplateVersion)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(service.FieldTemplateID)
	}
	query.Where(predicate.Service(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(templateversion.ServicesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TemplateID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "template_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tvq *TemplateVersionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tvq.querySpec()
	_spec.Node.Schema = tvq.schemaConfig.TemplateVersion
	ctx = internal.NewSchemaConfigContext(ctx, tvq.schemaConfig)
	if len(tvq.modifiers) > 0 {
		_spec.Modifiers = tvq.modifiers
	}
	_spec.Node.Columns = tvq.ctx.Fields
	if len(tvq.ctx.Fields) > 0 {
		_spec.Unique = tvq.ctx.Unique != nil && *tvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tvq.driver, _spec)
}

func (tvq *TemplateVersionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(templateversion.Table, templateversion.Columns, sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString))
	_spec.From = tvq.sql
	if unique := tvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tvq.path != nil {
		_spec.Unique = true
	}
	if fields := tvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, templateversion.FieldID)
		for i := range fields {
			if fields[i] != templateversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tvq.withTemplate != nil {
			_spec.Node.AddColumnOnce(templateversion.FieldTemplateID)
		}
	}
	if ps := tvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tvq *TemplateVersionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tvq.driver.Dialect())
	t1 := builder.Table(templateversion.Table)
	columns := tvq.ctx.Fields
	if len(columns) == 0 {
		columns = templateversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tvq.sql != nil {
		selector = tvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tvq.ctx.Unique != nil && *tvq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tvq.schemaConfig.TemplateVersion)
	ctx = internal.NewSchemaConfigContext(ctx, tvq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range tvq.modifiers {
		m(selector)
	}
	for _, p := range tvq.predicates {
		p(selector)
	}
	for _, p := range tvq.order {
		p(selector)
	}
	if offset := tvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tvq *TemplateVersionQuery) ForUpdate(opts ...sql.LockOption) *TemplateVersionQuery {
	if tvq.driver.Dialect() == dialect.Postgres {
		tvq.Unique(false)
	}
	tvq.modifiers = append(tvq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tvq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tvq *TemplateVersionQuery) ForShare(opts ...sql.LockOption) *TemplateVersionQuery {
	if tvq.driver.Dialect() == dialect.Postgres {
		tvq.Unique(false)
	}
	tvq.modifiers = append(tvq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tvq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tvq *TemplateVersionQuery) Modify(modifiers ...func(s *sql.Selector)) *TemplateVersionSelect {
	tvq.modifiers = append(tvq.modifiers, modifiers...)
	return tvq.Select()
}

// WhereP appends storage-level predicates to the TemplateVersionQuery builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (tvq *TemplateVersionQuery) WhereP(ps ...func(*sql.Selector)) {
	var wps = make([]predicate.TemplateVersion, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		wps = append(wps, predicate.TemplateVersion(ps[i]))
	}
	tvq.predicates = append(tvq.predicates, wps...)
}

// TemplateVersionGroupBy is the group-by builder for TemplateVersion entities.
type TemplateVersionGroupBy struct {
	selector
	build *TemplateVersionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tvgb *TemplateVersionGroupBy) Aggregate(fns ...AggregateFunc) *TemplateVersionGroupBy {
	tvgb.fns = append(tvgb.fns, fns...)
	return tvgb
}

// Scan applies the selector query and scans the result into the given value.
func (tvgb *TemplateVersionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tvgb.build.ctx, "GroupBy")
	if err := tvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TemplateVersionQuery, *TemplateVersionGroupBy](ctx, tvgb.build, tvgb, tvgb.build.inters, v)
}

func (tvgb *TemplateVersionGroupBy) sqlScan(ctx context.Context, root *TemplateVersionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tvgb.fns))
	for _, fn := range tvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tvgb.flds)+len(tvgb.fns))
		for _, f := range *tvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TemplateVersionSelect is the builder for selecting fields of TemplateVersion entities.
type TemplateVersionSelect struct {
	*TemplateVersionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tvs *TemplateVersionSelect) Aggregate(fns ...AggregateFunc) *TemplateVersionSelect {
	tvs.fns = append(tvs.fns, fns...)
	return tvs
}

// Scan applies the selector query and scans the result into the given value.
func (tvs *TemplateVersionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tvs.ctx, "Select")
	if err := tvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TemplateVersionQuery, *TemplateVersionSelect](ctx, tvs.TemplateVersionQuery, tvs, tvs.inters, v)
}

func (tvs *TemplateVersionSelect) sqlScan(ctx context.Context, root *TemplateVersionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tvs.fns))
	for _, fn := range tvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tvs *TemplateVersionSelect) Modify(modifiers ...func(s *sql.Selector)) *TemplateVersionSelect {
	tvs.modifiers = append(tvs.modifiers, modifiers...)
	return tvs
}
