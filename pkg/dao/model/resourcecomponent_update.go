// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/resourcecomponent"
	"github.com/seal-io/walrus/pkg/dao/model/resourcecomponentrelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// ResourceComponentUpdate is the builder for updating ResourceComponent entities.
type ResourceComponentUpdate struct {
	config
	hooks     []Hook
	mutation  *ResourceComponentMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceComponent
}

// Where appends a list predicates to the ResourceComponentUpdate builder.
func (rcu *ResourceComponentUpdate) Where(ps ...predicate.ResourceComponent) *ResourceComponentUpdate {
	rcu.mutation.Where(ps...)
	return rcu
}

// SetUpdateTime sets the "update_time" field.
func (rcu *ResourceComponentUpdate) SetUpdateTime(t time.Time) *ResourceComponentUpdate {
	rcu.mutation.SetUpdateTime(t)
	return rcu
}

// SetStatus sets the "status" field.
func (rcu *ResourceComponentUpdate) SetStatus(s status.Status) *ResourceComponentUpdate {
	rcu.mutation.SetStatus(s)
	return rcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rcu *ResourceComponentUpdate) SetNillableStatus(s *status.Status) *ResourceComponentUpdate {
	if s != nil {
		rcu.SetStatus(*s)
	}
	return rcu
}

// ClearStatus clears the value of the "status" field.
func (rcu *ResourceComponentUpdate) ClearStatus() *ResourceComponentUpdate {
	rcu.mutation.ClearStatus()
	return rcu
}

// AddComponentIDs adds the "components" edge to the ResourceComponent entity by IDs.
func (rcu *ResourceComponentUpdate) AddComponentIDs(ids ...object.ID) *ResourceComponentUpdate {
	rcu.mutation.AddComponentIDs(ids...)
	return rcu
}

// AddComponents adds the "components" edges to the ResourceComponent entity.
func (rcu *ResourceComponentUpdate) AddComponents(r ...*ResourceComponent) *ResourceComponentUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.AddComponentIDs(ids...)
}

// AddInstanceIDs adds the "instances" edge to the ResourceComponent entity by IDs.
func (rcu *ResourceComponentUpdate) AddInstanceIDs(ids ...object.ID) *ResourceComponentUpdate {
	rcu.mutation.AddInstanceIDs(ids...)
	return rcu
}

// AddInstances adds the "instances" edges to the ResourceComponent entity.
func (rcu *ResourceComponentUpdate) AddInstances(r ...*ResourceComponent) *ResourceComponentUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.AddInstanceIDs(ids...)
}

// AddDependencyIDs adds the "dependencies" edge to the ResourceComponentRelationship entity by IDs.
func (rcu *ResourceComponentUpdate) AddDependencyIDs(ids ...object.ID) *ResourceComponentUpdate {
	rcu.mutation.AddDependencyIDs(ids...)
	return rcu
}

// AddDependencies adds the "dependencies" edges to the ResourceComponentRelationship entity.
func (rcu *ResourceComponentUpdate) AddDependencies(r ...*ResourceComponentRelationship) *ResourceComponentUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.AddDependencyIDs(ids...)
}

// Mutation returns the ResourceComponentMutation object of the builder.
func (rcu *ResourceComponentUpdate) Mutation() *ResourceComponentMutation {
	return rcu.mutation
}

// ClearComponents clears all "components" edges to the ResourceComponent entity.
func (rcu *ResourceComponentUpdate) ClearComponents() *ResourceComponentUpdate {
	rcu.mutation.ClearComponents()
	return rcu
}

// RemoveComponentIDs removes the "components" edge to ResourceComponent entities by IDs.
func (rcu *ResourceComponentUpdate) RemoveComponentIDs(ids ...object.ID) *ResourceComponentUpdate {
	rcu.mutation.RemoveComponentIDs(ids...)
	return rcu
}

// RemoveComponents removes "components" edges to ResourceComponent entities.
func (rcu *ResourceComponentUpdate) RemoveComponents(r ...*ResourceComponent) *ResourceComponentUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.RemoveComponentIDs(ids...)
}

// ClearInstances clears all "instances" edges to the ResourceComponent entity.
func (rcu *ResourceComponentUpdate) ClearInstances() *ResourceComponentUpdate {
	rcu.mutation.ClearInstances()
	return rcu
}

// RemoveInstanceIDs removes the "instances" edge to ResourceComponent entities by IDs.
func (rcu *ResourceComponentUpdate) RemoveInstanceIDs(ids ...object.ID) *ResourceComponentUpdate {
	rcu.mutation.RemoveInstanceIDs(ids...)
	return rcu
}

// RemoveInstances removes "instances" edges to ResourceComponent entities.
func (rcu *ResourceComponentUpdate) RemoveInstances(r ...*ResourceComponent) *ResourceComponentUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.RemoveInstanceIDs(ids...)
}

// ClearDependencies clears all "dependencies" edges to the ResourceComponentRelationship entity.
func (rcu *ResourceComponentUpdate) ClearDependencies() *ResourceComponentUpdate {
	rcu.mutation.ClearDependencies()
	return rcu
}

// RemoveDependencyIDs removes the "dependencies" edge to ResourceComponentRelationship entities by IDs.
func (rcu *ResourceComponentUpdate) RemoveDependencyIDs(ids ...object.ID) *ResourceComponentUpdate {
	rcu.mutation.RemoveDependencyIDs(ids...)
	return rcu
}

// RemoveDependencies removes "dependencies" edges to ResourceComponentRelationship entities.
func (rcu *ResourceComponentUpdate) RemoveDependencies(r ...*ResourceComponentRelationship) *ResourceComponentUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.RemoveDependencyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rcu *ResourceComponentUpdate) Save(ctx context.Context) (int, error) {
	if err := rcu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, rcu.sqlSave, rcu.mutation, rcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcu *ResourceComponentUpdate) SaveX(ctx context.Context) int {
	affected, err := rcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rcu *ResourceComponentUpdate) Exec(ctx context.Context) error {
	_, err := rcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcu *ResourceComponentUpdate) ExecX(ctx context.Context) {
	if err := rcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcu *ResourceComponentUpdate) defaults() error {
	if _, ok := rcu.mutation.UpdateTime(); !ok {
		if resourcecomponent.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized resourcecomponent.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := resourcecomponent.UpdateDefaultUpdateTime()
		rcu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rcu *ResourceComponentUpdate) check() error {
	if _, ok := rcu.mutation.ProjectID(); rcu.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.project"`)
	}
	if _, ok := rcu.mutation.EnvironmentID(); rcu.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.environment"`)
	}
	if _, ok := rcu.mutation.ResourceID(); rcu.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.resource"`)
	}
	if _, ok := rcu.mutation.ConnectorID(); rcu.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.connector"`)
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value is not zero.
//
// For no default but required fields, Set calls directly.
//
// For no default but optional fields, Set calls if the value is not zero,
// or clears if the value is zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (rcu *ResourceComponentUpdate) Set(obj *ResourceComponent) *ResourceComponentUpdate {
	// Without Default.
	if !reflect.ValueOf(obj.Status).IsZero() {
		rcu.SetStatus(obj.Status)
	}

	// With Default.
	if obj.UpdateTime != nil {
		rcu.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	rcu.object = obj

	return rcu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rcu *ResourceComponentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceComponentUpdate {
	rcu.modifiers = append(rcu.modifiers, modifiers...)
	return rcu
}

func (rcu *ResourceComponentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcecomponent.Table, resourcecomponent.Columns, sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString))
	if ps := rcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcu.mutation.UpdateTime(); ok {
		_spec.SetField(resourcecomponent.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rcu.mutation.Status(); ok {
		_spec.SetField(resourcecomponent.FieldStatus, field.TypeJSON, value)
	}
	if rcu.mutation.StatusCleared() {
		_spec.ClearField(resourcecomponent.FieldStatus, field.TypeJSON)
	}
	if rcu.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.ComponentsTable,
			Columns: []string{resourcecomponent.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponent
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.RemovedComponentsIDs(); len(nodes) > 0 && !rcu.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.ComponentsTable,
			Columns: []string{resourcecomponent.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.ComponentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.ComponentsTable,
			Columns: []string{resourcecomponent.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rcu.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.InstancesTable,
			Columns: []string{resourcecomponent.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponent
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !rcu.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.InstancesTable,
			Columns: []string{resourcecomponent.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.InstancesTable,
			Columns: []string{resourcecomponent.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rcu.mutation.DependenciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   resourcecomponent.DependenciesTable,
			Columns: []string{resourcecomponent.DependenciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponentrelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponentRelationship
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.RemovedDependenciesIDs(); len(nodes) > 0 && !rcu.mutation.DependenciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   resourcecomponent.DependenciesTable,
			Columns: []string{resourcecomponent.DependenciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponentrelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponentRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.DependenciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   resourcecomponent.DependenciesTable,
			Columns: []string{resourcecomponent.DependenciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponentrelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcu.schemaConfig.ResourceComponentRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = rcu.schemaConfig.ResourceComponent
	ctx = internal.NewSchemaConfigContext(ctx, rcu.schemaConfig)
	_spec.AddModifiers(rcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcecomponent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rcu.mutation.done = true
	return n, nil
}

// ResourceComponentUpdateOne is the builder for updating a single ResourceComponent entity.
type ResourceComponentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ResourceComponentMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceComponent
}

// SetUpdateTime sets the "update_time" field.
func (rcuo *ResourceComponentUpdateOne) SetUpdateTime(t time.Time) *ResourceComponentUpdateOne {
	rcuo.mutation.SetUpdateTime(t)
	return rcuo
}

// SetStatus sets the "status" field.
func (rcuo *ResourceComponentUpdateOne) SetStatus(s status.Status) *ResourceComponentUpdateOne {
	rcuo.mutation.SetStatus(s)
	return rcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rcuo *ResourceComponentUpdateOne) SetNillableStatus(s *status.Status) *ResourceComponentUpdateOne {
	if s != nil {
		rcuo.SetStatus(*s)
	}
	return rcuo
}

// ClearStatus clears the value of the "status" field.
func (rcuo *ResourceComponentUpdateOne) ClearStatus() *ResourceComponentUpdateOne {
	rcuo.mutation.ClearStatus()
	return rcuo
}

// AddComponentIDs adds the "components" edge to the ResourceComponent entity by IDs.
func (rcuo *ResourceComponentUpdateOne) AddComponentIDs(ids ...object.ID) *ResourceComponentUpdateOne {
	rcuo.mutation.AddComponentIDs(ids...)
	return rcuo
}

// AddComponents adds the "components" edges to the ResourceComponent entity.
func (rcuo *ResourceComponentUpdateOne) AddComponents(r ...*ResourceComponent) *ResourceComponentUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.AddComponentIDs(ids...)
}

// AddInstanceIDs adds the "instances" edge to the ResourceComponent entity by IDs.
func (rcuo *ResourceComponentUpdateOne) AddInstanceIDs(ids ...object.ID) *ResourceComponentUpdateOne {
	rcuo.mutation.AddInstanceIDs(ids...)
	return rcuo
}

// AddInstances adds the "instances" edges to the ResourceComponent entity.
func (rcuo *ResourceComponentUpdateOne) AddInstances(r ...*ResourceComponent) *ResourceComponentUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.AddInstanceIDs(ids...)
}

// AddDependencyIDs adds the "dependencies" edge to the ResourceComponentRelationship entity by IDs.
func (rcuo *ResourceComponentUpdateOne) AddDependencyIDs(ids ...object.ID) *ResourceComponentUpdateOne {
	rcuo.mutation.AddDependencyIDs(ids...)
	return rcuo
}

// AddDependencies adds the "dependencies" edges to the ResourceComponentRelationship entity.
func (rcuo *ResourceComponentUpdateOne) AddDependencies(r ...*ResourceComponentRelationship) *ResourceComponentUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.AddDependencyIDs(ids...)
}

// Mutation returns the ResourceComponentMutation object of the builder.
func (rcuo *ResourceComponentUpdateOne) Mutation() *ResourceComponentMutation {
	return rcuo.mutation
}

// ClearComponents clears all "components" edges to the ResourceComponent entity.
func (rcuo *ResourceComponentUpdateOne) ClearComponents() *ResourceComponentUpdateOne {
	rcuo.mutation.ClearComponents()
	return rcuo
}

// RemoveComponentIDs removes the "components" edge to ResourceComponent entities by IDs.
func (rcuo *ResourceComponentUpdateOne) RemoveComponentIDs(ids ...object.ID) *ResourceComponentUpdateOne {
	rcuo.mutation.RemoveComponentIDs(ids...)
	return rcuo
}

// RemoveComponents removes "components" edges to ResourceComponent entities.
func (rcuo *ResourceComponentUpdateOne) RemoveComponents(r ...*ResourceComponent) *ResourceComponentUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.RemoveComponentIDs(ids...)
}

// ClearInstances clears all "instances" edges to the ResourceComponent entity.
func (rcuo *ResourceComponentUpdateOne) ClearInstances() *ResourceComponentUpdateOne {
	rcuo.mutation.ClearInstances()
	return rcuo
}

// RemoveInstanceIDs removes the "instances" edge to ResourceComponent entities by IDs.
func (rcuo *ResourceComponentUpdateOne) RemoveInstanceIDs(ids ...object.ID) *ResourceComponentUpdateOne {
	rcuo.mutation.RemoveInstanceIDs(ids...)
	return rcuo
}

// RemoveInstances removes "instances" edges to ResourceComponent entities.
func (rcuo *ResourceComponentUpdateOne) RemoveInstances(r ...*ResourceComponent) *ResourceComponentUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.RemoveInstanceIDs(ids...)
}

// ClearDependencies clears all "dependencies" edges to the ResourceComponentRelationship entity.
func (rcuo *ResourceComponentUpdateOne) ClearDependencies() *ResourceComponentUpdateOne {
	rcuo.mutation.ClearDependencies()
	return rcuo
}

// RemoveDependencyIDs removes the "dependencies" edge to ResourceComponentRelationship entities by IDs.
func (rcuo *ResourceComponentUpdateOne) RemoveDependencyIDs(ids ...object.ID) *ResourceComponentUpdateOne {
	rcuo.mutation.RemoveDependencyIDs(ids...)
	return rcuo
}

// RemoveDependencies removes "dependencies" edges to ResourceComponentRelationship entities.
func (rcuo *ResourceComponentUpdateOne) RemoveDependencies(r ...*ResourceComponentRelationship) *ResourceComponentUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.RemoveDependencyIDs(ids...)
}

// Where appends a list predicates to the ResourceComponentUpdate builder.
func (rcuo *ResourceComponentUpdateOne) Where(ps ...predicate.ResourceComponent) *ResourceComponentUpdateOne {
	rcuo.mutation.Where(ps...)
	return rcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rcuo *ResourceComponentUpdateOne) Select(field string, fields ...string) *ResourceComponentUpdateOne {
	rcuo.fields = append([]string{field}, fields...)
	return rcuo
}

// Save executes the query and returns the updated ResourceComponent entity.
func (rcuo *ResourceComponentUpdateOne) Save(ctx context.Context) (*ResourceComponent, error) {
	if err := rcuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rcuo.sqlSave, rcuo.mutation, rcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcuo *ResourceComponentUpdateOne) SaveX(ctx context.Context) *ResourceComponent {
	node, err := rcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rcuo *ResourceComponentUpdateOne) Exec(ctx context.Context) error {
	_, err := rcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *ResourceComponentUpdateOne) ExecX(ctx context.Context) {
	if err := rcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcuo *ResourceComponentUpdateOne) defaults() error {
	if _, ok := rcuo.mutation.UpdateTime(); !ok {
		if resourcecomponent.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized resourcecomponent.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := resourcecomponent.UpdateDefaultUpdateTime()
		rcuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rcuo *ResourceComponentUpdateOne) check() error {
	if _, ok := rcuo.mutation.ProjectID(); rcuo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.project"`)
	}
	if _, ok := rcuo.mutation.EnvironmentID(); rcuo.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.environment"`)
	}
	if _, ok := rcuo.mutation.ResourceID(); rcuo.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.resource"`)
	}
	if _, ok := rcuo.mutation.ConnectorID(); rcuo.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceComponent.connector"`)
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value changes from the original.
//
// For no default but required fields, Set calls if the value changes from the original.
//
// For no default but optional fields, Set calls if the value changes from the original,
// or clears if changes to zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   if _is_not_equal_(db.X, obj.X) {
//	      db.SetX(obj.X)
//	   }
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) && _is_not_equal_(db.X, obj.X) {
//	   db.SetX(obj.X)
//	}
func (rcuo *ResourceComponentUpdateOne) Set(obj *ResourceComponent) *ResourceComponentUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*ResourceComponentMutation)
			db, err := mt.Client().ResourceComponent.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting ResourceComponent with id: %v", *mt.id)
			}

			// Without Default.
			if !reflect.ValueOf(obj.Status).IsZero() {
				if !db.Status.Equal(obj.Status) {
					rcuo.SetStatus(obj.Status)
				}
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				rcuo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			rcuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	rcuo.hooks = append(rcuo.hooks, h)

	return rcuo
}

// getClientSet returns the ClientSet for the given builder.
func (rcuo *ResourceComponentUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := rcuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: rcuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rcuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the ResourceComponent entity,
// which is always good for cascading update operations.
func (rcuo *ResourceComponentUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceComponent) error) (*ResourceComponent, error) {
	obj, err := rcuo.Save(ctx)
	if err != nil &&
		(rcuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := rcuo.getClientSet()

	if obj == nil {
		obj = rcuo.object
	} else if x := rcuo.object; x != nil {
		if _, set := rcuo.mutation.Field(resourcecomponent.FieldStatus); set {
			obj.Status = x.Status
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
func (rcuo *ResourceComponentUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceComponent) error) *ResourceComponent {
	obj, err := rcuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (rcuo *ResourceComponentUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceComponent) error) error {
	_, err := rcuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *ResourceComponentUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceComponent) error) {
	if err := rcuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rcuo *ResourceComponentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceComponentUpdateOne {
	rcuo.modifiers = append(rcuo.modifiers, modifiers...)
	return rcuo
}

func (rcuo *ResourceComponentUpdateOne) sqlSave(ctx context.Context) (_node *ResourceComponent, err error) {
	if err := rcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcecomponent.Table, resourcecomponent.Columns, sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString))
	id, ok := rcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ResourceComponent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourcecomponent.FieldID)
		for _, f := range fields {
			if !resourcecomponent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != resourcecomponent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcuo.mutation.UpdateTime(); ok {
		_spec.SetField(resourcecomponent.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rcuo.mutation.Status(); ok {
		_spec.SetField(resourcecomponent.FieldStatus, field.TypeJSON, value)
	}
	if rcuo.mutation.StatusCleared() {
		_spec.ClearField(resourcecomponent.FieldStatus, field.TypeJSON)
	}
	if rcuo.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.ComponentsTable,
			Columns: []string{resourcecomponent.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponent
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.RemovedComponentsIDs(); len(nodes) > 0 && !rcuo.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.ComponentsTable,
			Columns: []string{resourcecomponent.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.ComponentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.ComponentsTable,
			Columns: []string{resourcecomponent.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rcuo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.InstancesTable,
			Columns: []string{resourcecomponent.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponent
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !rcuo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.InstancesTable,
			Columns: []string{resourcecomponent.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcecomponent.InstancesTable,
			Columns: []string{resourcecomponent.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponent.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rcuo.mutation.DependenciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   resourcecomponent.DependenciesTable,
			Columns: []string{resourcecomponent.DependenciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponentrelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponentRelationship
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.RemovedDependenciesIDs(); len(nodes) > 0 && !rcuo.mutation.DependenciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   resourcecomponent.DependenciesTable,
			Columns: []string{resourcecomponent.DependenciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponentrelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponentRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.DependenciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   resourcecomponent.DependenciesTable,
			Columns: []string{resourcecomponent.DependenciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourcecomponentrelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = rcuo.schemaConfig.ResourceComponentRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = rcuo.schemaConfig.ResourceComponent
	ctx = internal.NewSchemaConfigContext(ctx, rcuo.schemaConfig)
	_spec.AddModifiers(rcuo.modifiers...)
	_node = &ResourceComponent{config: rcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcecomponent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rcuo.mutation.done = true
	return _node, nil
}
