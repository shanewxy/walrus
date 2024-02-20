// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcestate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ResourceState is the model entity for the ResourceState schema.
type ResourceState struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// State data of the resource.
	Data string `json:"data,omitempty"`
	// ID of the resource to which the state belongs.
	ResourceID object.ID `json:"resource_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceStateQuery when eager-loading is set.
	Edges        ResourceStateEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ResourceStateEdges holds the relations/edges for other nodes in the graph.
type ResourceStateEdges struct {
	// Resource to which the state belongs.
	Resource *Resource `json:"resource,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceStateEdges) ResourceOrErr() (*Resource, error) {
	if e.loadedTypes[0] {
		if e.Resource == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resource.Label}
		}
		return e.Resource, nil
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResourceState) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resourcestate.FieldID, resourcestate.FieldResourceID:
			values[i] = new(object.ID)
		case resourcestate.FieldData:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourceState fields.
func (rs *ResourceState) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resourcestate.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rs.ID = *value
			}
		case resourcestate.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				rs.Data = value.String
			}
		case resourcestate.FieldResourceID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value != nil {
				rs.ResourceID = *value
			}
		default:
			rs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResourceState.
// This includes values selected through modifiers, order, etc.
func (rs *ResourceState) Value(name string) (ent.Value, error) {
	return rs.selectValues.Get(name)
}

// QueryResource queries the "resource" edge of the ResourceState entity.
func (rs *ResourceState) QueryResource() *ResourceQuery {
	return NewResourceStateClient(rs.config).QueryResource(rs)
}

// Update returns a builder for updating this ResourceState.
// Note that you need to call ResourceState.Unwrap() before calling this method if this ResourceState
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *ResourceState) Update() *ResourceStateUpdateOne {
	return NewResourceStateClient(rs.config).UpdateOne(rs)
}

// Unwrap unwraps the ResourceState entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rs *ResourceState) Unwrap() *ResourceState {
	_tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("model: ResourceState is not a transactional entity")
	}
	rs.config.driver = _tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *ResourceState) String() string {
	var builder strings.Builder
	builder.WriteString("ResourceState(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rs.ID))
	builder.WriteString("data=")
	builder.WriteString(rs.Data)
	builder.WriteString(", ")
	builder.WriteString("resource_id=")
	builder.WriteString(fmt.Sprintf("%v", rs.ResourceID))
	builder.WriteByte(')')
	return builder.String()
}

// ResourceStates is a parsable slice of ResourceState.
type ResourceStates []*ResourceState