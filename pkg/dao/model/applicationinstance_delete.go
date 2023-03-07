// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/applicationinstance"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// ApplicationInstanceDelete is the builder for deleting a ApplicationInstance entity.
type ApplicationInstanceDelete struct {
	config
	hooks    []Hook
	mutation *ApplicationInstanceMutation
}

// Where appends a list predicates to the ApplicationInstanceDelete builder.
func (aid *ApplicationInstanceDelete) Where(ps ...predicate.ApplicationInstance) *ApplicationInstanceDelete {
	aid.mutation.Where(ps...)
	return aid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aid *ApplicationInstanceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ApplicationInstanceMutation](ctx, aid.sqlExec, aid.mutation, aid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (aid *ApplicationInstanceDelete) ExecX(ctx context.Context) int {
	n, err := aid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aid *ApplicationInstanceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(applicationinstance.Table, sqlgraph.NewFieldSpec(applicationinstance.FieldID, field.TypeString))
	_spec.Node.Schema = aid.schemaConfig.ApplicationInstance
	ctx = internal.NewSchemaConfigContext(ctx, aid.schemaConfig)
	if ps := aid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	aid.mutation.done = true
	return affected, err
}

// ApplicationInstanceDeleteOne is the builder for deleting a single ApplicationInstance entity.
type ApplicationInstanceDeleteOne struct {
	aid *ApplicationInstanceDelete
}

// Where appends a list predicates to the ApplicationInstanceDelete builder.
func (aido *ApplicationInstanceDeleteOne) Where(ps ...predicate.ApplicationInstance) *ApplicationInstanceDeleteOne {
	aido.aid.mutation.Where(ps...)
	return aido
}

// Exec executes the deletion query.
func (aido *ApplicationInstanceDeleteOne) Exec(ctx context.Context) error {
	n, err := aido.aid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{applicationinstance.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aido *ApplicationInstanceDeleteOne) ExecX(ctx context.Context) {
	if err := aido.Exec(ctx); err != nil {
		panic(err)
	}
}