// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/service"
	"github.com/seal-io/seal/pkg/dao/model/servicerelationship"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// ServiceRelationshipCreate is the builder for creating a ServiceRelationship entity.
type ServiceRelationshipCreate struct {
	config
	mutation   *ServiceRelationshipMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *ServiceRelationship
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (src *ServiceRelationshipCreate) SetCreateTime(t time.Time) *ServiceRelationshipCreate {
	src.mutation.SetCreateTime(t)
	return src
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (src *ServiceRelationshipCreate) SetNillableCreateTime(t *time.Time) *ServiceRelationshipCreate {
	if t != nil {
		src.SetCreateTime(*t)
	}
	return src
}

// SetServiceID sets the "service_id" field.
func (src *ServiceRelationshipCreate) SetServiceID(o object.ID) *ServiceRelationshipCreate {
	src.mutation.SetServiceID(o)
	return src
}

// SetDependencyID sets the "dependency_id" field.
func (src *ServiceRelationshipCreate) SetDependencyID(o object.ID) *ServiceRelationshipCreate {
	src.mutation.SetDependencyID(o)
	return src
}

// SetPath sets the "path" field.
func (src *ServiceRelationshipCreate) SetPath(o []object.ID) *ServiceRelationshipCreate {
	src.mutation.SetPath(o)
	return src
}

// SetType sets the "type" field.
func (src *ServiceRelationshipCreate) SetType(s string) *ServiceRelationshipCreate {
	src.mutation.SetType(s)
	return src
}

// SetID sets the "id" field.
func (src *ServiceRelationshipCreate) SetID(o object.ID) *ServiceRelationshipCreate {
	src.mutation.SetID(o)
	return src
}

// SetService sets the "service" edge to the Service entity.
func (src *ServiceRelationshipCreate) SetService(s *Service) *ServiceRelationshipCreate {
	return src.SetServiceID(s.ID)
}

// SetDependency sets the "dependency" edge to the Service entity.
func (src *ServiceRelationshipCreate) SetDependency(s *Service) *ServiceRelationshipCreate {
	return src.SetDependencyID(s.ID)
}

// Mutation returns the ServiceRelationshipMutation object of the builder.
func (src *ServiceRelationshipCreate) Mutation() *ServiceRelationshipMutation {
	return src.mutation
}

// Save creates the ServiceRelationship in the database.
func (src *ServiceRelationshipCreate) Save(ctx context.Context) (*ServiceRelationship, error) {
	if err := src.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, src.sqlSave, src.mutation, src.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (src *ServiceRelationshipCreate) SaveX(ctx context.Context) *ServiceRelationship {
	v, err := src.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (src *ServiceRelationshipCreate) Exec(ctx context.Context) error {
	_, err := src.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (src *ServiceRelationshipCreate) ExecX(ctx context.Context) {
	if err := src.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (src *ServiceRelationshipCreate) defaults() error {
	if _, ok := src.mutation.CreateTime(); !ok {
		if servicerelationship.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized servicerelationship.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := servicerelationship.DefaultCreateTime()
		src.mutation.SetCreateTime(v)
	}
	if _, ok := src.mutation.Path(); !ok {
		v := servicerelationship.DefaultPath
		src.mutation.SetPath(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (src *ServiceRelationshipCreate) check() error {
	if _, ok := src.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "ServiceRelationship.create_time"`)}
	}
	if _, ok := src.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "service_id", err: errors.New(`model: missing required field "ServiceRelationship.service_id"`)}
	}
	if v, ok := src.mutation.ServiceID(); ok {
		if err := servicerelationship.ServiceIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "service_id", err: fmt.Errorf(`model: validator failed for field "ServiceRelationship.service_id": %w`, err)}
		}
	}
	if _, ok := src.mutation.DependencyID(); !ok {
		return &ValidationError{Name: "dependency_id", err: errors.New(`model: missing required field "ServiceRelationship.dependency_id"`)}
	}
	if v, ok := src.mutation.DependencyID(); ok {
		if err := servicerelationship.DependencyIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "dependency_id", err: fmt.Errorf(`model: validator failed for field "ServiceRelationship.dependency_id": %w`, err)}
		}
	}
	if _, ok := src.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`model: missing required field "ServiceRelationship.path"`)}
	}
	if _, ok := src.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "ServiceRelationship.type"`)}
	}
	if v, ok := src.mutation.GetType(); ok {
		if err := servicerelationship.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`model: validator failed for field "ServiceRelationship.type": %w`, err)}
		}
	}
	if _, ok := src.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "service", err: errors.New(`model: missing required edge "ServiceRelationship.service"`)}
	}
	if _, ok := src.mutation.DependencyID(); !ok {
		return &ValidationError{Name: "dependency", err: errors.New(`model: missing required edge "ServiceRelationship.dependency"`)}
	}
	return nil
}

func (src *ServiceRelationshipCreate) sqlSave(ctx context.Context) (*ServiceRelationship, error) {
	if err := src.check(); err != nil {
		return nil, err
	}
	_node, _spec := src.createSpec()
	if err := sqlgraph.CreateNode(ctx, src.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*object.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	src.mutation.id = &_node.ID
	src.mutation.done = true
	return _node, nil
}

func (src *ServiceRelationshipCreate) createSpec() (*ServiceRelationship, *sqlgraph.CreateSpec) {
	var (
		_node = &ServiceRelationship{config: src.config}
		_spec = sqlgraph.NewCreateSpec(servicerelationship.Table, sqlgraph.NewFieldSpec(servicerelationship.FieldID, field.TypeString))
	)
	_spec.Schema = src.schemaConfig.ServiceRelationship
	_spec.OnConflict = src.conflict
	if id, ok := src.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := src.mutation.CreateTime(); ok {
		_spec.SetField(servicerelationship.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := src.mutation.Path(); ok {
		_spec.SetField(servicerelationship.FieldPath, field.TypeJSON, value)
		_node.Path = value
	}
	if value, ok := src.mutation.GetType(); ok {
		_spec.SetField(servicerelationship.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := src.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servicerelationship.ServiceTable,
			Columns: []string{servicerelationship.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = src.schemaConfig.ServiceRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ServiceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := src.mutation.DependencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servicerelationship.DependencyTable,
			Columns: []string{servicerelationship.DependencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = src.schemaConfig.ServiceRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DependencyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (src *ServiceRelationshipCreate) Set(obj *ServiceRelationship) *ServiceRelationshipCreate {
	// Required.
	src.SetServiceID(obj.ServiceID)
	src.SetDependencyID(obj.DependencyID)
	src.SetPath(obj.Path)
	src.SetType(obj.Type)

	// Optional.
	if obj.CreateTime != nil {
		src.SetCreateTime(*obj.CreateTime)
	}

	// Record the given object.
	src.object = obj

	return src
}

// getClientSet returns the ClientSet for the given builder.
func (src *ServiceRelationshipCreate) getClientSet() (mc ClientSet) {
	if _, ok := src.config.driver.(*txDriver); ok {
		tx := &Tx{config: src.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: src.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the ServiceRelationship entity,
// which is always good for cascading create operations.
func (src *ServiceRelationshipCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) (*ServiceRelationship, error) {
	obj, err := src.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := src.getClientSet()

	if x := src.object; x != nil {
		if _, set := src.mutation.Field(servicerelationship.FieldServiceID); set {
			obj.ServiceID = x.ServiceID
		}
		if _, set := src.mutation.Field(servicerelationship.FieldDependencyID); set {
			obj.DependencyID = x.DependencyID
		}
		if _, set := src.mutation.Field(servicerelationship.FieldType); set {
			obj.Type = x.Type
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (src *ServiceRelationshipCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) *ServiceRelationship {
	obj, err := src.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (src *ServiceRelationshipCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) error {
	_, err := src.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (src *ServiceRelationshipCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) {
	if err := src.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the ServiceRelationshipCreate Set method,
// it sets the value by judging the definition of each field within the entire item of the given list.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (srcb *ServiceRelationshipCreateBulk) Set(objs ...*ServiceRelationship) *ServiceRelationshipCreateBulk {
	if len(objs) != 0 {
		client := NewServiceRelationshipClient(srcb.config)

		srcb.builders = make([]*ServiceRelationshipCreate, len(objs))
		for i := range objs {
			srcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		srcb.objects = objs
	}

	return srcb
}

// getClientSet returns the ClientSet for the given builder.
func (srcb *ServiceRelationshipCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := srcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: srcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: srcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the ServiceRelationship entities,
// which is always good for cascading create operations.
func (srcb *ServiceRelationshipCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) ([]*ServiceRelationship, error) {
	objs, err := srcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := srcb.getClientSet()

	if x := srcb.objects; x != nil {
		for i := range x {
			if _, set := srcb.builders[i].mutation.Field(servicerelationship.FieldServiceID); set {
				objs[i].ServiceID = x[i].ServiceID
			}
			if _, set := srcb.builders[i].mutation.Field(servicerelationship.FieldDependencyID); set {
				objs[i].DependencyID = x[i].DependencyID
			}
			if _, set := srcb.builders[i].mutation.Field(servicerelationship.FieldType); set {
				objs[i].Type = x[i].Type
			}
			objs[i].Edges = x[i].Edges
		}
	}

	for i := range objs {
		for j := range cbs {
			if err = cbs[j](ctx, mc, objs[i]); err != nil {
				return nil, err
			}
		}
	}

	return objs, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (srcb *ServiceRelationshipCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) []*ServiceRelationship {
	objs, err := srcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (srcb *ServiceRelationshipCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) error {
	_, err := srcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (srcb *ServiceRelationshipCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ServiceRelationship) error) {
	if err := srcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *ServiceRelationshipUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRelationship) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ServiceRelationshipUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *ServiceRelationshipUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRelationship) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *ServiceRelationshipUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRelationship) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ServiceRelationshipUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ServiceRelationshipUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *ServiceRelationshipUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRelationship) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ServiceRelationship.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServiceRelationshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (src *ServiceRelationshipCreate) OnConflict(opts ...sql.ConflictOption) *ServiceRelationshipUpsertOne {
	src.conflict = opts
	return &ServiceRelationshipUpsertOne{
		create: src,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ServiceRelationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (src *ServiceRelationshipCreate) OnConflictColumns(columns ...string) *ServiceRelationshipUpsertOne {
	src.conflict = append(src.conflict, sql.ConflictColumns(columns...))
	return &ServiceRelationshipUpsertOne{
		create: src,
	}
}

type (
	// ServiceRelationshipUpsertOne is the builder for "upsert"-ing
	//  one ServiceRelationship node.
	ServiceRelationshipUpsertOne struct {
		create *ServiceRelationshipCreate
	}

	// ServiceRelationshipUpsert is the "OnConflict" setter.
	ServiceRelationshipUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ServiceRelationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(servicerelationship.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServiceRelationshipUpsertOne) UpdateNewValues() *ServiceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(servicerelationship.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(servicerelationship.FieldCreateTime)
		}
		if _, exists := u.create.mutation.ServiceID(); exists {
			s.SetIgnore(servicerelationship.FieldServiceID)
		}
		if _, exists := u.create.mutation.DependencyID(); exists {
			s.SetIgnore(servicerelationship.FieldDependencyID)
		}
		if _, exists := u.create.mutation.Path(); exists {
			s.SetIgnore(servicerelationship.FieldPath)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(servicerelationship.FieldType)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ServiceRelationship.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ServiceRelationshipUpsertOne) Ignore() *ServiceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServiceRelationshipUpsertOne) DoNothing() *ServiceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServiceRelationshipCreate.OnConflict
// documentation for more info.
func (u *ServiceRelationshipUpsertOne) Update(set func(*ServiceRelationshipUpsert)) *ServiceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServiceRelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ServiceRelationshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ServiceRelationshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServiceRelationshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ServiceRelationshipUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: ServiceRelationshipUpsertOne.ID is not supported by MySQL driver. Use ServiceRelationshipUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ServiceRelationshipUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ServiceRelationshipCreateBulk is the builder for creating many ServiceRelationship entities in bulk.
type ServiceRelationshipCreateBulk struct {
	config
	builders   []*ServiceRelationshipCreate
	conflict   []sql.ConflictOption
	objects    []*ServiceRelationship
	fromUpsert bool
}

// Save creates the ServiceRelationship entities in the database.
func (srcb *ServiceRelationshipCreateBulk) Save(ctx context.Context) ([]*ServiceRelationship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(srcb.builders))
	nodes := make([]*ServiceRelationship, len(srcb.builders))
	mutators := make([]Mutator, len(srcb.builders))
	for i := range srcb.builders {
		func(i int, root context.Context) {
			builder := srcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceRelationshipMutation)
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
					_, err = mutators[i+1].Mutate(root, srcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = srcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, srcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (srcb *ServiceRelationshipCreateBulk) SaveX(ctx context.Context) []*ServiceRelationship {
	v, err := srcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (srcb *ServiceRelationshipCreateBulk) Exec(ctx context.Context) error {
	_, err := srcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srcb *ServiceRelationshipCreateBulk) ExecX(ctx context.Context) {
	if err := srcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ServiceRelationship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServiceRelationshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (srcb *ServiceRelationshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *ServiceRelationshipUpsertBulk {
	srcb.conflict = opts
	return &ServiceRelationshipUpsertBulk{
		create: srcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ServiceRelationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (srcb *ServiceRelationshipCreateBulk) OnConflictColumns(columns ...string) *ServiceRelationshipUpsertBulk {
	srcb.conflict = append(srcb.conflict, sql.ConflictColumns(columns...))
	return &ServiceRelationshipUpsertBulk{
		create: srcb,
	}
}

// ServiceRelationshipUpsertBulk is the builder for "upsert"-ing
// a bulk of ServiceRelationship nodes.
type ServiceRelationshipUpsertBulk struct {
	create *ServiceRelationshipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ServiceRelationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(servicerelationship.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServiceRelationshipUpsertBulk) UpdateNewValues() *ServiceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(servicerelationship.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(servicerelationship.FieldCreateTime)
			}
			if _, exists := b.mutation.ServiceID(); exists {
				s.SetIgnore(servicerelationship.FieldServiceID)
			}
			if _, exists := b.mutation.DependencyID(); exists {
				s.SetIgnore(servicerelationship.FieldDependencyID)
			}
			if _, exists := b.mutation.Path(); exists {
				s.SetIgnore(servicerelationship.FieldPath)
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(servicerelationship.FieldType)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ServiceRelationship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ServiceRelationshipUpsertBulk) Ignore() *ServiceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServiceRelationshipUpsertBulk) DoNothing() *ServiceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServiceRelationshipCreateBulk.OnConflict
// documentation for more info.
func (u *ServiceRelationshipUpsertBulk) Update(set func(*ServiceRelationshipUpsert)) *ServiceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServiceRelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ServiceRelationshipUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ServiceRelationshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ServiceRelationshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServiceRelationshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
