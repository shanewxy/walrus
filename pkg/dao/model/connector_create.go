// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

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

	"github.com/seal-io/seal/pkg/dao/model/connector"
	"github.com/seal-io/seal/pkg/dao/oid"
)

// ConnectorCreate is the builder for creating a Connector entity.
type ConnectorCreate struct {
	config
	mutation *ConnectorMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStatus sets the "status" field.
func (cc *ConnectorCreate) SetStatus(s string) *ConnectorCreate {
	cc.mutation.SetStatus(s)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableStatus(s *string) *ConnectorCreate {
	if s != nil {
		cc.SetStatus(*s)
	}
	return cc
}

// SetStatusMessage sets the "statusMessage" field.
func (cc *ConnectorCreate) SetStatusMessage(s string) *ConnectorCreate {
	cc.mutation.SetStatusMessage(s)
	return cc
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableStatusMessage(s *string) *ConnectorCreate {
	if s != nil {
		cc.SetStatusMessage(*s)
	}
	return cc
}

// SetCreateTime sets the "createTime" field.
func (cc *ConnectorCreate) SetCreateTime(t time.Time) *ConnectorCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableCreateTime(t *time.Time) *ConnectorCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "updateTime" field.
func (cc *ConnectorCreate) SetUpdateTime(t time.Time) *ConnectorCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableUpdateTime(t *time.Time) *ConnectorCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetDriver sets the "driver" field.
func (cc *ConnectorCreate) SetDriver(s string) *ConnectorCreate {
	cc.mutation.SetDriver(s)
	return cc
}

// SetConfigVersion sets the "configVersion" field.
func (cc *ConnectorCreate) SetConfigVersion(s string) *ConnectorCreate {
	cc.mutation.SetConfigVersion(s)
	return cc
}

// SetConfigData sets the "configData" field.
func (cc *ConnectorCreate) SetConfigData(m map[string]interface{}) *ConnectorCreate {
	cc.mutation.SetConfigData(m)
	return cc
}

// SetID sets the "id" field.
func (cc *ConnectorCreate) SetID(o oid.ID) *ConnectorCreate {
	cc.mutation.SetID(o)
	return cc
}

// Mutation returns the ConnectorMutation object of the builder.
func (cc *ConnectorCreate) Mutation() *ConnectorMutation {
	return cc.mutation
}

// Save creates the Connector in the database.
func (cc *ConnectorCreate) Save(ctx context.Context) (*Connector, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Connector, ConnectorMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConnectorCreate) SaveX(ctx context.Context) *Connector {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConnectorCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConnectorCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ConnectorCreate) defaults() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		if connector.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized connector.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := connector.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		if connector.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized connector.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := connector.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConnectorCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "createTime", err: errors.New(`model: missing required field "Connector.createTime"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "updateTime", err: errors.New(`model: missing required field "Connector.updateTime"`)}
	}
	if _, ok := cc.mutation.Driver(); !ok {
		return &ValidationError{Name: "driver", err: errors.New(`model: missing required field "Connector.driver"`)}
	}
	if _, ok := cc.mutation.ConfigVersion(); !ok {
		return &ValidationError{Name: "configVersion", err: errors.New(`model: missing required field "Connector.configVersion"`)}
	}
	return nil
}

func (cc *ConnectorCreate) sqlSave(ctx context.Context) (*Connector, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*oid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ConnectorCreate) createSpec() (*Connector, *sqlgraph.CreateSpec) {
	var (
		_node = &Connector{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: connector.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: connector.FieldID,
			},
		}
	)
	_spec.Schema = cc.schemaConfig.Connector
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(connector.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.StatusMessage(); ok {
		_spec.SetField(connector.FieldStatusMessage, field.TypeString, value)
		_node.StatusMessage = value
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(connector.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(connector.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := cc.mutation.Driver(); ok {
		_spec.SetField(connector.FieldDriver, field.TypeString, value)
		_node.Driver = value
	}
	if value, ok := cc.mutation.ConfigVersion(); ok {
		_spec.SetField(connector.FieldConfigVersion, field.TypeString, value)
		_node.ConfigVersion = value
	}
	if value, ok := cc.mutation.ConfigData(); ok {
		_spec.SetField(connector.FieldConfigData, field.TypeJSON, value)
		_node.ConfigData = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Connector.Create().
//		SetStatus(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ConnectorUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (cc *ConnectorCreate) OnConflict(opts ...sql.ConflictOption) *ConnectorUpsertOne {
	cc.conflict = opts
	return &ConnectorUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Connector.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ConnectorCreate) OnConflictColumns(columns ...string) *ConnectorUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ConnectorUpsertOne{
		create: cc,
	}
}

type (
	// ConnectorUpsertOne is the builder for "upsert"-ing
	//  one Connector node.
	ConnectorUpsertOne struct {
		create *ConnectorCreate
	}

	// ConnectorUpsert is the "OnConflict" setter.
	ConnectorUpsert struct {
		*sql.UpdateSet
	}
)

// SetStatus sets the "status" field.
func (u *ConnectorUpsert) SetStatus(v string) *ConnectorUpsert {
	u.Set(connector.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ConnectorUpsert) UpdateStatus() *ConnectorUpsert {
	u.SetExcluded(connector.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *ConnectorUpsert) ClearStatus() *ConnectorUpsert {
	u.SetNull(connector.FieldStatus)
	return u
}

// SetStatusMessage sets the "statusMessage" field.
func (u *ConnectorUpsert) SetStatusMessage(v string) *ConnectorUpsert {
	u.Set(connector.FieldStatusMessage, v)
	return u
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *ConnectorUpsert) UpdateStatusMessage() *ConnectorUpsert {
	u.SetExcluded(connector.FieldStatusMessage)
	return u
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *ConnectorUpsert) ClearStatusMessage() *ConnectorUpsert {
	u.SetNull(connector.FieldStatusMessage)
	return u
}

// SetUpdateTime sets the "updateTime" field.
func (u *ConnectorUpsert) SetUpdateTime(v time.Time) *ConnectorUpsert {
	u.Set(connector.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *ConnectorUpsert) UpdateUpdateTime() *ConnectorUpsert {
	u.SetExcluded(connector.FieldUpdateTime)
	return u
}

// SetDriver sets the "driver" field.
func (u *ConnectorUpsert) SetDriver(v string) *ConnectorUpsert {
	u.Set(connector.FieldDriver, v)
	return u
}

// UpdateDriver sets the "driver" field to the value that was provided on create.
func (u *ConnectorUpsert) UpdateDriver() *ConnectorUpsert {
	u.SetExcluded(connector.FieldDriver)
	return u
}

// SetConfigVersion sets the "configVersion" field.
func (u *ConnectorUpsert) SetConfigVersion(v string) *ConnectorUpsert {
	u.Set(connector.FieldConfigVersion, v)
	return u
}

// UpdateConfigVersion sets the "configVersion" field to the value that was provided on create.
func (u *ConnectorUpsert) UpdateConfigVersion() *ConnectorUpsert {
	u.SetExcluded(connector.FieldConfigVersion)
	return u
}

// SetConfigData sets the "configData" field.
func (u *ConnectorUpsert) SetConfigData(v map[string]interface{}) *ConnectorUpsert {
	u.Set(connector.FieldConfigData, v)
	return u
}

// UpdateConfigData sets the "configData" field to the value that was provided on create.
func (u *ConnectorUpsert) UpdateConfigData() *ConnectorUpsert {
	u.SetExcluded(connector.FieldConfigData)
	return u
}

// ClearConfigData clears the value of the "configData" field.
func (u *ConnectorUpsert) ClearConfigData() *ConnectorUpsert {
	u.SetNull(connector.FieldConfigData)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Connector.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(connector.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ConnectorUpsertOne) UpdateNewValues() *ConnectorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(connector.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(connector.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Connector.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ConnectorUpsertOne) Ignore() *ConnectorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ConnectorUpsertOne) DoNothing() *ConnectorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ConnectorCreate.OnConflict
// documentation for more info.
func (u *ConnectorUpsertOne) Update(set func(*ConnectorUpsert)) *ConnectorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ConnectorUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *ConnectorUpsertOne) SetStatus(v string) *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ConnectorUpsertOne) UpdateStatus() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *ConnectorUpsertOne) ClearStatus() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.ClearStatus()
	})
}

// SetStatusMessage sets the "statusMessage" field.
func (u *ConnectorUpsertOne) SetStatusMessage(v string) *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetStatusMessage(v)
	})
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *ConnectorUpsertOne) UpdateStatusMessage() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateStatusMessage()
	})
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *ConnectorUpsertOne) ClearStatusMessage() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.ClearStatusMessage()
	})
}

// SetUpdateTime sets the "updateTime" field.
func (u *ConnectorUpsertOne) SetUpdateTime(v time.Time) *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *ConnectorUpsertOne) UpdateUpdateTime() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDriver sets the "driver" field.
func (u *ConnectorUpsertOne) SetDriver(v string) *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetDriver(v)
	})
}

// UpdateDriver sets the "driver" field to the value that was provided on create.
func (u *ConnectorUpsertOne) UpdateDriver() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateDriver()
	})
}

// SetConfigVersion sets the "configVersion" field.
func (u *ConnectorUpsertOne) SetConfigVersion(v string) *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetConfigVersion(v)
	})
}

// UpdateConfigVersion sets the "configVersion" field to the value that was provided on create.
func (u *ConnectorUpsertOne) UpdateConfigVersion() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateConfigVersion()
	})
}

// SetConfigData sets the "configData" field.
func (u *ConnectorUpsertOne) SetConfigData(v map[string]interface{}) *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetConfigData(v)
	})
}

// UpdateConfigData sets the "configData" field to the value that was provided on create.
func (u *ConnectorUpsertOne) UpdateConfigData() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateConfigData()
	})
}

// ClearConfigData clears the value of the "configData" field.
func (u *ConnectorUpsertOne) ClearConfigData() *ConnectorUpsertOne {
	return u.Update(func(s *ConnectorUpsert) {
		s.ClearConfigData()
	})
}

// Exec executes the query.
func (u *ConnectorUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ConnectorCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ConnectorUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ConnectorUpsertOne) ID(ctx context.Context) (id oid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: ConnectorUpsertOne.ID is not supported by MySQL driver. Use ConnectorUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ConnectorUpsertOne) IDX(ctx context.Context) oid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ConnectorCreateBulk is the builder for creating many Connector entities in bulk.
type ConnectorCreateBulk struct {
	config
	builders []*ConnectorCreate
	conflict []sql.ConflictOption
}

// Save creates the Connector entities in the database.
func (ccb *ConnectorCreateBulk) Save(ctx context.Context) ([]*Connector, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Connector, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConnectorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConnectorCreateBulk) SaveX(ctx context.Context) []*Connector {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConnectorCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConnectorCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Connector.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ConnectorUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (ccb *ConnectorCreateBulk) OnConflict(opts ...sql.ConflictOption) *ConnectorUpsertBulk {
	ccb.conflict = opts
	return &ConnectorUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Connector.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ConnectorCreateBulk) OnConflictColumns(columns ...string) *ConnectorUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ConnectorUpsertBulk{
		create: ccb,
	}
}

// ConnectorUpsertBulk is the builder for "upsert"-ing
// a bulk of Connector nodes.
type ConnectorUpsertBulk struct {
	create *ConnectorCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Connector.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(connector.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ConnectorUpsertBulk) UpdateNewValues() *ConnectorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(connector.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(connector.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Connector.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ConnectorUpsertBulk) Ignore() *ConnectorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ConnectorUpsertBulk) DoNothing() *ConnectorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ConnectorCreateBulk.OnConflict
// documentation for more info.
func (u *ConnectorUpsertBulk) Update(set func(*ConnectorUpsert)) *ConnectorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ConnectorUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *ConnectorUpsertBulk) SetStatus(v string) *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ConnectorUpsertBulk) UpdateStatus() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *ConnectorUpsertBulk) ClearStatus() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.ClearStatus()
	})
}

// SetStatusMessage sets the "statusMessage" field.
func (u *ConnectorUpsertBulk) SetStatusMessage(v string) *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetStatusMessage(v)
	})
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *ConnectorUpsertBulk) UpdateStatusMessage() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateStatusMessage()
	})
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *ConnectorUpsertBulk) ClearStatusMessage() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.ClearStatusMessage()
	})
}

// SetUpdateTime sets the "updateTime" field.
func (u *ConnectorUpsertBulk) SetUpdateTime(v time.Time) *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *ConnectorUpsertBulk) UpdateUpdateTime() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDriver sets the "driver" field.
func (u *ConnectorUpsertBulk) SetDriver(v string) *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetDriver(v)
	})
}

// UpdateDriver sets the "driver" field to the value that was provided on create.
func (u *ConnectorUpsertBulk) UpdateDriver() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateDriver()
	})
}

// SetConfigVersion sets the "configVersion" field.
func (u *ConnectorUpsertBulk) SetConfigVersion(v string) *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetConfigVersion(v)
	})
}

// UpdateConfigVersion sets the "configVersion" field to the value that was provided on create.
func (u *ConnectorUpsertBulk) UpdateConfigVersion() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateConfigVersion()
	})
}

// SetConfigData sets the "configData" field.
func (u *ConnectorUpsertBulk) SetConfigData(v map[string]interface{}) *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.SetConfigData(v)
	})
}

// UpdateConfigData sets the "configData" field to the value that was provided on create.
func (u *ConnectorUpsertBulk) UpdateConfigData() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.UpdateConfigData()
	})
}

// ClearConfigData clears the value of the "configData" field.
func (u *ConnectorUpsertBulk) ClearConfigData() *ConnectorUpsertBulk {
	return u.Update(func(s *ConnectorUpsert) {
		s.ClearConfigData()
	})
}

// Exec executes the query.
func (u *ConnectorUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ConnectorCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ConnectorCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ConnectorUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}