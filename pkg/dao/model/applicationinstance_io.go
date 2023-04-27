// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types/oid"
	"github.com/seal-io/seal/pkg/dao/types/property"
	"github.com/seal-io/seal/pkg/dao/types/status"
)

// ApplicationInstanceQueryInput is the input for the ApplicationInstance query.
type ApplicationInstanceQueryInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the ApplicationInstanceQueryInput to ApplicationInstance.
func (in ApplicationInstanceQueryInput) Model() *ApplicationInstance {
	return &ApplicationInstance{
		ID: in.ID,
	}
}

// ApplicationInstanceCreateInput is the input for the ApplicationInstance creation.
type ApplicationInstanceCreateInput struct {
	// Name of the instance.
	Name string `json:"name"`
	// Variables of the instance.
	Variables property.Values `json:"variables,omitempty"`
	// Status of the instance.
	Status status.Status `json:"status,omitempty"`
	// Application to which the instance belongs.
	Application ApplicationQueryInput `json:"application"`
	// Environment to which the instance belongs.
	Environment EnvironmentQueryInput `json:"environment"`
}

// Model converts the ApplicationInstanceCreateInput to ApplicationInstance.
func (in ApplicationInstanceCreateInput) Model() *ApplicationInstance {
	var entity = &ApplicationInstance{
		Name:      in.Name,
		Variables: in.Variables,
		Status:    in.Status,
	}
	entity.ApplicationID = in.Application.ID
	entity.EnvironmentID = in.Environment.ID
	return entity
}

// ApplicationInstanceUpdateInput is the input for the ApplicationInstance modification.
type ApplicationInstanceUpdateInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id" json:"-"`
	// Variables of the instance.
	Variables property.Values `json:"variables,omitempty"`
	// Status of the instance.
	Status status.Status `json:"status,omitempty"`
}

// Model converts the ApplicationInstanceUpdateInput to ApplicationInstance.
func (in ApplicationInstanceUpdateInput) Model() *ApplicationInstance {
	var entity = &ApplicationInstance{
		ID:        in.ID,
		Variables: in.Variables,
		Status:    in.Status,
	}
	return entity
}

// ApplicationInstanceOutput is the output for the ApplicationInstance.
type ApplicationInstanceOutput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `json:"id,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Name of the instance.
	Name string `json:"name,omitempty"`
	// Variables of the instance.
	Variables property.Values `json:"variables,omitempty"`
	// Status of the instance.
	Status status.Status `json:"status,omitempty"`
	// Application to which the instance belongs.
	Application *ApplicationOutput `json:"application,omitempty"`
	// Environment to which the instance belongs.
	Environment *EnvironmentOutput `json:"environment,omitempty"`
	// Application revisions that belong to this instance.
	Revisions []*ApplicationRevisionOutput `json:"revisions,omitempty"`
	// Application resources that belong to the instance.
	Resources []*ApplicationResourceOutput `json:"resources,omitempty"`
}

// ExposeApplicationInstance converts the ApplicationInstance to ApplicationInstanceOutput.
func ExposeApplicationInstance(in *ApplicationInstance) *ApplicationInstanceOutput {
	if in == nil {
		return nil
	}
	var entity = &ApplicationInstanceOutput{
		ID:          in.ID,
		CreateTime:  in.CreateTime,
		UpdateTime:  in.UpdateTime,
		Name:        in.Name,
		Variables:   in.Variables,
		Status:      in.Status,
		Application: ExposeApplication(in.Edges.Application),
		Environment: ExposeEnvironment(in.Edges.Environment),
		Revisions:   ExposeApplicationRevisions(in.Edges.Revisions),
		Resources:   ExposeApplicationResources(in.Edges.Resources),
	}
	if entity.Application == nil {
		entity.Application = &ApplicationOutput{}
	}
	entity.Application.ID = in.ApplicationID
	if entity.Environment == nil {
		entity.Environment = &EnvironmentOutput{}
	}
	entity.Environment.ID = in.EnvironmentID
	return entity
}

// ExposeApplicationInstances converts the ApplicationInstance slice to ApplicationInstanceOutput pointer slice.
func ExposeApplicationInstances(in []*ApplicationInstance) []*ApplicationInstanceOutput {
	var out = make([]*ApplicationInstanceOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeApplicationInstance(in[i])
		if o == nil {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
