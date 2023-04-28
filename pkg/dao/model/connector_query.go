// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

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

	"github.com/seal-io/seal/pkg/dao/model/allocationcost"
	"github.com/seal-io/seal/pkg/dao/model/applicationresource"
	"github.com/seal-io/seal/pkg/dao/model/clustercost"
	"github.com/seal-io/seal/pkg/dao/model/connector"
	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// ConnectorQuery is the builder for querying Connector entities.
type ConnectorQuery struct {
	config
	ctx                 *QueryContext
	order               []connector.OrderOption
	inters              []Interceptor
	predicates          []predicate.Connector
	withEnvironments    *EnvironmentConnectorRelationshipQuery
	withResources       *ApplicationResourceQuery
	withClusterCosts    *ClusterCostQuery
	withAllocationCosts *AllocationCostQuery
	modifiers           []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConnectorQuery builder.
func (cq *ConnectorQuery) Where(ps ...predicate.Connector) *ConnectorQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConnectorQuery) Limit(limit int) *ConnectorQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConnectorQuery) Offset(offset int) *ConnectorQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConnectorQuery) Unique(unique bool) *ConnectorQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConnectorQuery) Order(o ...connector.OrderOption) *ConnectorQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryEnvironments chains the current query on the "environments" edge.
func (cq *ConnectorQuery) QueryEnvironments() *EnvironmentConnectorRelationshipQuery {
	query := (&EnvironmentConnectorRelationshipClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(environmentconnectorrelationship.Table, environmentconnectorrelationship.ConnectorColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, connector.EnvironmentsTable, connector.EnvironmentsColumn),
		)
		schemaConfig := cq.schemaConfig
		step.To.Schema = schemaConfig.EnvironmentConnectorRelationship
		step.Edge.Schema = schemaConfig.EnvironmentConnectorRelationship
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResources chains the current query on the "resources" edge.
func (cq *ConnectorQuery) QueryResources() *ApplicationResourceQuery {
	query := (&ApplicationResourceClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(applicationresource.Table, applicationresource.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, connector.ResourcesTable, connector.ResourcesColumn),
		)
		schemaConfig := cq.schemaConfig
		step.To.Schema = schemaConfig.ApplicationResource
		step.Edge.Schema = schemaConfig.ApplicationResource
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClusterCosts chains the current query on the "clusterCosts" edge.
func (cq *ConnectorQuery) QueryClusterCosts() *ClusterCostQuery {
	query := (&ClusterCostClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(clustercost.Table, clustercost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, connector.ClusterCostsTable, connector.ClusterCostsColumn),
		)
		schemaConfig := cq.schemaConfig
		step.To.Schema = schemaConfig.ClusterCost
		step.Edge.Schema = schemaConfig.ClusterCost
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAllocationCosts chains the current query on the "allocationCosts" edge.
func (cq *ConnectorQuery) QueryAllocationCosts() *AllocationCostQuery {
	query := (&AllocationCostClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(allocationcost.Table, allocationcost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, connector.AllocationCostsTable, connector.AllocationCostsColumn),
		)
		schemaConfig := cq.schemaConfig
		step.To.Schema = schemaConfig.AllocationCost
		step.Edge.Schema = schemaConfig.AllocationCost
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Connector entity from the query.
// Returns a *NotFoundError when no Connector was found.
func (cq *ConnectorQuery) First(ctx context.Context) (*Connector, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{connector.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConnectorQuery) FirstX(ctx context.Context) *Connector {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Connector ID from the query.
// Returns a *NotFoundError when no Connector ID was found.
func (cq *ConnectorQuery) FirstID(ctx context.Context) (id oid.ID, err error) {
	var ids []oid.ID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{connector.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConnectorQuery) FirstIDX(ctx context.Context) oid.ID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Connector entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Connector entity is found.
// Returns a *NotFoundError when no Connector entities are found.
func (cq *ConnectorQuery) Only(ctx context.Context) (*Connector, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{connector.Label}
	default:
		return nil, &NotSingularError{connector.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConnectorQuery) OnlyX(ctx context.Context) *Connector {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Connector ID in the query.
// Returns a *NotSingularError when more than one Connector ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConnectorQuery) OnlyID(ctx context.Context) (id oid.ID, err error) {
	var ids []oid.ID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{connector.Label}
	default:
		err = &NotSingularError{connector.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConnectorQuery) OnlyIDX(ctx context.Context) oid.ID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Connectors.
func (cq *ConnectorQuery) All(ctx context.Context) ([]*Connector, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Connector, *ConnectorQuery]()
	return withInterceptors[[]*Connector](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConnectorQuery) AllX(ctx context.Context) []*Connector {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Connector IDs.
func (cq *ConnectorQuery) IDs(ctx context.Context) (ids []oid.ID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(connector.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConnectorQuery) IDsX(ctx context.Context) []oid.ID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConnectorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConnectorQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConnectorQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConnectorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConnectorQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConnectorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConnectorQuery) Clone() *ConnectorQuery {
	if cq == nil {
		return nil
	}
	return &ConnectorQuery{
		config:              cq.config,
		ctx:                 cq.ctx.Clone(),
		order:               append([]connector.OrderOption{}, cq.order...),
		inters:              append([]Interceptor{}, cq.inters...),
		predicates:          append([]predicate.Connector{}, cq.predicates...),
		withEnvironments:    cq.withEnvironments.Clone(),
		withResources:       cq.withResources.Clone(),
		withClusterCosts:    cq.withClusterCosts.Clone(),
		withAllocationCosts: cq.withAllocationCosts.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithEnvironments tells the query-builder to eager-load the nodes that are connected to
// the "environments" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithEnvironments(opts ...func(*EnvironmentConnectorRelationshipQuery)) *ConnectorQuery {
	query := (&EnvironmentConnectorRelationshipClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withEnvironments = query
	return cq
}

// WithResources tells the query-builder to eager-load the nodes that are connected to
// the "resources" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithResources(opts ...func(*ApplicationResourceQuery)) *ConnectorQuery {
	query := (&ApplicationResourceClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withResources = query
	return cq
}

// WithClusterCosts tells the query-builder to eager-load the nodes that are connected to
// the "clusterCosts" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithClusterCosts(opts ...func(*ClusterCostQuery)) *ConnectorQuery {
	query := (&ClusterCostClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withClusterCosts = query
	return cq
}

// WithAllocationCosts tells the query-builder to eager-load the nodes that are connected to
// the "allocationCosts" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithAllocationCosts(opts ...func(*AllocationCostQuery)) *ConnectorQuery {
	query := (&AllocationCostClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAllocationCosts = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty" sql:"name"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Connector.Query().
//		GroupBy(connector.FieldName).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (cq *ConnectorQuery) GroupBy(field string, fields ...string) *ConnectorGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ConnectorGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = connector.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty" sql:"name"`
//	}
//
//	client.Connector.Query().
//		Select(connector.FieldName).
//		Scan(ctx, &v)
func (cq *ConnectorQuery) Select(fields ...string) *ConnectorSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ConnectorSelect{ConnectorQuery: cq}
	sbuild.label = connector.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConnectorSelect configured with the given aggregations.
func (cq *ConnectorQuery) Aggregate(fns ...AggregateFunc) *ConnectorSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConnectorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !connector.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConnectorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Connector, error) {
	var (
		nodes       = []*Connector{}
		_spec       = cq.querySpec()
		loadedTypes = [4]bool{
			cq.withEnvironments != nil,
			cq.withResources != nil,
			cq.withClusterCosts != nil,
			cq.withAllocationCosts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Connector).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Connector{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = cq.schemaConfig.Connector
	ctx = internal.NewSchemaConfigContext(ctx, cq.schemaConfig)
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withEnvironments; query != nil {
		if err := cq.loadEnvironments(ctx, query, nodes,
			func(n *Connector) { n.Edges.Environments = []*EnvironmentConnectorRelationship{} },
			func(n *Connector, e *EnvironmentConnectorRelationship) {
				n.Edges.Environments = append(n.Edges.Environments, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := cq.withResources; query != nil {
		if err := cq.loadResources(ctx, query, nodes,
			func(n *Connector) { n.Edges.Resources = []*ApplicationResource{} },
			func(n *Connector, e *ApplicationResource) { n.Edges.Resources = append(n.Edges.Resources, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withClusterCosts; query != nil {
		if err := cq.loadClusterCosts(ctx, query, nodes,
			func(n *Connector) { n.Edges.ClusterCosts = []*ClusterCost{} },
			func(n *Connector, e *ClusterCost) { n.Edges.ClusterCosts = append(n.Edges.ClusterCosts, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withAllocationCosts; query != nil {
		if err := cq.loadAllocationCosts(ctx, query, nodes,
			func(n *Connector) { n.Edges.AllocationCosts = []*AllocationCost{} },
			func(n *Connector, e *AllocationCost) { n.Edges.AllocationCosts = append(n.Edges.AllocationCosts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ConnectorQuery) loadEnvironments(ctx context.Context, query *EnvironmentConnectorRelationshipQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *EnvironmentConnectorRelationship)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[oid.ID]*Connector)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(environmentconnectorrelationship.FieldConnectorID)
	}
	query.Where(predicate.EnvironmentConnectorRelationship(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(connector.EnvironmentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ConnectorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "connector_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ConnectorQuery) loadResources(ctx context.Context, query *ApplicationResourceQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *ApplicationResource)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[oid.ID]*Connector)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(applicationresource.FieldConnectorID)
	}
	query.Where(predicate.ApplicationResource(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(connector.ResourcesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ConnectorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "connectorID" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ConnectorQuery) loadClusterCosts(ctx context.Context, query *ClusterCostQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *ClusterCost)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[oid.ID]*Connector)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(clustercost.FieldConnectorID)
	}
	query.Where(predicate.ClusterCost(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(connector.ClusterCostsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ConnectorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "connectorID" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ConnectorQuery) loadAllocationCosts(ctx context.Context, query *AllocationCostQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *AllocationCost)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[oid.ID]*Connector)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(allocationcost.FieldConnectorID)
	}
	query.Where(predicate.AllocationCost(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(connector.AllocationCostsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ConnectorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "connectorID" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ConnectorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Schema = cq.schemaConfig.Connector
	ctx = internal.NewSchemaConfigContext(ctx, cq.schemaConfig)
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConnectorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(connector.Table, connector.Columns, sqlgraph.NewFieldSpec(connector.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connector.FieldID)
		for i := range fields {
			if fields[i] != connector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConnectorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(connector.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = connector.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(cq.schemaConfig.Connector)
	ctx = internal.NewSchemaConfigContext(ctx, cq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cq *ConnectorQuery) ForUpdate(opts ...sql.LockOption) *ConnectorQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cq *ConnectorQuery) ForShare(opts ...sql.LockOption) *ConnectorQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *ConnectorQuery) Modify(modifiers ...func(s *sql.Selector)) *ConnectorSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// WhereP appends storage-level predicates to the ConnectorQuery builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (cq *ConnectorQuery) WhereP(ps ...func(*sql.Selector)) {
	var wps = make([]predicate.Connector, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		wps = append(wps, predicate.Connector(ps[i]))
	}
	cq.predicates = append(cq.predicates, wps...)
}

// ConnectorGroupBy is the group-by builder for Connector entities.
type ConnectorGroupBy struct {
	selector
	build *ConnectorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConnectorGroupBy) Aggregate(fns ...AggregateFunc) *ConnectorGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConnectorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConnectorQuery, *ConnectorGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConnectorGroupBy) sqlScan(ctx context.Context, root *ConnectorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConnectorSelect is the builder for selecting fields of Connector entities.
type ConnectorSelect struct {
	*ConnectorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConnectorSelect) Aggregate(fns ...AggregateFunc) *ConnectorSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConnectorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConnectorQuery, *ConnectorSelect](ctx, cs.ConnectorQuery, cs, cs.inters, v)
}

func (cs *ConnectorSelect) sqlScan(ctx context.Context, root *ConnectorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *ConnectorSelect) Modify(modifiers ...func(s *sql.Selector)) *ConnectorSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
