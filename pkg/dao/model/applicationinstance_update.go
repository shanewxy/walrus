// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/applicationinstance"
	"github.com/seal-io/seal/pkg/dao/model/applicationresource"
	"github.com/seal-io/seal/pkg/dao/model/applicationrevision"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types"
)

// ApplicationInstanceUpdate is the builder for updating ApplicationInstance entities.
type ApplicationInstanceUpdate struct {
	config
	hooks     []Hook
	mutation  *ApplicationInstanceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ApplicationInstanceUpdate builder.
func (aiu *ApplicationInstanceUpdate) Where(ps ...predicate.ApplicationInstance) *ApplicationInstanceUpdate {
	aiu.mutation.Where(ps...)
	return aiu
}

// SetStatus sets the "status" field.
func (aiu *ApplicationInstanceUpdate) SetStatus(s string) *ApplicationInstanceUpdate {
	aiu.mutation.SetStatus(s)
	return aiu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aiu *ApplicationInstanceUpdate) SetNillableStatus(s *string) *ApplicationInstanceUpdate {
	if s != nil {
		aiu.SetStatus(*s)
	}
	return aiu
}

// ClearStatus clears the value of the "status" field.
func (aiu *ApplicationInstanceUpdate) ClearStatus() *ApplicationInstanceUpdate {
	aiu.mutation.ClearStatus()
	return aiu
}

// SetStatusMessage sets the "statusMessage" field.
func (aiu *ApplicationInstanceUpdate) SetStatusMessage(s string) *ApplicationInstanceUpdate {
	aiu.mutation.SetStatusMessage(s)
	return aiu
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (aiu *ApplicationInstanceUpdate) SetNillableStatusMessage(s *string) *ApplicationInstanceUpdate {
	if s != nil {
		aiu.SetStatusMessage(*s)
	}
	return aiu
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (aiu *ApplicationInstanceUpdate) ClearStatusMessage() *ApplicationInstanceUpdate {
	aiu.mutation.ClearStatusMessage()
	return aiu
}

// SetUpdateTime sets the "updateTime" field.
func (aiu *ApplicationInstanceUpdate) SetUpdateTime(t time.Time) *ApplicationInstanceUpdate {
	aiu.mutation.SetUpdateTime(t)
	return aiu
}

// SetVariables sets the "variables" field.
func (aiu *ApplicationInstanceUpdate) SetVariables(m map[string]interface{}) *ApplicationInstanceUpdate {
	aiu.mutation.SetVariables(m)
	return aiu
}

// ClearVariables clears the value of the "variables" field.
func (aiu *ApplicationInstanceUpdate) ClearVariables() *ApplicationInstanceUpdate {
	aiu.mutation.ClearVariables()
	return aiu
}

// AddRevisionIDs adds the "revisions" edge to the ApplicationRevision entity by IDs.
func (aiu *ApplicationInstanceUpdate) AddRevisionIDs(ids ...types.ID) *ApplicationInstanceUpdate {
	aiu.mutation.AddRevisionIDs(ids...)
	return aiu
}

// AddRevisions adds the "revisions" edges to the ApplicationRevision entity.
func (aiu *ApplicationInstanceUpdate) AddRevisions(a ...*ApplicationRevision) *ApplicationInstanceUpdate {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiu.AddRevisionIDs(ids...)
}

// AddResourceIDs adds the "resources" edge to the ApplicationResource entity by IDs.
func (aiu *ApplicationInstanceUpdate) AddResourceIDs(ids ...types.ID) *ApplicationInstanceUpdate {
	aiu.mutation.AddResourceIDs(ids...)
	return aiu
}

// AddResources adds the "resources" edges to the ApplicationResource entity.
func (aiu *ApplicationInstanceUpdate) AddResources(a ...*ApplicationResource) *ApplicationInstanceUpdate {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiu.AddResourceIDs(ids...)
}

// Mutation returns the ApplicationInstanceMutation object of the builder.
func (aiu *ApplicationInstanceUpdate) Mutation() *ApplicationInstanceMutation {
	return aiu.mutation
}

// ClearRevisions clears all "revisions" edges to the ApplicationRevision entity.
func (aiu *ApplicationInstanceUpdate) ClearRevisions() *ApplicationInstanceUpdate {
	aiu.mutation.ClearRevisions()
	return aiu
}

// RemoveRevisionIDs removes the "revisions" edge to ApplicationRevision entities by IDs.
func (aiu *ApplicationInstanceUpdate) RemoveRevisionIDs(ids ...types.ID) *ApplicationInstanceUpdate {
	aiu.mutation.RemoveRevisionIDs(ids...)
	return aiu
}

// RemoveRevisions removes "revisions" edges to ApplicationRevision entities.
func (aiu *ApplicationInstanceUpdate) RemoveRevisions(a ...*ApplicationRevision) *ApplicationInstanceUpdate {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiu.RemoveRevisionIDs(ids...)
}

// ClearResources clears all "resources" edges to the ApplicationResource entity.
func (aiu *ApplicationInstanceUpdate) ClearResources() *ApplicationInstanceUpdate {
	aiu.mutation.ClearResources()
	return aiu
}

// RemoveResourceIDs removes the "resources" edge to ApplicationResource entities by IDs.
func (aiu *ApplicationInstanceUpdate) RemoveResourceIDs(ids ...types.ID) *ApplicationInstanceUpdate {
	aiu.mutation.RemoveResourceIDs(ids...)
	return aiu
}

// RemoveResources removes "resources" edges to ApplicationResource entities.
func (aiu *ApplicationInstanceUpdate) RemoveResources(a ...*ApplicationResource) *ApplicationInstanceUpdate {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiu.RemoveResourceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *ApplicationInstanceUpdate) Save(ctx context.Context) (int, error) {
	if err := aiu.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, ApplicationInstanceMutation](ctx, aiu.sqlSave, aiu.mutation, aiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *ApplicationInstanceUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *ApplicationInstanceUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *ApplicationInstanceUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiu *ApplicationInstanceUpdate) defaults() error {
	if _, ok := aiu.mutation.UpdateTime(); !ok {
		if applicationinstance.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized applicationinstance.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := applicationinstance.UpdateDefaultUpdateTime()
		aiu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aiu *ApplicationInstanceUpdate) check() error {
	if _, ok := aiu.mutation.ApplicationID(); aiu.mutation.ApplicationCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationInstance.application"`)
	}
	if _, ok := aiu.mutation.EnvironmentID(); aiu.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationInstance.environment"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aiu *ApplicationInstanceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApplicationInstanceUpdate {
	aiu.modifiers = append(aiu.modifiers, modifiers...)
	return aiu
}

func (aiu *ApplicationInstanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(applicationinstance.Table, applicationinstance.Columns, sqlgraph.NewFieldSpec(applicationinstance.FieldID, field.TypeString))
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiu.mutation.Status(); ok {
		_spec.SetField(applicationinstance.FieldStatus, field.TypeString, value)
	}
	if aiu.mutation.StatusCleared() {
		_spec.ClearField(applicationinstance.FieldStatus, field.TypeString)
	}
	if value, ok := aiu.mutation.StatusMessage(); ok {
		_spec.SetField(applicationinstance.FieldStatusMessage, field.TypeString, value)
	}
	if aiu.mutation.StatusMessageCleared() {
		_spec.ClearField(applicationinstance.FieldStatusMessage, field.TypeString)
	}
	if value, ok := aiu.mutation.UpdateTime(); ok {
		_spec.SetField(applicationinstance.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := aiu.mutation.Variables(); ok {
		_spec.SetField(applicationinstance.FieldVariables, field.TypeJSON, value)
	}
	if aiu.mutation.VariablesCleared() {
		_spec.ClearField(applicationinstance.FieldVariables, field.TypeJSON)
	}
	if aiu.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.RevisionsTable,
			Columns: []string{applicationinstance.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationrevision.FieldID,
				},
			},
		}
		edge.Schema = aiu.schemaConfig.ApplicationRevision
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RemovedRevisionsIDs(); len(nodes) > 0 && !aiu.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.RevisionsTable,
			Columns: []string{applicationinstance.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationrevision.FieldID,
				},
			},
		}
		edge.Schema = aiu.schemaConfig.ApplicationRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.RevisionsTable,
			Columns: []string{applicationinstance.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationrevision.FieldID,
				},
			},
		}
		edge.Schema = aiu.schemaConfig.ApplicationRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiu.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.ResourcesTable,
			Columns: []string{applicationinstance.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aiu.schemaConfig.ApplicationResource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !aiu.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.ResourcesTable,
			Columns: []string{applicationinstance.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aiu.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.ResourcesTable,
			Columns: []string{applicationinstance.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aiu.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = aiu.schemaConfig.ApplicationInstance
	ctx = internal.NewSchemaConfigContext(ctx, aiu.schemaConfig)
	_spec.AddModifiers(aiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aiu.mutation.done = true
	return n, nil
}

// ApplicationInstanceUpdateOne is the builder for updating a single ApplicationInstance entity.
type ApplicationInstanceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ApplicationInstanceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStatus sets the "status" field.
func (aiuo *ApplicationInstanceUpdateOne) SetStatus(s string) *ApplicationInstanceUpdateOne {
	aiuo.mutation.SetStatus(s)
	return aiuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aiuo *ApplicationInstanceUpdateOne) SetNillableStatus(s *string) *ApplicationInstanceUpdateOne {
	if s != nil {
		aiuo.SetStatus(*s)
	}
	return aiuo
}

// ClearStatus clears the value of the "status" field.
func (aiuo *ApplicationInstanceUpdateOne) ClearStatus() *ApplicationInstanceUpdateOne {
	aiuo.mutation.ClearStatus()
	return aiuo
}

// SetStatusMessage sets the "statusMessage" field.
func (aiuo *ApplicationInstanceUpdateOne) SetStatusMessage(s string) *ApplicationInstanceUpdateOne {
	aiuo.mutation.SetStatusMessage(s)
	return aiuo
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (aiuo *ApplicationInstanceUpdateOne) SetNillableStatusMessage(s *string) *ApplicationInstanceUpdateOne {
	if s != nil {
		aiuo.SetStatusMessage(*s)
	}
	return aiuo
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (aiuo *ApplicationInstanceUpdateOne) ClearStatusMessage() *ApplicationInstanceUpdateOne {
	aiuo.mutation.ClearStatusMessage()
	return aiuo
}

// SetUpdateTime sets the "updateTime" field.
func (aiuo *ApplicationInstanceUpdateOne) SetUpdateTime(t time.Time) *ApplicationInstanceUpdateOne {
	aiuo.mutation.SetUpdateTime(t)
	return aiuo
}

// SetVariables sets the "variables" field.
func (aiuo *ApplicationInstanceUpdateOne) SetVariables(m map[string]interface{}) *ApplicationInstanceUpdateOne {
	aiuo.mutation.SetVariables(m)
	return aiuo
}

// ClearVariables clears the value of the "variables" field.
func (aiuo *ApplicationInstanceUpdateOne) ClearVariables() *ApplicationInstanceUpdateOne {
	aiuo.mutation.ClearVariables()
	return aiuo
}

// AddRevisionIDs adds the "revisions" edge to the ApplicationRevision entity by IDs.
func (aiuo *ApplicationInstanceUpdateOne) AddRevisionIDs(ids ...types.ID) *ApplicationInstanceUpdateOne {
	aiuo.mutation.AddRevisionIDs(ids...)
	return aiuo
}

// AddRevisions adds the "revisions" edges to the ApplicationRevision entity.
func (aiuo *ApplicationInstanceUpdateOne) AddRevisions(a ...*ApplicationRevision) *ApplicationInstanceUpdateOne {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiuo.AddRevisionIDs(ids...)
}

// AddResourceIDs adds the "resources" edge to the ApplicationResource entity by IDs.
func (aiuo *ApplicationInstanceUpdateOne) AddResourceIDs(ids ...types.ID) *ApplicationInstanceUpdateOne {
	aiuo.mutation.AddResourceIDs(ids...)
	return aiuo
}

// AddResources adds the "resources" edges to the ApplicationResource entity.
func (aiuo *ApplicationInstanceUpdateOne) AddResources(a ...*ApplicationResource) *ApplicationInstanceUpdateOne {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiuo.AddResourceIDs(ids...)
}

// Mutation returns the ApplicationInstanceMutation object of the builder.
func (aiuo *ApplicationInstanceUpdateOne) Mutation() *ApplicationInstanceMutation {
	return aiuo.mutation
}

// ClearRevisions clears all "revisions" edges to the ApplicationRevision entity.
func (aiuo *ApplicationInstanceUpdateOne) ClearRevisions() *ApplicationInstanceUpdateOne {
	aiuo.mutation.ClearRevisions()
	return aiuo
}

// RemoveRevisionIDs removes the "revisions" edge to ApplicationRevision entities by IDs.
func (aiuo *ApplicationInstanceUpdateOne) RemoveRevisionIDs(ids ...types.ID) *ApplicationInstanceUpdateOne {
	aiuo.mutation.RemoveRevisionIDs(ids...)
	return aiuo
}

// RemoveRevisions removes "revisions" edges to ApplicationRevision entities.
func (aiuo *ApplicationInstanceUpdateOne) RemoveRevisions(a ...*ApplicationRevision) *ApplicationInstanceUpdateOne {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiuo.RemoveRevisionIDs(ids...)
}

// ClearResources clears all "resources" edges to the ApplicationResource entity.
func (aiuo *ApplicationInstanceUpdateOne) ClearResources() *ApplicationInstanceUpdateOne {
	aiuo.mutation.ClearResources()
	return aiuo
}

// RemoveResourceIDs removes the "resources" edge to ApplicationResource entities by IDs.
func (aiuo *ApplicationInstanceUpdateOne) RemoveResourceIDs(ids ...types.ID) *ApplicationInstanceUpdateOne {
	aiuo.mutation.RemoveResourceIDs(ids...)
	return aiuo
}

// RemoveResources removes "resources" edges to ApplicationResource entities.
func (aiuo *ApplicationInstanceUpdateOne) RemoveResources(a ...*ApplicationResource) *ApplicationInstanceUpdateOne {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aiuo.RemoveResourceIDs(ids...)
}

// Where appends a list predicates to the ApplicationInstanceUpdate builder.
func (aiuo *ApplicationInstanceUpdateOne) Where(ps ...predicate.ApplicationInstance) *ApplicationInstanceUpdateOne {
	aiuo.mutation.Where(ps...)
	return aiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *ApplicationInstanceUpdateOne) Select(field string, fields ...string) *ApplicationInstanceUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated ApplicationInstance entity.
func (aiuo *ApplicationInstanceUpdateOne) Save(ctx context.Context) (*ApplicationInstance, error) {
	if err := aiuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*ApplicationInstance, ApplicationInstanceMutation](ctx, aiuo.sqlSave, aiuo.mutation, aiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *ApplicationInstanceUpdateOne) SaveX(ctx context.Context) *ApplicationInstance {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *ApplicationInstanceUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *ApplicationInstanceUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiuo *ApplicationInstanceUpdateOne) defaults() error {
	if _, ok := aiuo.mutation.UpdateTime(); !ok {
		if applicationinstance.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized applicationinstance.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := applicationinstance.UpdateDefaultUpdateTime()
		aiuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aiuo *ApplicationInstanceUpdateOne) check() error {
	if _, ok := aiuo.mutation.ApplicationID(); aiuo.mutation.ApplicationCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationInstance.application"`)
	}
	if _, ok := aiuo.mutation.EnvironmentID(); aiuo.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationInstance.environment"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aiuo *ApplicationInstanceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApplicationInstanceUpdateOne {
	aiuo.modifiers = append(aiuo.modifiers, modifiers...)
	return aiuo
}

func (aiuo *ApplicationInstanceUpdateOne) sqlSave(ctx context.Context) (_node *ApplicationInstance, err error) {
	if err := aiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(applicationinstance.Table, applicationinstance.Columns, sqlgraph.NewFieldSpec(applicationinstance.FieldID, field.TypeString))
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ApplicationInstance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationinstance.FieldID)
		for _, f := range fields {
			if !applicationinstance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != applicationinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiuo.mutation.Status(); ok {
		_spec.SetField(applicationinstance.FieldStatus, field.TypeString, value)
	}
	if aiuo.mutation.StatusCleared() {
		_spec.ClearField(applicationinstance.FieldStatus, field.TypeString)
	}
	if value, ok := aiuo.mutation.StatusMessage(); ok {
		_spec.SetField(applicationinstance.FieldStatusMessage, field.TypeString, value)
	}
	if aiuo.mutation.StatusMessageCleared() {
		_spec.ClearField(applicationinstance.FieldStatusMessage, field.TypeString)
	}
	if value, ok := aiuo.mutation.UpdateTime(); ok {
		_spec.SetField(applicationinstance.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := aiuo.mutation.Variables(); ok {
		_spec.SetField(applicationinstance.FieldVariables, field.TypeJSON, value)
	}
	if aiuo.mutation.VariablesCleared() {
		_spec.ClearField(applicationinstance.FieldVariables, field.TypeJSON)
	}
	if aiuo.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.RevisionsTable,
			Columns: []string{applicationinstance.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationrevision.FieldID,
				},
			},
		}
		edge.Schema = aiuo.schemaConfig.ApplicationRevision
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RemovedRevisionsIDs(); len(nodes) > 0 && !aiuo.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.RevisionsTable,
			Columns: []string{applicationinstance.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationrevision.FieldID,
				},
			},
		}
		edge.Schema = aiuo.schemaConfig.ApplicationRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.RevisionsTable,
			Columns: []string{applicationinstance.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationrevision.FieldID,
				},
			},
		}
		edge.Schema = aiuo.schemaConfig.ApplicationRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiuo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.ResourcesTable,
			Columns: []string{applicationinstance.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aiuo.schemaConfig.ApplicationResource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !aiuo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.ResourcesTable,
			Columns: []string{applicationinstance.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aiuo.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationinstance.ResourcesTable,
			Columns: []string{applicationinstance.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aiuo.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = aiuo.schemaConfig.ApplicationInstance
	ctx = internal.NewSchemaConfigContext(ctx, aiuo.schemaConfig)
	_spec.AddModifiers(aiuo.modifiers...)
	_node = &ApplicationInstance{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aiuo.mutation.done = true
	return _node, nil
}