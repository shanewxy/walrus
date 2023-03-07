// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package perspective

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types"
)

// ID filters vertices based on their ID field.
func ID(id types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id types.ID) predicate.Perspective {
	return predicate.Perspective(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "createTime" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "updateTime" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldName, v))
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "endTime" field. It's identical to EndTimeEQ.
func EndTime(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldEndTime, v))
}

// Builtin applies equality check predicate on the "builtin" field. It's identical to BuiltinEQ.
func Builtin(v bool) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldBuiltin, v))
}

// CreateTimeEQ applies the EQ predicate on the "createTime" field.
func CreateTimeEQ(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "createTime" field.
func CreateTimeNEQ(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "createTime" field.
func CreateTimeIn(vs ...time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "createTime" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "createTime" field.
func CreateTimeGT(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "createTime" field.
func CreateTimeGTE(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "createTime" field.
func CreateTimeLT(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "createTime" field.
func CreateTimeLTE(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "updateTime" field.
func UpdateTimeEQ(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "updateTime" field.
func UpdateTimeNEQ(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "updateTime" field.
func UpdateTimeIn(vs ...time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "updateTime" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "updateTime" field.
func UpdateTimeGT(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "updateTime" field.
func UpdateTimeGTE(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "updateTime" field.
func UpdateTimeLT(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "updateTime" field.
func UpdateTimeLTE(v time.Time) predicate.Perspective {
	return predicate.Perspective(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Perspective {
	return predicate.Perspective(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Perspective {
	return predicate.Perspective(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldContainsFold(FieldName, v))
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...string) predicate.Perspective {
	return predicate.Perspective(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "startTime" field.
func StartTimeNotIn(vs ...string) predicate.Perspective {
	return predicate.Perspective(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "startTime" field.
func StartTimeGT(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeContains applies the Contains predicate on the "startTime" field.
func StartTimeContains(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldContains(FieldStartTime, v))
}

// StartTimeHasPrefix applies the HasPrefix predicate on the "startTime" field.
func StartTimeHasPrefix(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldHasPrefix(FieldStartTime, v))
}

// StartTimeHasSuffix applies the HasSuffix predicate on the "startTime" field.
func StartTimeHasSuffix(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldHasSuffix(FieldStartTime, v))
}

// StartTimeEqualFold applies the EqualFold predicate on the "startTime" field.
func StartTimeEqualFold(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEqualFold(FieldStartTime, v))
}

// StartTimeContainsFold applies the ContainsFold predicate on the "startTime" field.
func StartTimeContainsFold(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldContainsFold(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "endTime" field.
func EndTimeEQ(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "endTime" field.
func EndTimeNEQ(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "endTime" field.
func EndTimeIn(vs ...string) predicate.Perspective {
	return predicate.Perspective(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "endTime" field.
func EndTimeNotIn(vs ...string) predicate.Perspective {
	return predicate.Perspective(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "endTime" field.
func EndTimeGT(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "endTime" field.
func EndTimeGTE(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "endTime" field.
func EndTimeLT(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "endTime" field.
func EndTimeLTE(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldLTE(FieldEndTime, v))
}

// EndTimeContains applies the Contains predicate on the "endTime" field.
func EndTimeContains(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldContains(FieldEndTime, v))
}

// EndTimeHasPrefix applies the HasPrefix predicate on the "endTime" field.
func EndTimeHasPrefix(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldHasPrefix(FieldEndTime, v))
}

// EndTimeHasSuffix applies the HasSuffix predicate on the "endTime" field.
func EndTimeHasSuffix(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldHasSuffix(FieldEndTime, v))
}

// EndTimeEqualFold applies the EqualFold predicate on the "endTime" field.
func EndTimeEqualFold(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldEqualFold(FieldEndTime, v))
}

// EndTimeContainsFold applies the ContainsFold predicate on the "endTime" field.
func EndTimeContainsFold(v string) predicate.Perspective {
	return predicate.Perspective(sql.FieldContainsFold(FieldEndTime, v))
}

// BuiltinEQ applies the EQ predicate on the "builtin" field.
func BuiltinEQ(v bool) predicate.Perspective {
	return predicate.Perspective(sql.FieldEQ(FieldBuiltin, v))
}

// BuiltinNEQ applies the NEQ predicate on the "builtin" field.
func BuiltinNEQ(v bool) predicate.Perspective {
	return predicate.Perspective(sql.FieldNEQ(FieldBuiltin, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Perspective) predicate.Perspective {
	return predicate.Perspective(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Perspective) predicate.Perspective {
	return predicate.Perspective(func(s *sql.Selector) {
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
func Not(p predicate.Perspective) predicate.Perspective {
	return predicate.Perspective(func(s *sql.Selector) {
		p(s.Not())
	})
}