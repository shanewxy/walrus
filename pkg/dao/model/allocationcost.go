// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/seal/pkg/dao/model/allocationcost"
	"github.com/seal-io/seal/pkg/dao/model/connector"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/object"
	"github.com/seal-io/seal/utils/json"
)

// AllocationCost is the model entity for the AllocationCost schema.
type AllocationCost struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Usage start time for current cost.
	StartTime time.Time `json:"start_time,omitempty"`
	// Usage end time for current cost.
	EndTime time.Time `json:"end_time,omitempty"`
	// Usage minutes from start time to end time.
	Minutes float64 `json:"minutes,omitempty"`
	// ID of the connector.
	ConnectorID object.ID `json:"connector_id,omitempty"`
	// Resource name for current cost, could be __unmounted__.
	Name string `json:"name,omitempty"`
	// String generated from resource properties, used to identify this cost.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Cluster name for current cost.
	ClusterName string `json:"cluster_name,omitempty"`
	// Namespace for current cost.
	Namespace string `json:"namespace,omitempty"`
	// Node for current cost.
	Node string `json:"node,omitempty"`
	// Controller name for the cost linked resource.
	Controller string `json:"controller,omitempty"`
	// Controller kind for the cost linked resource, deployment, statefulSet etc.
	ControllerKind string `json:"controller_kind,omitempty"`
	// Pod name for current cost.
	Pod string `json:"pod,omitempty"`
	// Container name for current cost.
	Container string `json:"container,omitempty"`
	// PV list for current cost linked.
	Pvs map[string]types.PVCost `json:"pvs,omitempty"`
	// Labels for the cost linked resource.
	Labels map[string]string `json:"labels,omitempty"`
	// Cost number.
	TotalCost float64 `json:"totalCost,omitempty"`
	// Cost currency.
	Currency int `json:"currency,omitempty"`
	// Cpu cost for current cost.
	CPUCost float64 `json:"cpu_cost,omitempty"`
	// Cpu core requested.
	CPUCoreRequest float64 `json:"cpu_core_request,omitempty"`
	// GPU cost for current cost.
	GpuCost float64 `json:"gpu_cost,omitempty"`
	// GPU core count.
	GpuCount float64 `json:"gpu_count,omitempty"`
	// Ram cost for current cost.
	RAMCost float64 `json:"ram_cost,omitempty"`
	// Ram requested in byte.
	RAMByteRequest float64 `json:"ram_byte_request,omitempty"`
	// PV cost for current cost linked.
	PvCost float64 `json:"pv_cost,omitempty"`
	// PV bytes for current cost linked.
	PvBytes float64 `json:"pv_bytes,omitempty"`
	// LoadBalancer cost for current cost linked.
	LoadBalancerCost float64 `json:"load_balancer_cost,omitempty"`
	// CPU core average usage.
	CPUCoreUsageAverage float64 `json:"cpu_core_usage_average,omitempty"`
	// CPU core max usage.
	CPUCoreUsageMax float64 `json:"cpu_core_usage_max,omitempty"`
	// Ram average usage in byte.
	RAMByteUsageAverage float64 `json:"ram_byte_usage_average,omitempty"`
	// Ram max usage in byte.
	RAMByteUsageMax float64 `json:"ram_byte_usage_max,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AllocationCostQuery when eager-loading is set.
	Edges        AllocationCostEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// AllocationCostEdges holds the relations/edges for other nodes in the graph.
type AllocationCostEdges struct {
	// Connector current cost linked.
	Connector *Connector `json:"connector,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ConnectorOrErr returns the Connector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AllocationCostEdges) ConnectorOrErr() (*Connector, error) {
	if e.loadedTypes[0] {
		if e.Connector == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: connector.Label}
		}
		return e.Connector, nil
	}
	return nil, &NotLoadedError{edge: "connector"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AllocationCost) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case allocationcost.FieldPvs, allocationcost.FieldLabels:
			values[i] = new([]byte)
		case allocationcost.FieldConnectorID:
			values[i] = new(object.ID)
		case allocationcost.FieldMinutes, allocationcost.FieldTotalCost, allocationcost.FieldCPUCost, allocationcost.FieldCPUCoreRequest, allocationcost.FieldGpuCost, allocationcost.FieldGpuCount, allocationcost.FieldRAMCost, allocationcost.FieldRAMByteRequest, allocationcost.FieldPvCost, allocationcost.FieldPvBytes, allocationcost.FieldLoadBalancerCost, allocationcost.FieldCPUCoreUsageAverage, allocationcost.FieldCPUCoreUsageMax, allocationcost.FieldRAMByteUsageAverage, allocationcost.FieldRAMByteUsageMax:
			values[i] = new(sql.NullFloat64)
		case allocationcost.FieldID, allocationcost.FieldCurrency:
			values[i] = new(sql.NullInt64)
		case allocationcost.FieldName, allocationcost.FieldFingerprint, allocationcost.FieldClusterName, allocationcost.FieldNamespace, allocationcost.FieldNode, allocationcost.FieldController, allocationcost.FieldControllerKind, allocationcost.FieldPod, allocationcost.FieldContainer:
			values[i] = new(sql.NullString)
		case allocationcost.FieldStartTime, allocationcost.FieldEndTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AllocationCost fields.
func (ac *AllocationCost) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case allocationcost.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		case allocationcost.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				ac.StartTime = value.Time
			}
		case allocationcost.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				ac.EndTime = value.Time
			}
		case allocationcost.FieldMinutes:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field minutes", values[i])
			} else if value.Valid {
				ac.Minutes = value.Float64
			}
		case allocationcost.FieldConnectorID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field connector_id", values[i])
			} else if value != nil {
				ac.ConnectorID = *value
			}
		case allocationcost.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ac.Name = value.String
			}
		case allocationcost.FieldFingerprint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fingerprint", values[i])
			} else if value.Valid {
				ac.Fingerprint = value.String
			}
		case allocationcost.FieldClusterName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cluster_name", values[i])
			} else if value.Valid {
				ac.ClusterName = value.String
			}
		case allocationcost.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				ac.Namespace = value.String
			}
		case allocationcost.FieldNode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node", values[i])
			} else if value.Valid {
				ac.Node = value.String
			}
		case allocationcost.FieldController:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field controller", values[i])
			} else if value.Valid {
				ac.Controller = value.String
			}
		case allocationcost.FieldControllerKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field controller_kind", values[i])
			} else if value.Valid {
				ac.ControllerKind = value.String
			}
		case allocationcost.FieldPod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pod", values[i])
			} else if value.Valid {
				ac.Pod = value.String
			}
		case allocationcost.FieldContainer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field container", values[i])
			} else if value.Valid {
				ac.Container = value.String
			}
		case allocationcost.FieldPvs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field pvs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ac.Pvs); err != nil {
					return fmt.Errorf("unmarshal field pvs: %w", err)
				}
			}
		case allocationcost.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ac.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case allocationcost.FieldTotalCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field totalCost", values[i])
			} else if value.Valid {
				ac.TotalCost = value.Float64
			}
		case allocationcost.FieldCurrency:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				ac.Currency = int(value.Int64)
			}
		case allocationcost.FieldCPUCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu_cost", values[i])
			} else if value.Valid {
				ac.CPUCost = value.Float64
			}
		case allocationcost.FieldCPUCoreRequest:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu_core_request", values[i])
			} else if value.Valid {
				ac.CPUCoreRequest = value.Float64
			}
		case allocationcost.FieldGpuCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_cost", values[i])
			} else if value.Valid {
				ac.GpuCost = value.Float64
			}
		case allocationcost.FieldGpuCount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_count", values[i])
			} else if value.Valid {
				ac.GpuCount = value.Float64
			}
		case allocationcost.FieldRAMCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ram_cost", values[i])
			} else if value.Valid {
				ac.RAMCost = value.Float64
			}
		case allocationcost.FieldRAMByteRequest:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ram_byte_request", values[i])
			} else if value.Valid {
				ac.RAMByteRequest = value.Float64
			}
		case allocationcost.FieldPvCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pv_cost", values[i])
			} else if value.Valid {
				ac.PvCost = value.Float64
			}
		case allocationcost.FieldPvBytes:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pv_bytes", values[i])
			} else if value.Valid {
				ac.PvBytes = value.Float64
			}
		case allocationcost.FieldLoadBalancerCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field load_balancer_cost", values[i])
			} else if value.Valid {
				ac.LoadBalancerCost = value.Float64
			}
		case allocationcost.FieldCPUCoreUsageAverage:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu_core_usage_average", values[i])
			} else if value.Valid {
				ac.CPUCoreUsageAverage = value.Float64
			}
		case allocationcost.FieldCPUCoreUsageMax:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu_core_usage_max", values[i])
			} else if value.Valid {
				ac.CPUCoreUsageMax = value.Float64
			}
		case allocationcost.FieldRAMByteUsageAverage:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ram_byte_usage_average", values[i])
			} else if value.Valid {
				ac.RAMByteUsageAverage = value.Float64
			}
		case allocationcost.FieldRAMByteUsageMax:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ram_byte_usage_max", values[i])
			} else if value.Valid {
				ac.RAMByteUsageMax = value.Float64
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AllocationCost.
// This includes values selected through modifiers, order, etc.
func (ac *AllocationCost) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// QueryConnector queries the "connector" edge of the AllocationCost entity.
func (ac *AllocationCost) QueryConnector() *ConnectorQuery {
	return NewAllocationCostClient(ac.config).QueryConnector(ac)
}

// Update returns a builder for updating this AllocationCost.
// Note that you need to call AllocationCost.Unwrap() before calling this method if this AllocationCost
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AllocationCost) Update() *AllocationCostUpdateOne {
	return NewAllocationCostClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AllocationCost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AllocationCost) Unwrap() *AllocationCost {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("model: AllocationCost is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AllocationCost) String() string {
	var builder strings.Builder
	builder.WriteString("AllocationCost(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("start_time=")
	builder.WriteString(ac.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(ac.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("minutes=")
	builder.WriteString(fmt.Sprintf("%v", ac.Minutes))
	builder.WriteString(", ")
	builder.WriteString("connector_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.ConnectorID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ac.Name)
	builder.WriteString(", ")
	builder.WriteString("fingerprint=")
	builder.WriteString(ac.Fingerprint)
	builder.WriteString(", ")
	builder.WriteString("cluster_name=")
	builder.WriteString(ac.ClusterName)
	builder.WriteString(", ")
	builder.WriteString("namespace=")
	builder.WriteString(ac.Namespace)
	builder.WriteString(", ")
	builder.WriteString("node=")
	builder.WriteString(ac.Node)
	builder.WriteString(", ")
	builder.WriteString("controller=")
	builder.WriteString(ac.Controller)
	builder.WriteString(", ")
	builder.WriteString("controller_kind=")
	builder.WriteString(ac.ControllerKind)
	builder.WriteString(", ")
	builder.WriteString("pod=")
	builder.WriteString(ac.Pod)
	builder.WriteString(", ")
	builder.WriteString("container=")
	builder.WriteString(ac.Container)
	builder.WriteString(", ")
	builder.WriteString("pvs=")
	builder.WriteString(fmt.Sprintf("%v", ac.Pvs))
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", ac.Labels))
	builder.WriteString(", ")
	builder.WriteString("totalCost=")
	builder.WriteString(fmt.Sprintf("%v", ac.TotalCost))
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(fmt.Sprintf("%v", ac.Currency))
	builder.WriteString(", ")
	builder.WriteString("cpu_cost=")
	builder.WriteString(fmt.Sprintf("%v", ac.CPUCost))
	builder.WriteString(", ")
	builder.WriteString("cpu_core_request=")
	builder.WriteString(fmt.Sprintf("%v", ac.CPUCoreRequest))
	builder.WriteString(", ")
	builder.WriteString("gpu_cost=")
	builder.WriteString(fmt.Sprintf("%v", ac.GpuCost))
	builder.WriteString(", ")
	builder.WriteString("gpu_count=")
	builder.WriteString(fmt.Sprintf("%v", ac.GpuCount))
	builder.WriteString(", ")
	builder.WriteString("ram_cost=")
	builder.WriteString(fmt.Sprintf("%v", ac.RAMCost))
	builder.WriteString(", ")
	builder.WriteString("ram_byte_request=")
	builder.WriteString(fmt.Sprintf("%v", ac.RAMByteRequest))
	builder.WriteString(", ")
	builder.WriteString("pv_cost=")
	builder.WriteString(fmt.Sprintf("%v", ac.PvCost))
	builder.WriteString(", ")
	builder.WriteString("pv_bytes=")
	builder.WriteString(fmt.Sprintf("%v", ac.PvBytes))
	builder.WriteString(", ")
	builder.WriteString("load_balancer_cost=")
	builder.WriteString(fmt.Sprintf("%v", ac.LoadBalancerCost))
	builder.WriteString(", ")
	builder.WriteString("cpu_core_usage_average=")
	builder.WriteString(fmt.Sprintf("%v", ac.CPUCoreUsageAverage))
	builder.WriteString(", ")
	builder.WriteString("cpu_core_usage_max=")
	builder.WriteString(fmt.Sprintf("%v", ac.CPUCoreUsageMax))
	builder.WriteString(", ")
	builder.WriteString("ram_byte_usage_average=")
	builder.WriteString(fmt.Sprintf("%v", ac.RAMByteUsageAverage))
	builder.WriteString(", ")
	builder.WriteString("ram_byte_usage_max=")
	builder.WriteString(fmt.Sprintf("%v", ac.RAMByteUsageMax))
	builder.WriteByte(')')
	return builder.String()
}

// AllocationCosts is a parsable slice of AllocationCost.
type AllocationCosts []*AllocationCost
