// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package servicerevision

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
)

const (
	// Label holds the string label denoting the servicerevision type in the database.
	Label = "service_revision"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProjectID holds the string denoting the projectid field in the database.
	FieldProjectID = "project_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStatusMessage holds the string denoting the statusmessage field in the database.
	FieldStatusMessage = "status_message"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldServiceID holds the string denoting the serviceid field in the database.
	FieldServiceID = "service_id"
	// FieldEnvironmentID holds the string denoting the environmentid field in the database.
	FieldEnvironmentID = "environment_id"
	// FieldTemplateID holds the string denoting the templateid field in the database.
	FieldTemplateID = "template_id"
	// FieldTemplateVersion holds the string denoting the templateversion field in the database.
	FieldTemplateVersion = "template_version"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// FieldSecrets holds the string denoting the secrets field in the database.
	FieldSecrets = "secrets"
	// FieldInputPlan holds the string denoting the inputplan field in the database.
	FieldInputPlan = "input_plan"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldDeployerType holds the string denoting the deployertype field in the database.
	FieldDeployerType = "deployer_type"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldPreviousRequiredProviders holds the string denoting the previousrequiredproviders field in the database.
	FieldPreviousRequiredProviders = "previous_required_providers"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// EdgeEnvironment holds the string denoting the environment edge name in mutations.
	EdgeEnvironment = "environment"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the servicerevision in the database.
	Table = "service_revisions"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "service_revisions"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "service_id"
	// EnvironmentTable is the table that holds the environment relation/edge.
	EnvironmentTable = "service_revisions"
	// EnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	EnvironmentInverseTable = "environments"
	// EnvironmentColumn is the table column denoting the environment relation/edge.
	EnvironmentColumn = "environment_id"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "service_revisions"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
)

// Columns holds all SQL columns for servicerevision fields.
var Columns = []string{
	FieldID,
	FieldProjectID,
	FieldStatus,
	FieldStatusMessage,
	FieldCreateTime,
	FieldServiceID,
	FieldEnvironmentID,
	FieldTemplateID,
	FieldTemplateVersion,
	FieldAttributes,
	FieldSecrets,
	FieldInputPlan,
	FieldOutput,
	FieldDeployerType,
	FieldDuration,
	FieldPreviousRequiredProviders,
	FieldTags,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/seal-io/seal/pkg/dao/model/runtime"
var (
	Hooks        [3]ent.Hook
	Interceptors [1]ent.Interceptor
	// ProjectIDValidator is a validator for the "projectID" field. It is called by the builders before save.
	ProjectIDValidator func(string) error
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime func() time.Time
	// ServiceIDValidator is a validator for the "serviceID" field. It is called by the builders before save.
	ServiceIDValidator func(string) error
	// EnvironmentIDValidator is a validator for the "environmentID" field. It is called by the builders before save.
	EnvironmentIDValidator func(string) error
	// TemplateIDValidator is a validator for the "templateID" field. It is called by the builders before save.
	TemplateIDValidator func(string) error
	// TemplateVersionValidator is a validator for the "templateVersion" field. It is called by the builders before save.
	TemplateVersionValidator func(string) error
	// DefaultSecrets holds the default value on creation for the "secrets" field.
	DefaultSecrets crypto.Map[string, string]
	// DefaultDeployerType holds the default value on creation for the "deployerType" field.
	DefaultDeployerType string
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
	// DefaultPreviousRequiredProviders holds the default value on creation for the "previousRequiredProviders" field.
	DefaultPreviousRequiredProviders []types.ProviderRequirement
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
)

// OrderOption defines the ordering options for the ServiceRevision queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProjectID orders the results by the projectID field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByStatusMessage orders the results by the statusMessage field.
func ByStatusMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusMessage, opts...).ToFunc()
}

// ByCreateTime orders the results by the createTime field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByServiceID orders the results by the serviceID field.
func ByServiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServiceID, opts...).ToFunc()
}

// ByEnvironmentID orders the results by the environmentID field.
func ByEnvironmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnvironmentID, opts...).ToFunc()
}

// ByTemplateID orders the results by the templateID field.
func ByTemplateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateID, opts...).ToFunc()
}

// ByTemplateVersion orders the results by the templateVersion field.
func ByTemplateVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateVersion, opts...).ToFunc()
}

// ByAttributes orders the results by the attributes field.
func ByAttributes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttributes, opts...).ToFunc()
}

// BySecrets orders the results by the secrets field.
func BySecrets(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecrets, opts...).ToFunc()
}

// ByInputPlan orders the results by the inputPlan field.
func ByInputPlan(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInputPlan, opts...).ToFunc()
}

// ByOutput orders the results by the output field.
func ByOutput(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutput, opts...).ToFunc()
}

// ByDeployerType orders the results by the deployerType field.
func ByDeployerType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeployerType, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByServiceField orders the results by service field.
func ByServiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByEnvironmentField orders the results by environment field.
func ByEnvironmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEnvironmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}
func newServiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
	)
}
func newEnvironmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EnvironmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EnvironmentTable, EnvironmentColumn),
	)
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return Columns
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}