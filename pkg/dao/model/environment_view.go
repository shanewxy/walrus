// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/schema/intercept"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// EnvironmentCreateInput holds the creation input of the Environment entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Environment entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Connectors specifies full inserting the new EnvironmentConnectorRelationship entities of the Environment entity.
	Connectors []*EnvironmentConnectorRelationshipCreateInput `uri:"-" query:"-" json:"connectors,omitempty"`
	// Services specifies full inserting the new Service entities of the Environment entity.
	Services []*ServiceCreateInput `uri:"-" query:"-" json:"services,cli-ignore,omitempty"`
}

// Model returns the Environment entity for creating,
// after validating.
func (eci *EnvironmentCreateInput) Model() *Environment {
	if eci == nil {
		return nil
	}

	_e := &Environment{
		Name:        eci.Name,
		Description: eci.Description,
		Labels:      eci.Labels,
	}

	if eci.Project != nil {
		_e.ProjectID = eci.Project.ID
	}

	for j := range eci.Connectors {
		if eci.Connectors[j] == nil {
			continue
		}
		_e.Edges.Connectors = append(_e.Edges.Connectors,
			eci.Connectors[j].Model())
	}
	for j := range eci.Services {
		if eci.Services[j] == nil {
			continue
		}
		_e.Edges.Services = append(_e.Edges.Services,
			eci.Services[j].Model())
	}
	return _e
}

// Validate checks the EnvironmentCreateInput entity.
func (eci *EnvironmentCreateInput) Validate() error {
	if eci == nil {
		return errors.New("nil receiver")
	}

	return eci.ValidateWith(eci.inputConfig.Context, eci.inputConfig.Client)
}

// ValidateWith checks the EnvironmentCreateInput entity with the given context and client set.
func (eci *EnvironmentCreateInput) ValidateWith(ctx context.Context, cs ClientSet) error {
	if eci == nil {
		return errors.New("nil receiver")
	}

	// Validate when creating under the Project route.
	if eci.Project != nil {
		if err := eci.Project.ValidateWith(ctx, cs); err != nil {
			return err
		}
	}

	for i := range eci.Connectors {
		if eci.Connectors[i] == nil {
			continue
		}

		if err := eci.Connectors[i].ValidateWith(ctx, cs); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				eci.Connectors[i] = nil
			}
		}
	}

	for i := range eci.Services {
		if eci.Services[i] == nil {
			continue
		}

		if err := eci.Services[i].ValidateWith(ctx, cs); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				eci.Services[i] = nil
			}
		}
	}

	return nil
}

// EnvironmentCreateInputs holds the creation input item of the Environment entities.
type EnvironmentCreateInputsItem struct {
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Connectors specifies full inserting the new EnvironmentConnectorRelationship entities.
	Connectors []*EnvironmentConnectorRelationshipCreateInput `uri:"-" query:"-" json:"connectors,omitempty"`
	// Services specifies full inserting the new Service entities.
	Services []*ServiceCreateInput `uri:"-" query:"-" json:"services,cli-ignore,omitempty"`
}

// ValidateWith checks the EnvironmentCreateInputsItem entity with the given context and client set.
func (eci *EnvironmentCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet) error {
	if eci == nil {
		return errors.New("nil receiver")
	}

	for i := range eci.Connectors {
		if eci.Connectors[i] == nil {
			continue
		}

		if err := eci.Connectors[i].ValidateWith(ctx, cs); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				eci.Connectors[i] = nil
			}
		}
	}

	for i := range eci.Services {
		if eci.Services[i] == nil {
			continue
		}

		if err := eci.Services[i].ValidateWith(ctx, cs); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				eci.Services[i] = nil
			}
		}
	}

	return nil
}

// EnvironmentCreateInputs holds the creation input of the Environment entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Environment entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*EnvironmentCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Environment entities for creating,
// after validating.
func (eci *EnvironmentCreateInputs) Model() []*Environment {
	if eci == nil || len(eci.Items) == 0 {
		return nil
	}

	_es := make([]*Environment, len(eci.Items))

	for i := range eci.Items {
		_e := &Environment{
			Name:        eci.Items[i].Name,
			Description: eci.Items[i].Description,
			Labels:      eci.Items[i].Labels,
		}

		if eci.Project != nil {
			_e.ProjectID = eci.Project.ID
		}

		for j := range eci.Items[i].Connectors {
			if eci.Items[i].Connectors[j] == nil {
				continue
			}
			_e.Edges.Connectors = append(_e.Edges.Connectors,
				eci.Items[i].Connectors[j].Model())
		}
		for j := range eci.Items[i].Services {
			if eci.Items[i].Services[j] == nil {
				continue
			}
			_e.Edges.Services = append(_e.Edges.Services,
				eci.Items[i].Services[j].Model())
		}

		_es[i] = _e
	}

	return _es
}

// Validate checks the EnvironmentCreateInputs entity .
func (eci *EnvironmentCreateInputs) Validate() error {
	if eci == nil {
		return errors.New("nil receiver")
	}

	return eci.ValidateWith(eci.inputConfig.Context, eci.inputConfig.Client)
}

// ValidateWith checks the EnvironmentCreateInputs entity with the given context and client set.
func (eci *EnvironmentCreateInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if eci == nil {
		return errors.New("nil receiver")
	}

	if len(eci.Items) == 0 {
		return errors.New("empty items")
	}

	// Validate when creating under the Project route.
	if eci.Project != nil {
		if err := eci.Project.ValidateWith(ctx, cs); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				eci.Project = nil
			}
		}
	}

	for i := range eci.Items {
		if eci.Items[i] == nil {
			continue
		}

		if err := eci.Items[i].ValidateWith(ctx, cs); err != nil {
			return err
		}
	}

	return nil
}

// EnvironmentDeleteInput holds the deletion input of the Environment entity,
// please tags with `path:",inline"` if embedding.
type EnvironmentDeleteInput struct {
	EnvironmentQueryInput `path:",inline"`
}

// EnvironmentDeleteInputs holds the deletion input item of the Environment entities.
type EnvironmentDeleteInputsItem struct {
	// ID of the Environment entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Environment entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// EnvironmentDeleteInputs holds the deletion input of the Environment entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete Environment entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*EnvironmentDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Environment entities for deleting,
// after validating.
func (edi *EnvironmentDeleteInputs) Model() []*Environment {
	if edi == nil || len(edi.Items) == 0 {
		return nil
	}

	_es := make([]*Environment, len(edi.Items))
	for i := range edi.Items {
		_es[i] = &Environment{
			ID: edi.Items[i].ID,
		}
	}
	return _es
}

// IDs returns the ID list of the Environment entities for deleting,
// after validating.
func (edi *EnvironmentDeleteInputs) IDs() []object.ID {
	if edi == nil || len(edi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(edi.Items))
	for i := range edi.Items {
		ids[i] = edi.Items[i].ID
	}
	return ids
}

// Validate checks the EnvironmentDeleteInputs entity.
func (edi *EnvironmentDeleteInputs) Validate() error {
	if edi == nil {
		return errors.New("nil receiver")
	}

	return edi.ValidateWith(edi.inputConfig.Context, edi.inputConfig.Client)
}

// ValidateWith checks the EnvironmentDeleteInputs entity with the given context and client set.
func (edi *EnvironmentDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if edi == nil {
		return errors.New("nil receiver")
	}

	if len(edi.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Environments().Query()

	// Validate when deleting under the Project route.
	if edi.Project != nil {
		if err := edi.Project.ValidateWith(ctx, cs); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				environment.ProjectID(edi.Project.ID))
		}
	}

	ids := make([]object.ID, 0, len(edi.Items))
	ors := make([]predicate.Environment, 0, len(edi.Items))

	for i := range edi.Items {
		if edi.Items[i] == nil {
			return errors.New("nil item")
		}

		if edi.Items[i].ID != "" {
			ids = append(ids, edi.Items[i].ID)
			ors = append(ors, environment.ID(edi.Items[i].ID))
		} else if edi.Items[i].Name != "" {
			ors = append(ors, environment.And(
				environment.Name(edi.Items[i].Name)))
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := environment.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = environment.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			environment.FieldID,
			environment.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		edi.Items[i].ID = es[i].ID
		edi.Items[i].Name = es[i].Name
	}

	return nil
}

// EnvironmentQueryInput holds the query input of the Environment entity,
// please tags with `path:",inline"` if embedding.
type EnvironmentQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Environment entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project"`

	// Refer holds the route path reference of the Environment entity.
	Refer *object.Refer `path:"environment,default=" query:"-" json:"-"`
	// ID of the Environment entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Environment entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Environment entity for querying,
// after validating.
func (eqi *EnvironmentQueryInput) Model() *Environment {
	if eqi == nil {
		return nil
	}

	return &Environment{
		ID:   eqi.ID,
		Name: eqi.Name,
	}
}

// Validate checks the EnvironmentQueryInput entity.
func (eqi *EnvironmentQueryInput) Validate() error {
	if eqi == nil {
		return errors.New("nil receiver")
	}

	return eqi.ValidateWith(eqi.inputConfig.Context, eqi.inputConfig.Client)
}

// ValidateWith checks the EnvironmentQueryInput entity with the given context and client set.
func (eqi *EnvironmentQueryInput) ValidateWith(ctx context.Context, cs ClientSet) error {
	if eqi == nil {
		return errors.New("nil receiver")
	}

	if eqi.Refer != nil && *eqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", environment.Label, ErrBlankResourceRefer)
	}

	q := cs.Environments().Query()

	// Validate when querying under the Project route.
	if eqi.Project != nil {
		if err := eqi.Project.ValidateWith(ctx, cs); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				environment.ProjectID(eqi.Project.ID))
		}
	}

	if eqi.Refer != nil {
		if eqi.Refer.IsID() {
			q.Where(
				environment.ID(eqi.Refer.ID()))
		} else if refers := eqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				environment.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of environment")
		}
	} else if eqi.ID != "" {
		q.Where(
			environment.ID(eqi.ID))
	} else if eqi.Name != "" {
		q.Where(
			environment.Name(eqi.Name))
	} else {
		return errors.New("invalid identify of environment")
	}

	e, err := q.
		Select(
			environment.FieldID,
			environment.FieldName,
		).
		Only(ctx)
	if err == nil {
		eqi.ID = e.ID
		eqi.Name = e.Name
	}
	return err
}

// EnvironmentQueryInputs holds the query input of the Environment entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type EnvironmentQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Environment entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the EnvironmentQueryInputs entity.
func (eqi *EnvironmentQueryInputs) Validate() error {
	if eqi == nil {
		return errors.New("nil receiver")
	}

	return eqi.ValidateWith(eqi.inputConfig.Context, eqi.inputConfig.Client)
}

// ValidateWith checks the EnvironmentQueryInputs entity with the given context and client set.
func (eqi *EnvironmentQueryInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if eqi == nil {
		return errors.New("nil receiver")
	}

	// Validate when querying under the Project route.
	if eqi.Project != nil {
		if err := eqi.Project.ValidateWith(ctx, cs); err != nil {
			return err
		}
	}

	return nil
}

// EnvironmentUpdateInput holds the modification input of the Environment entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentUpdateInput struct {
	EnvironmentQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Connectors indicates replacing the stale EnvironmentConnectorRelationship entities.
	Connectors []*EnvironmentConnectorRelationshipCreateInput `uri:"-" query:"-" json:"connectors,omitempty"`
}

// Model returns the Environment entity for modifying,
// after validating.
func (eui *EnvironmentUpdateInput) Model() *Environment {
	if eui == nil {
		return nil
	}

	_e := &Environment{
		ID:          eui.ID,
		Description: eui.Description,
		Labels:      eui.Labels,
	}

	for j := range eui.Connectors {
		if eui.Connectors[j] == nil {
			continue
		}
		_e.Edges.Connectors = append(_e.Edges.Connectors,
			eui.Connectors[j].Model())
	}
	return _e
}

// Validate checks the EnvironmentUpdateInput entity.
func (eui *EnvironmentUpdateInput) Validate() error {
	if eui == nil {
		return errors.New("nil receiver")
	}

	return eui.ValidateWith(eui.inputConfig.Context, eui.inputConfig.Client)
}

// ValidateWith checks the EnvironmentUpdateInput entity with the given context and client set.
func (eui *EnvironmentUpdateInput) ValidateWith(ctx context.Context, cs ClientSet) error {
	if err := eui.EnvironmentQueryInput.ValidateWith(ctx, cs); err != nil {
		return err
	}

	for i := range eui.Connectors {
		if eui.Connectors[i] == nil {
			continue
		}

		if err := eui.Connectors[i].ValidateWith(ctx, cs); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				eui.Connectors[i] = nil
			}
		}
	}

	return nil
}

// EnvironmentUpdateInputs holds the modification input item of the Environment entities.
type EnvironmentUpdateInputsItem struct {
	// ID of the Environment entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Environment entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Connectors indicates replacing the stale EnvironmentConnectorRelationship entities.
	Connectors []*EnvironmentConnectorRelationshipCreateInput `uri:"-" query:"-" json:"connectors,omitempty"`
}

// ValidateWith checks the EnvironmentUpdateInputsItem entity with the given context and client set.
func (eui *EnvironmentUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet) error {
	if eui == nil {
		return errors.New("nil receiver")
	}

	for i := range eui.Connectors {
		if eui.Connectors[i] == nil {
			continue
		}

		if err := eui.Connectors[i].ValidateWith(ctx, cs); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				eui.Connectors[i] = nil
			}
		}
	}

	return nil
}

// EnvironmentUpdateInputs holds the modification input of the Environment entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update Environment entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*EnvironmentUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Environment entities for modifying,
// after validating.
func (eui *EnvironmentUpdateInputs) Model() []*Environment {
	if eui == nil || len(eui.Items) == 0 {
		return nil
	}

	_es := make([]*Environment, len(eui.Items))

	for i := range eui.Items {
		_e := &Environment{
			ID:          eui.Items[i].ID,
			Description: eui.Items[i].Description,
			Labels:      eui.Items[i].Labels,
		}

		for j := range eui.Items[i].Connectors {
			if eui.Items[i].Connectors[j] == nil {
				continue
			}
			_e.Edges.Connectors = append(_e.Edges.Connectors,
				eui.Items[i].Connectors[j].Model())
		}

		_es[i] = _e
	}

	return _es
}

// IDs returns the ID list of the Environment entities for modifying,
// after validating.
func (eui *EnvironmentUpdateInputs) IDs() []object.ID {
	if eui == nil || len(eui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(eui.Items))
	for i := range eui.Items {
		ids[i] = eui.Items[i].ID
	}
	return ids
}

// Validate checks the EnvironmentUpdateInputs entity.
func (eui *EnvironmentUpdateInputs) Validate() error {
	if eui == nil {
		return errors.New("nil receiver")
	}

	return eui.ValidateWith(eui.inputConfig.Context, eui.inputConfig.Client)
}

// ValidateWith checks the EnvironmentUpdateInputs entity with the given context and client set.
func (eui *EnvironmentUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet) error {
	if eui == nil {
		return errors.New("nil receiver")
	}

	if len(eui.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Environments().Query()

	// Validate when updating under the Project route.
	if eui.Project != nil {
		if err := eui.Project.ValidateWith(ctx, cs); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				environment.ProjectID(eui.Project.ID))
		}
	}

	ids := make([]object.ID, 0, len(eui.Items))
	ors := make([]predicate.Environment, 0, len(eui.Items))

	for i := range eui.Items {
		if eui.Items[i] == nil {
			return errors.New("nil item")
		}

		if eui.Items[i].ID != "" {
			ids = append(ids, eui.Items[i].ID)
			ors = append(ors, environment.ID(eui.Items[i].ID))
		} else if eui.Items[i].Name != "" {
			ors = append(ors, environment.And(
				environment.Name(eui.Items[i].Name)))
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := environment.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = environment.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			environment.FieldID,
			environment.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		eui.Items[i].ID = es[i].ID
		eui.Items[i].Name = es[i].Name
	}

	for i := range eui.Items {
		if eui.Items[i] == nil {
			continue
		}

		if err := eui.Items[i].ValidateWith(ctx, cs); err != nil {
			return err
		}
	}

	return nil
}

// EnvironmentOutput holds the output of the Environment entity.
type EnvironmentOutput struct {
	ID          object.ID         `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	CreateTime  *time.Time        `json:"createTime,omitempty"`
	UpdateTime  *time.Time        `json:"updateTime,omitempty"`

	Project    *ProjectOutput                            `json:"project,omitempty"`
	Connectors []*EnvironmentConnectorRelationshipOutput `json:"connectors,omitempty"`
}

// View returns the output of Environment entity.
func (_e *Environment) View() *EnvironmentOutput {
	return ExposeEnvironment(_e)
}

// View returns the output of Environment entities.
func (_es Environments) View() []*EnvironmentOutput {
	return ExposeEnvironments(_es)
}

// ExposeEnvironment converts the Environment to EnvironmentOutput.
func ExposeEnvironment(_e *Environment) *EnvironmentOutput {
	if _e == nil {
		return nil
	}

	eo := &EnvironmentOutput{
		ID:          _e.ID,
		Name:        _e.Name,
		Description: _e.Description,
		Labels:      _e.Labels,
		CreateTime:  _e.CreateTime,
		UpdateTime:  _e.UpdateTime,
	}

	if _e.Edges.Project != nil {
		eo.Project = ExposeProject(_e.Edges.Project)
	} else if _e.ProjectID != "" {
		eo.Project = &ProjectOutput{
			ID: _e.ProjectID,
		}
	}
	if _e.Edges.Connectors != nil {
		eo.Connectors = ExposeEnvironmentConnectorRelationships(_e.Edges.Connectors)
	}
	return eo
}

// ExposeEnvironments converts the Environment slice to EnvironmentOutput pointer slice.
func ExposeEnvironments(_es []*Environment) []*EnvironmentOutput {
	if len(_es) == 0 {
		return nil
	}

	eos := make([]*EnvironmentOutput, len(_es))
	for i := range _es {
		eos[i] = ExposeEnvironment(_es[i])
	}
	return eos
}
