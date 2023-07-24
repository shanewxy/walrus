// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"time"

	"github.com/seal-io/seal/pkg/dao/model/setting"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// SettingCreateInput holds the creation input of the Setting entity.
type SettingCreateInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Value crypto.String `uri:"-" query:"-" json:"value"`
}

// Model returns the Setting entity for creating,
// after validating.
func (sci *SettingCreateInput) Model() *Setting {
	if sci == nil {
		return nil
	}

	s := &Setting{
		Value: sci.Value,
	}

	return s
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (sci *SettingCreateInput) Load() error {
	if sci == nil {
		return errors.New("nil receiver")
	}

	return sci.LoadWith(sci.inputConfig.Context, sci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (sci *SettingCreateInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if sci == nil {
		return errors.New("nil receiver")
	}

	return nil
}

// SettingCreateInputs holds the creation input item of the Setting entities.
type SettingCreateInputsItem struct {
	Value crypto.String `uri:"-" query:"-" json:"value"`
}

// SettingCreateInputs holds the creation input of the Setting entities.
type SettingCreateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*SettingCreateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the Setting entities for creating,
// after validating.
func (sci *SettingCreateInputs) Model() []*Setting {
	if sci == nil || len(sci.Items) == 0 {
		return nil
	}

	ss := make([]*Setting, len(sci.Items))

	for i := range sci.Items {
		s := &Setting{
			Value: sci.Items[i].Value,
		}

		ss[i] = s
	}

	return ss
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (sci *SettingCreateInputs) Load() error {
	if sci == nil {
		return errors.New("nil receiver")
	}

	return sci.LoadWith(sci.inputConfig.Context, sci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (sci *SettingCreateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if sci == nil {
		return errors.New("nil receiver")
	}

	if len(sci.Items) == 0 {
		return errors.New("empty items")
	}

	return nil
}

// SettingDeleteInput holds the deletion input of the Setting entity.
type SettingDeleteInput = SettingQueryInput

// SettingDeleteInputs holds the deletion input item of the Setting entities.
type SettingDeleteInputsItem struct {
	ID object.ID `uri:"-" query:"-" json:"id"`
}

// SettingDeleteInputs holds the deletion input of the Setting entities.
type SettingDeleteInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*SettingDeleteInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the Setting entities for deleting,
// after validating.
func (sdi *SettingDeleteInputs) Model() []*Setting {
	if sdi == nil || len(sdi.Items) == 0 {
		return nil
	}

	ss := make([]*Setting, len(sdi.Items))
	for i := range sdi.Items {
		ss[i] = &Setting{
			ID: sdi.Items[i].ID,
		}
	}
	return ss
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (sdi *SettingDeleteInputs) Load() error {
	if sdi == nil {
		return errors.New("nil receiver")
	}

	return sdi.LoadWith(sdi.inputConfig.Context, sdi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (sdi *SettingDeleteInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if sdi == nil {
		return errors.New("nil receiver")
	}

	if len(sdi.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Settings().Query()

	ids := make([]object.ID, 0, len(sdi.Items))

	for i := range sdi.Items {
		if sdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if sdi.Items[i].ID != "" {
			ids = append(ids, sdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(setting.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// SettingQueryInput holds the query input of the Setting entity.
type SettingQueryInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Refer *object.Refer `uri:"setting,default=\"\"" query:"-" json:"-"`
	ID    object.ID     `uri:"id" query:"-" json:"id"` // TODO(thxCode): remove the uri:"id" after supporting hierarchical routes.
}

// Model returns the Setting entity for querying,
// after validating.
func (sqi *SettingQueryInput) Model() *Setting {
	if sqi == nil {
		return nil
	}

	return &Setting{
		ID: sqi.ID,
	}
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (sqi *SettingQueryInput) Load() error {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	return sqi.LoadWith(sqi.inputConfig.Context, sqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (sqi *SettingQueryInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	if sqi.Refer != nil && *sqi.Refer == "" {
		return nil
	}

	q := cs.Settings().Query()

	if sqi.Refer != nil {
		if sqi.Refer.IsID() {
			q.Where(
				setting.ID(sqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of setting")
		}
	} else if sqi.ID != "" {
		q.Where(
			setting.ID(sqi.ID))
	} else {
		return errors.New("invalid identify of setting")
	}

	sqi.ID, err = q.OnlyID(ctx)
	return err
}

// SettingQueryInputs holds the query input of the Setting entities.
type SettingQueryInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (sqi *SettingQueryInputs) Load() error {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	return sqi.LoadWith(sqi.inputConfig.Context, sqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (sqi *SettingQueryInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	return err
}

// SettingUpdateInput holds the modification input of the Setting entity.
type SettingUpdateInput struct {
	SettingQueryInput `uri:",inline" query:"-" json:",inline"`

	Name  string        `uri:"-" query:"-" json:"name,omitempty"`
	Value crypto.String `uri:"-" query:"-" json:"value,omitempty"`
}

// Model returns the Setting entity for modifying,
// after validating.
func (sui *SettingUpdateInput) Model() *Setting {
	if sui == nil {
		return nil
	}

	s := &Setting{
		ID:    sui.ID,
		Name:  sui.Name,
		Value: sui.Value,
	}

	return s
}

// SettingUpdateInputs holds the modification input item of the Setting entities.
type SettingUpdateInputsItem struct {
	ID object.ID `uri:"-" query:"-" json:"id"`

	Name  string        `uri:"-" query:"-" json:"name,omitempty"`
	Value crypto.String `uri:"-" query:"-" json:"value,omitempty"`
}

// SettingUpdateInputs holds the modification input of the Setting entities.
type SettingUpdateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*SettingUpdateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the Setting entities for modifying,
// after validating.
func (sui *SettingUpdateInputs) Model() []*Setting {
	if sui == nil || len(sui.Items) == 0 {
		return nil
	}

	ss := make([]*Setting, len(sui.Items))

	for i := range sui.Items {
		s := &Setting{
			ID:    sui.Items[i].ID,
			Name:  sui.Items[i].Name,
			Value: sui.Items[i].Value,
		}

		ss[i] = s
	}

	return ss
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (sui *SettingUpdateInputs) Load() error {
	if sui == nil {
		return errors.New("nil receiver")
	}

	return sui.LoadWith(sui.inputConfig.Context, sui.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (sui *SettingUpdateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if sui == nil {
		return errors.New("nil receiver")
	}

	if len(sui.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.Settings().Query()

	ids := make([]object.ID, 0, len(sui.Items))

	for i := range sui.Items {
		if sui.Items[i] == nil {
			return errors.New("nil item")
		}

		if sui.Items[i].ID != "" {
			ids = append(ids, sui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(setting.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// SettingOutput holds the output of the Setting entity.
type SettingOutput struct {
	ID         object.ID     `json:"id,omitempty"`
	CreateTime *time.Time    `json:"createTime,omitempty"`
	UpdateTime *time.Time    `json:"updateTime,omitempty"`
	Name       string        `json:"name,omitempty"`
	Value      crypto.String `json:"value,omitempty"`
	Hidden     *bool         `json:"hidden,omitempty"`
	Editable   *bool         `json:"editable,omitempty"`
	Sensitive  *bool         `json:"sensitive,omitempty"`
}

// View returns the output of Setting.
func (s *Setting) View() *SettingOutput {
	return ExposeSetting(s)
}

// View returns the output of Settings.
func (ss Settings) View() []*SettingOutput {
	return ExposeSettings(ss)
}

// ExposeSetting converts the Setting to SettingOutput.
func ExposeSetting(s *Setting) *SettingOutput {
	if s == nil {
		return nil
	}

	so := &SettingOutput{
		ID:         s.ID,
		CreateTime: s.CreateTime,
		UpdateTime: s.UpdateTime,
		Name:       s.Name,
		Value:      s.Value,
		Hidden:     s.Hidden,
		Editable:   s.Editable,
		Sensitive:  s.Sensitive,
	}

	return so
}

// ExposeSettings converts the Setting slice to SettingOutput pointer slice.
func ExposeSettings(ss []*Setting) []*SettingOutput {
	if len(ss) == 0 {
		return nil
	}

	sos := make([]*SettingOutput, len(ss))
	for i := range ss {
		sos[i] = ExposeSetting(ss[i])
	}
	return sos
}
