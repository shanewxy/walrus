// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package service

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types/object"
	"github.com/seal-io/seal/pkg/dao/types/property"
)

// ID filters vertices based on their ID field.
func ID(id object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id object.ID) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id object.ID) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id object.ID) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id object.ID) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id object.ID) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldDescription, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldUpdateTime, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldProjectID, v))
}

// EnvironmentID applies equality check predicate on the "environment_id" field. It's identical to EnvironmentIDEQ.
func EnvironmentID(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldEnvironmentID, v))
}

// TemplateID applies equality check predicate on the "template_id" field. It's identical to TemplateIDEQ.
func TemplateID(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldTemplateID, v))
}

// Attributes applies equality check predicate on the "attributes" field. It's identical to AttributesEQ.
func Attributes(v property.Values) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldAttributes, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Service {
	return predicate.Service(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Service {
	return predicate.Service(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Service {
	return predicate.Service(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Service {
	return predicate.Service(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Service {
	return predicate.Service(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Service {
	return predicate.Service(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Service {
	return predicate.Service(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Service {
	return predicate.Service(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Service {
	return predicate.Service(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Service {
	return predicate.Service(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Service {
	return predicate.Service(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Service {
	return predicate.Service(sql.FieldContainsFold(FieldDescription, v))
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.Service {
	return predicate.Service(sql.FieldIsNull(FieldLabels))
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.Service {
	return predicate.Service(sql.FieldNotNull(FieldLabels))
}

// AnnotationsIsNil applies the IsNil predicate on the "annotations" field.
func AnnotationsIsNil() predicate.Service {
	return predicate.Service(sql.FieldIsNull(FieldAnnotations))
}

// AnnotationsNotNil applies the NotNil predicate on the "annotations" field.
func AnnotationsNotNil() predicate.Service {
	return predicate.Service(sql.FieldNotNull(FieldAnnotations))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldUpdateTime, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Service {
	return predicate.Service(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Service {
	return predicate.Service(sql.FieldNotNull(FieldStatus))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldContains(FieldProjectID, vc))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldHasPrefix(FieldProjectID, vc))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldHasSuffix(FieldProjectID, vc))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldEqualFold(FieldProjectID, vc))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldContainsFold(FieldProjectID, vc))
}

// EnvironmentIDEQ applies the EQ predicate on the "environment_id" field.
func EnvironmentIDEQ(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldEnvironmentID, v))
}

// EnvironmentIDNEQ applies the NEQ predicate on the "environment_id" field.
func EnvironmentIDNEQ(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldEnvironmentID, v))
}

// EnvironmentIDIn applies the In predicate on the "environment_id" field.
func EnvironmentIDIn(vs ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldEnvironmentID, vs...))
}

// EnvironmentIDNotIn applies the NotIn predicate on the "environment_id" field.
func EnvironmentIDNotIn(vs ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldEnvironmentID, vs...))
}

// EnvironmentIDGT applies the GT predicate on the "environment_id" field.
func EnvironmentIDGT(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldEnvironmentID, v))
}

// EnvironmentIDGTE applies the GTE predicate on the "environment_id" field.
func EnvironmentIDGTE(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldEnvironmentID, v))
}

// EnvironmentIDLT applies the LT predicate on the "environment_id" field.
func EnvironmentIDLT(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldEnvironmentID, v))
}

// EnvironmentIDLTE applies the LTE predicate on the "environment_id" field.
func EnvironmentIDLTE(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldEnvironmentID, v))
}

// EnvironmentIDContains applies the Contains predicate on the "environment_id" field.
func EnvironmentIDContains(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldContains(FieldEnvironmentID, vc))
}

// EnvironmentIDHasPrefix applies the HasPrefix predicate on the "environment_id" field.
func EnvironmentIDHasPrefix(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldHasPrefix(FieldEnvironmentID, vc))
}

// EnvironmentIDHasSuffix applies the HasSuffix predicate on the "environment_id" field.
func EnvironmentIDHasSuffix(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldHasSuffix(FieldEnvironmentID, vc))
}

// EnvironmentIDEqualFold applies the EqualFold predicate on the "environment_id" field.
func EnvironmentIDEqualFold(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldEqualFold(FieldEnvironmentID, vc))
}

// EnvironmentIDContainsFold applies the ContainsFold predicate on the "environment_id" field.
func EnvironmentIDContainsFold(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldContainsFold(FieldEnvironmentID, vc))
}

// TemplateIDEQ applies the EQ predicate on the "template_id" field.
func TemplateIDEQ(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldTemplateID, v))
}

// TemplateIDNEQ applies the NEQ predicate on the "template_id" field.
func TemplateIDNEQ(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldTemplateID, v))
}

// TemplateIDIn applies the In predicate on the "template_id" field.
func TemplateIDIn(vs ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldTemplateID, vs...))
}

// TemplateIDNotIn applies the NotIn predicate on the "template_id" field.
func TemplateIDNotIn(vs ...object.ID) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldTemplateID, vs...))
}

// TemplateIDGT applies the GT predicate on the "template_id" field.
func TemplateIDGT(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldTemplateID, v))
}

// TemplateIDGTE applies the GTE predicate on the "template_id" field.
func TemplateIDGTE(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldTemplateID, v))
}

// TemplateIDLT applies the LT predicate on the "template_id" field.
func TemplateIDLT(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldTemplateID, v))
}

// TemplateIDLTE applies the LTE predicate on the "template_id" field.
func TemplateIDLTE(v object.ID) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldTemplateID, v))
}

// TemplateIDContains applies the Contains predicate on the "template_id" field.
func TemplateIDContains(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldContains(FieldTemplateID, vc))
}

// TemplateIDHasPrefix applies the HasPrefix predicate on the "template_id" field.
func TemplateIDHasPrefix(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldHasPrefix(FieldTemplateID, vc))
}

// TemplateIDHasSuffix applies the HasSuffix predicate on the "template_id" field.
func TemplateIDHasSuffix(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldHasSuffix(FieldTemplateID, vc))
}

// TemplateIDEqualFold applies the EqualFold predicate on the "template_id" field.
func TemplateIDEqualFold(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldEqualFold(FieldTemplateID, vc))
}

// TemplateIDContainsFold applies the ContainsFold predicate on the "template_id" field.
func TemplateIDContainsFold(v object.ID) predicate.Service {
	vc := string(v)
	return predicate.Service(sql.FieldContainsFold(FieldTemplateID, vc))
}

// AttributesEQ applies the EQ predicate on the "attributes" field.
func AttributesEQ(v property.Values) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldAttributes, v))
}

// AttributesNEQ applies the NEQ predicate on the "attributes" field.
func AttributesNEQ(v property.Values) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldAttributes, v))
}

// AttributesIn applies the In predicate on the "attributes" field.
func AttributesIn(vs ...property.Values) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldAttributes, vs...))
}

// AttributesNotIn applies the NotIn predicate on the "attributes" field.
func AttributesNotIn(vs ...property.Values) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldAttributes, vs...))
}

// AttributesGT applies the GT predicate on the "attributes" field.
func AttributesGT(v property.Values) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldAttributes, v))
}

// AttributesGTE applies the GTE predicate on the "attributes" field.
func AttributesGTE(v property.Values) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldAttributes, v))
}

// AttributesLT applies the LT predicate on the "attributes" field.
func AttributesLT(v property.Values) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldAttributes, v))
}

// AttributesLTE applies the LTE predicate on the "attributes" field.
func AttributesLTE(v property.Values) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldAttributes, v))
}

// AttributesIsNil applies the IsNil predicate on the "attributes" field.
func AttributesIsNil() predicate.Service {
	return predicate.Service(sql.FieldIsNull(FieldAttributes))
}

// AttributesNotNil applies the NotNil predicate on the "attributes" field.
func AttributesNotNil() predicate.Service {
	return predicate.Service(sql.FieldNotNull(FieldAttributes))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Service
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newProjectStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Service
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEnvironment applies the HasEdge predicate on the "environment" edge.
func HasEnvironment() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EnvironmentTable, EnvironmentColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Environment
		step.Edge.Schema = schemaConfig.Service
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEnvironmentWith applies the HasEdge predicate on the "environment" edge with a given conditions (other predicates).
func HasEnvironmentWith(preds ...predicate.Environment) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newEnvironmentStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Environment
		step.Edge.Schema = schemaConfig.Service
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTemplate applies the HasEdge predicate on the "template" edge.
func HasTemplate() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TemplateTable, TemplateColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TemplateVersion
		step.Edge.Schema = schemaConfig.Service
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTemplateWith applies the HasEdge predicate on the "template" edge with a given conditions (other predicates).
func HasTemplateWith(preds ...predicate.TemplateVersion) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newTemplateStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TemplateVersion
		step.Edge.Schema = schemaConfig.Service
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRevisions applies the HasEdge predicate on the "revisions" edge.
func HasRevisions() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RevisionsTable, RevisionsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ServiceRevision
		step.Edge.Schema = schemaConfig.ServiceRevision
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRevisionsWith applies the HasEdge predicate on the "revisions" edge with a given conditions (other predicates).
func HasRevisionsWith(preds ...predicate.ServiceRevision) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newRevisionsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ServiceRevision
		step.Edge.Schema = schemaConfig.ServiceRevision
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResources applies the HasEdge predicate on the "resources" edge.
func HasResources() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ServiceResource
		step.Edge.Schema = schemaConfig.ServiceResource
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourcesWith applies the HasEdge predicate on the "resources" edge with a given conditions (other predicates).
func HasResourcesWith(preds ...predicate.ServiceResource) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newResourcesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ServiceResource
		step.Edge.Schema = schemaConfig.ServiceResource
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDependencies applies the HasEdge predicate on the "dependencies" edge.
func HasDependencies() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DependenciesTable, DependenciesColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ServiceRelationship
		step.Edge.Schema = schemaConfig.ServiceRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDependenciesWith applies the HasEdge predicate on the "dependencies" edge with a given conditions (other predicates).
func HasDependenciesWith(preds ...predicate.ServiceRelationship) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newDependenciesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ServiceRelationship
		step.Edge.Schema = schemaConfig.ServiceRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Service) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Service) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Service) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		p(s.Not())
	})
}
