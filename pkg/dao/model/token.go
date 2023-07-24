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

	"github.com/seal-io/seal/pkg/dao/model/subject"
	"github.com/seal-io/seal/pkg/dao/model/token"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// Token is the model entity for the Token schema.
type Token struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// ID of the subject to belong.
	SubjectID object.ID `json:"subject_id,omitempty"`
	// The kind of token.
	Kind string `json:"kind,omitempty"`
	// The name of token.
	Name string `json:"name,omitempty"`
	// The time of expiration, empty means forever.
	Expiration *time.Time `json:"expiration,omitempty"`
	// The value of token, store in string.
	Value crypto.String `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TokenQuery when eager-loading is set.
	Edges        TokenEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues

	// AccessToken is the token used for authentication.
	// AccessToken does not store in the database.
	AccessToken string `json:"access_token,omitempty"`
}

// TokenEdges holds the relations/edges for other nodes in the graph.
type TokenEdges struct {
	// Subject to which the token belongs.
	Subject *Subject `json:"subject,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TokenEdges) SubjectOrErr() (*Subject, error) {
	if e.loadedTypes[0] {
		if e.Subject == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: subject.Label}
		}
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Token) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case token.FieldValue:
			values[i] = new(crypto.String)
		case token.FieldID, token.FieldSubjectID:
			values[i] = new(object.ID)
		case token.FieldKind, token.FieldName:
			values[i] = new(sql.NullString)
		case token.FieldCreateTime, token.FieldExpiration:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Token fields.
func (t *Token) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case token.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case token.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = new(time.Time)
				*t.CreateTime = value.Time
			}
		case token.FieldSubjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field subject_id", values[i])
			} else if value != nil {
				t.SubjectID = *value
			}
		case token.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				t.Kind = value.String
			}
		case token.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case token.FieldExpiration:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiration", values[i])
			} else if value.Valid {
				t.Expiration = new(time.Time)
				*t.Expiration = value.Time
			}
		case token.FieldValue:
			if value, ok := values[i].(*crypto.String); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value != nil {
				t.Value = *value
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Token.
// This includes values selected through modifiers, order, etc.
func (t *Token) GetValue(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QuerySubject queries the "subject" edge of the Token entity.
func (t *Token) QuerySubject() *SubjectQuery {
	return NewTokenClient(t.config).QuerySubject(t)
}

// Update returns a builder for updating this Token.
// Note that you need to call Token.Unwrap() before calling this method if this Token
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Token) Update() *TokenUpdateOne {
	return NewTokenClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Token entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Token) Unwrap() *Token {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("model: Token is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Token) String() string {
	var builder strings.Builder
	builder.WriteString("Token(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	if v := t.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("subject_id=")
	builder.WriteString(fmt.Sprintf("%v", t.SubjectID))
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(t.Kind)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	if v := t.Expiration; v != nil {
		builder.WriteString("expiration=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("value=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// Tokens is a parsable slice of Token.
type Tokens []*Token
