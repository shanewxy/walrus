// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/clustercost"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// ClusterCostUpdate is the builder for updating ClusterCost entities.
type ClusterCostUpdate struct {
	config
	hooks     []Hook
	mutation  *ClusterCostMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ClusterCost
}

// Where appends a list predicates to the ClusterCostUpdate builder.
func (ccu *ClusterCostUpdate) Where(ps ...predicate.ClusterCost) *ClusterCostUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetTotalCost sets the "total_cost" field.
func (ccu *ClusterCostUpdate) SetTotalCost(f float64) *ClusterCostUpdate {
	ccu.mutation.ResetTotalCost()
	ccu.mutation.SetTotalCost(f)
	return ccu
}

// SetNillableTotalCost sets the "total_cost" field if the given value is not nil.
func (ccu *ClusterCostUpdate) SetNillableTotalCost(f *float64) *ClusterCostUpdate {
	if f != nil {
		ccu.SetTotalCost(*f)
	}
	return ccu
}

// AddTotalCost adds f to the "total_cost" field.
func (ccu *ClusterCostUpdate) AddTotalCost(f float64) *ClusterCostUpdate {
	ccu.mutation.AddTotalCost(f)
	return ccu
}

// SetCurrency sets the "currency" field.
func (ccu *ClusterCostUpdate) SetCurrency(i int) *ClusterCostUpdate {
	ccu.mutation.ResetCurrency()
	ccu.mutation.SetCurrency(i)
	return ccu
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (ccu *ClusterCostUpdate) SetNillableCurrency(i *int) *ClusterCostUpdate {
	if i != nil {
		ccu.SetCurrency(*i)
	}
	return ccu
}

// AddCurrency adds i to the "currency" field.
func (ccu *ClusterCostUpdate) AddCurrency(i int) *ClusterCostUpdate {
	ccu.mutation.AddCurrency(i)
	return ccu
}

// ClearCurrency clears the value of the "currency" field.
func (ccu *ClusterCostUpdate) ClearCurrency() *ClusterCostUpdate {
	ccu.mutation.ClearCurrency()
	return ccu
}

// SetAllocationCost sets the "allocation_cost" field.
func (ccu *ClusterCostUpdate) SetAllocationCost(f float64) *ClusterCostUpdate {
	ccu.mutation.ResetAllocationCost()
	ccu.mutation.SetAllocationCost(f)
	return ccu
}

// SetNillableAllocationCost sets the "allocation_cost" field if the given value is not nil.
func (ccu *ClusterCostUpdate) SetNillableAllocationCost(f *float64) *ClusterCostUpdate {
	if f != nil {
		ccu.SetAllocationCost(*f)
	}
	return ccu
}

// AddAllocationCost adds f to the "allocation_cost" field.
func (ccu *ClusterCostUpdate) AddAllocationCost(f float64) *ClusterCostUpdate {
	ccu.mutation.AddAllocationCost(f)
	return ccu
}

// SetIdleCost sets the "idle_cost" field.
func (ccu *ClusterCostUpdate) SetIdleCost(f float64) *ClusterCostUpdate {
	ccu.mutation.ResetIdleCost()
	ccu.mutation.SetIdleCost(f)
	return ccu
}

// SetNillableIdleCost sets the "idle_cost" field if the given value is not nil.
func (ccu *ClusterCostUpdate) SetNillableIdleCost(f *float64) *ClusterCostUpdate {
	if f != nil {
		ccu.SetIdleCost(*f)
	}
	return ccu
}

// AddIdleCost adds f to the "idle_cost" field.
func (ccu *ClusterCostUpdate) AddIdleCost(f float64) *ClusterCostUpdate {
	ccu.mutation.AddIdleCost(f)
	return ccu
}

// SetManagementCost sets the "management_cost" field.
func (ccu *ClusterCostUpdate) SetManagementCost(f float64) *ClusterCostUpdate {
	ccu.mutation.ResetManagementCost()
	ccu.mutation.SetManagementCost(f)
	return ccu
}

// SetNillableManagementCost sets the "management_cost" field if the given value is not nil.
func (ccu *ClusterCostUpdate) SetNillableManagementCost(f *float64) *ClusterCostUpdate {
	if f != nil {
		ccu.SetManagementCost(*f)
	}
	return ccu
}

// AddManagementCost adds f to the "management_cost" field.
func (ccu *ClusterCostUpdate) AddManagementCost(f float64) *ClusterCostUpdate {
	ccu.mutation.AddManagementCost(f)
	return ccu
}

// Mutation returns the ClusterCostMutation object of the builder.
func (ccu *ClusterCostUpdate) Mutation() *ClusterCostMutation {
	return ccu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *ClusterCostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ccu.sqlSave, ccu.mutation, ccu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *ClusterCostUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *ClusterCostUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *ClusterCostUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccu *ClusterCostUpdate) check() error {
	if v, ok := ccu.mutation.TotalCost(); ok {
		if err := clustercost.TotalCostValidator(v); err != nil {
			return &ValidationError{Name: "total_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.total_cost": %w`, err)}
		}
	}
	if v, ok := ccu.mutation.AllocationCost(); ok {
		if err := clustercost.AllocationCostValidator(v); err != nil {
			return &ValidationError{Name: "allocation_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.allocation_cost": %w`, err)}
		}
	}
	if v, ok := ccu.mutation.IdleCost(); ok {
		if err := clustercost.IdleCostValidator(v); err != nil {
			return &ValidationError{Name: "idle_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.idle_cost": %w`, err)}
		}
	}
	if v, ok := ccu.mutation.ManagementCost(); ok {
		if err := clustercost.ManagementCostValidator(v); err != nil {
			return &ValidationError{Name: "management_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.management_cost": %w`, err)}
		}
	}
	if _, ok := ccu.mutation.ConnectorID(); ccu.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ClusterCost.connector"`)
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
func (ccu *ClusterCostUpdate) Set(obj *ClusterCost) *ClusterCostUpdate {
	// Without Default.
	ccu.SetTotalCost(obj.TotalCost)
	if obj.Currency != 0 {
		ccu.SetCurrency(obj.Currency)
	}
	ccu.SetAllocationCost(obj.AllocationCost)
	ccu.SetIdleCost(obj.IdleCost)
	ccu.SetManagementCost(obj.ManagementCost)

	// With Default.

	// Record the given object.
	ccu.object = obj

	return ccu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ccu *ClusterCostUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ClusterCostUpdate {
	ccu.modifiers = append(ccu.modifiers, modifiers...)
	return ccu
}

func (ccu *ClusterCostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ccu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(clustercost.Table, clustercost.Columns, sqlgraph.NewFieldSpec(clustercost.FieldID, field.TypeInt))
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.TotalCost(); ok {
		_spec.SetField(clustercost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AddedTotalCost(); ok {
		_spec.AddField(clustercost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.Currency(); ok {
		_spec.SetField(clustercost.FieldCurrency, field.TypeInt, value)
	}
	if value, ok := ccu.mutation.AddedCurrency(); ok {
		_spec.AddField(clustercost.FieldCurrency, field.TypeInt, value)
	}
	if ccu.mutation.CurrencyCleared() {
		_spec.ClearField(clustercost.FieldCurrency, field.TypeInt)
	}
	if value, ok := ccu.mutation.AllocationCost(); ok {
		_spec.SetField(clustercost.FieldAllocationCost, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AddedAllocationCost(); ok {
		_spec.AddField(clustercost.FieldAllocationCost, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.IdleCost(); ok {
		_spec.SetField(clustercost.FieldIdleCost, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AddedIdleCost(); ok {
		_spec.AddField(clustercost.FieldIdleCost, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.ManagementCost(); ok {
		_spec.SetField(clustercost.FieldManagementCost, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AddedManagementCost(); ok {
		_spec.AddField(clustercost.FieldManagementCost, field.TypeFloat64, value)
	}
	_spec.Node.Schema = ccu.schemaConfig.ClusterCost
	ctx = internal.NewSchemaConfigContext(ctx, ccu.schemaConfig)
	_spec.AddModifiers(ccu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clustercost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ccu.mutation.done = true
	return n, nil
}

// ClusterCostUpdateOne is the builder for updating a single ClusterCost entity.
type ClusterCostUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ClusterCostMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ClusterCost
}

// SetTotalCost sets the "total_cost" field.
func (ccuo *ClusterCostUpdateOne) SetTotalCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.ResetTotalCost()
	ccuo.mutation.SetTotalCost(f)
	return ccuo
}

// SetNillableTotalCost sets the "total_cost" field if the given value is not nil.
func (ccuo *ClusterCostUpdateOne) SetNillableTotalCost(f *float64) *ClusterCostUpdateOne {
	if f != nil {
		ccuo.SetTotalCost(*f)
	}
	return ccuo
}

// AddTotalCost adds f to the "total_cost" field.
func (ccuo *ClusterCostUpdateOne) AddTotalCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.AddTotalCost(f)
	return ccuo
}

// SetCurrency sets the "currency" field.
func (ccuo *ClusterCostUpdateOne) SetCurrency(i int) *ClusterCostUpdateOne {
	ccuo.mutation.ResetCurrency()
	ccuo.mutation.SetCurrency(i)
	return ccuo
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (ccuo *ClusterCostUpdateOne) SetNillableCurrency(i *int) *ClusterCostUpdateOne {
	if i != nil {
		ccuo.SetCurrency(*i)
	}
	return ccuo
}

// AddCurrency adds i to the "currency" field.
func (ccuo *ClusterCostUpdateOne) AddCurrency(i int) *ClusterCostUpdateOne {
	ccuo.mutation.AddCurrency(i)
	return ccuo
}

// ClearCurrency clears the value of the "currency" field.
func (ccuo *ClusterCostUpdateOne) ClearCurrency() *ClusterCostUpdateOne {
	ccuo.mutation.ClearCurrency()
	return ccuo
}

// SetAllocationCost sets the "allocation_cost" field.
func (ccuo *ClusterCostUpdateOne) SetAllocationCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.ResetAllocationCost()
	ccuo.mutation.SetAllocationCost(f)
	return ccuo
}

// SetNillableAllocationCost sets the "allocation_cost" field if the given value is not nil.
func (ccuo *ClusterCostUpdateOne) SetNillableAllocationCost(f *float64) *ClusterCostUpdateOne {
	if f != nil {
		ccuo.SetAllocationCost(*f)
	}
	return ccuo
}

// AddAllocationCost adds f to the "allocation_cost" field.
func (ccuo *ClusterCostUpdateOne) AddAllocationCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.AddAllocationCost(f)
	return ccuo
}

// SetIdleCost sets the "idle_cost" field.
func (ccuo *ClusterCostUpdateOne) SetIdleCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.ResetIdleCost()
	ccuo.mutation.SetIdleCost(f)
	return ccuo
}

// SetNillableIdleCost sets the "idle_cost" field if the given value is not nil.
func (ccuo *ClusterCostUpdateOne) SetNillableIdleCost(f *float64) *ClusterCostUpdateOne {
	if f != nil {
		ccuo.SetIdleCost(*f)
	}
	return ccuo
}

// AddIdleCost adds f to the "idle_cost" field.
func (ccuo *ClusterCostUpdateOne) AddIdleCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.AddIdleCost(f)
	return ccuo
}

// SetManagementCost sets the "management_cost" field.
func (ccuo *ClusterCostUpdateOne) SetManagementCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.ResetManagementCost()
	ccuo.mutation.SetManagementCost(f)
	return ccuo
}

// SetNillableManagementCost sets the "management_cost" field if the given value is not nil.
func (ccuo *ClusterCostUpdateOne) SetNillableManagementCost(f *float64) *ClusterCostUpdateOne {
	if f != nil {
		ccuo.SetManagementCost(*f)
	}
	return ccuo
}

// AddManagementCost adds f to the "management_cost" field.
func (ccuo *ClusterCostUpdateOne) AddManagementCost(f float64) *ClusterCostUpdateOne {
	ccuo.mutation.AddManagementCost(f)
	return ccuo
}

// Mutation returns the ClusterCostMutation object of the builder.
func (ccuo *ClusterCostUpdateOne) Mutation() *ClusterCostMutation {
	return ccuo.mutation
}

// Where appends a list predicates to the ClusterCostUpdate builder.
func (ccuo *ClusterCostUpdateOne) Where(ps ...predicate.ClusterCost) *ClusterCostUpdateOne {
	ccuo.mutation.Where(ps...)
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *ClusterCostUpdateOne) Select(field string, fields ...string) *ClusterCostUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated ClusterCost entity.
func (ccuo *ClusterCostUpdateOne) Save(ctx context.Context) (*ClusterCost, error) {
	return withHooks(ctx, ccuo.sqlSave, ccuo.mutation, ccuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *ClusterCostUpdateOne) SaveX(ctx context.Context) *ClusterCost {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *ClusterCostUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *ClusterCostUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccuo *ClusterCostUpdateOne) check() error {
	if v, ok := ccuo.mutation.TotalCost(); ok {
		if err := clustercost.TotalCostValidator(v); err != nil {
			return &ValidationError{Name: "total_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.total_cost": %w`, err)}
		}
	}
	if v, ok := ccuo.mutation.AllocationCost(); ok {
		if err := clustercost.AllocationCostValidator(v); err != nil {
			return &ValidationError{Name: "allocation_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.allocation_cost": %w`, err)}
		}
	}
	if v, ok := ccuo.mutation.IdleCost(); ok {
		if err := clustercost.IdleCostValidator(v); err != nil {
			return &ValidationError{Name: "idle_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.idle_cost": %w`, err)}
		}
	}
	if v, ok := ccuo.mutation.ManagementCost(); ok {
		if err := clustercost.ManagementCostValidator(v); err != nil {
			return &ValidationError{Name: "management_cost", err: fmt.Errorf(`model: validator failed for field "ClusterCost.management_cost": %w`, err)}
		}
	}
	if _, ok := ccuo.mutation.ConnectorID(); ccuo.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ClusterCost.connector"`)
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
func (ccuo *ClusterCostUpdateOne) Set(obj *ClusterCost) *ClusterCostUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*ClusterCostMutation)
			db, err := mt.Client().ClusterCost.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting ClusterCost with id: %v", *mt.id)
			}

			// Without Default.
			if db.TotalCost != obj.TotalCost {
				ccuo.SetTotalCost(obj.TotalCost)
			}
			if obj.Currency != 0 {
				if db.Currency != obj.Currency {
					ccuo.SetCurrency(obj.Currency)
				}
			}
			if db.AllocationCost != obj.AllocationCost {
				ccuo.SetAllocationCost(obj.AllocationCost)
			}
			if db.IdleCost != obj.IdleCost {
				ccuo.SetIdleCost(obj.IdleCost)
			}
			if db.ManagementCost != obj.ManagementCost {
				ccuo.SetManagementCost(obj.ManagementCost)
			}

			// With Default.

			// Record the given object.
			ccuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	ccuo.hooks = append(ccuo.hooks, h)

	return ccuo
}

// getClientSet returns the ClientSet for the given builder.
func (ccuo *ClusterCostUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := ccuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: ccuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: ccuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the ClusterCost entity,
// which is always good for cascading update operations.
func (ccuo *ClusterCostUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ClusterCost) error) (*ClusterCost, error) {
	obj, err := ccuo.Save(ctx)
	if err != nil &&
		(ccuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := ccuo.getClientSet()

	if obj == nil {
		obj = ccuo.object
	} else if x := ccuo.object; x != nil {
		if _, set := ccuo.mutation.Field(clustercost.FieldTotalCost); set {
			obj.TotalCost = x.TotalCost
		}
		if _, set := ccuo.mutation.Field(clustercost.FieldCurrency); set {
			obj.Currency = x.Currency
		}
		if _, set := ccuo.mutation.Field(clustercost.FieldAllocationCost); set {
			obj.AllocationCost = x.AllocationCost
		}
		if _, set := ccuo.mutation.Field(clustercost.FieldIdleCost); set {
			obj.IdleCost = x.IdleCost
		}
		if _, set := ccuo.mutation.Field(clustercost.FieldManagementCost); set {
			obj.ManagementCost = x.ManagementCost
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
func (ccuo *ClusterCostUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ClusterCost) error) *ClusterCost {
	obj, err := ccuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (ccuo *ClusterCostUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ClusterCost) error) error {
	_, err := ccuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *ClusterCostUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ClusterCost) error) {
	if err := ccuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ccuo *ClusterCostUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ClusterCostUpdateOne {
	ccuo.modifiers = append(ccuo.modifiers, modifiers...)
	return ccuo
}

func (ccuo *ClusterCostUpdateOne) sqlSave(ctx context.Context) (_node *ClusterCost, err error) {
	if err := ccuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(clustercost.Table, clustercost.Columns, sqlgraph.NewFieldSpec(clustercost.FieldID, field.TypeInt))
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ClusterCost.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clustercost.FieldID)
		for _, f := range fields {
			if !clustercost.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != clustercost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.TotalCost(); ok {
		_spec.SetField(clustercost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AddedTotalCost(); ok {
		_spec.AddField(clustercost.FieldTotalCost, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.Currency(); ok {
		_spec.SetField(clustercost.FieldCurrency, field.TypeInt, value)
	}
	if value, ok := ccuo.mutation.AddedCurrency(); ok {
		_spec.AddField(clustercost.FieldCurrency, field.TypeInt, value)
	}
	if ccuo.mutation.CurrencyCleared() {
		_spec.ClearField(clustercost.FieldCurrency, field.TypeInt)
	}
	if value, ok := ccuo.mutation.AllocationCost(); ok {
		_spec.SetField(clustercost.FieldAllocationCost, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AddedAllocationCost(); ok {
		_spec.AddField(clustercost.FieldAllocationCost, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.IdleCost(); ok {
		_spec.SetField(clustercost.FieldIdleCost, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AddedIdleCost(); ok {
		_spec.AddField(clustercost.FieldIdleCost, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.ManagementCost(); ok {
		_spec.SetField(clustercost.FieldManagementCost, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AddedManagementCost(); ok {
		_spec.AddField(clustercost.FieldManagementCost, field.TypeFloat64, value)
	}
	_spec.Node.Schema = ccuo.schemaConfig.ClusterCost
	ctx = internal.NewSchemaConfigContext(ctx, ccuo.schemaConfig)
	_spec.AddModifiers(ccuo.modifiers...)
	_node = &ClusterCost{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clustercost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ccuo.mutation.done = true
	return _node, nil
}
