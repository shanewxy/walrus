// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/authrequest"
)

// AuthRequestCreate is the builder for creating a AuthRequest entity.
type AuthRequestCreate struct {
	config
	mutation *AuthRequestMutation
	hooks    []Hook
}

// SetClientID sets the "client_id" field.
func (arc *AuthRequestCreate) SetClientID(s string) *AuthRequestCreate {
	arc.mutation.SetClientID(s)
	return arc
}

// SetScopes sets the "scopes" field.
func (arc *AuthRequestCreate) SetScopes(s []string) *AuthRequestCreate {
	arc.mutation.SetScopes(s)
	return arc
}

// SetResponseTypes sets the "response_types" field.
func (arc *AuthRequestCreate) SetResponseTypes(s []string) *AuthRequestCreate {
	arc.mutation.SetResponseTypes(s)
	return arc
}

// SetRedirectURI sets the "redirect_uri" field.
func (arc *AuthRequestCreate) SetRedirectURI(s string) *AuthRequestCreate {
	arc.mutation.SetRedirectURI(s)
	return arc
}

// SetNonce sets the "nonce" field.
func (arc *AuthRequestCreate) SetNonce(s string) *AuthRequestCreate {
	arc.mutation.SetNonce(s)
	return arc
}

// SetState sets the "state" field.
func (arc *AuthRequestCreate) SetState(s string) *AuthRequestCreate {
	arc.mutation.SetState(s)
	return arc
}

// SetForceApprovalPrompt sets the "force_approval_prompt" field.
func (arc *AuthRequestCreate) SetForceApprovalPrompt(b bool) *AuthRequestCreate {
	arc.mutation.SetForceApprovalPrompt(b)
	return arc
}

// SetLoggedIn sets the "logged_in" field.
func (arc *AuthRequestCreate) SetLoggedIn(b bool) *AuthRequestCreate {
	arc.mutation.SetLoggedIn(b)
	return arc
}

// SetClaimsUserID sets the "claims_user_id" field.
func (arc *AuthRequestCreate) SetClaimsUserID(s string) *AuthRequestCreate {
	arc.mutation.SetClaimsUserID(s)
	return arc
}

// SetClaimsUsername sets the "claims_username" field.
func (arc *AuthRequestCreate) SetClaimsUsername(s string) *AuthRequestCreate {
	arc.mutation.SetClaimsUsername(s)
	return arc
}

// SetClaimsEmail sets the "claims_email" field.
func (arc *AuthRequestCreate) SetClaimsEmail(s string) *AuthRequestCreate {
	arc.mutation.SetClaimsEmail(s)
	return arc
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (arc *AuthRequestCreate) SetClaimsEmailVerified(b bool) *AuthRequestCreate {
	arc.mutation.SetClaimsEmailVerified(b)
	return arc
}

// SetClaimsGroups sets the "claims_groups" field.
func (arc *AuthRequestCreate) SetClaimsGroups(s []string) *AuthRequestCreate {
	arc.mutation.SetClaimsGroups(s)
	return arc
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (arc *AuthRequestCreate) SetClaimsPreferredUsername(s string) *AuthRequestCreate {
	arc.mutation.SetClaimsPreferredUsername(s)
	return arc
}

// SetNillableClaimsPreferredUsername sets the "claims_preferred_username" field if the given value is not nil.
func (arc *AuthRequestCreate) SetNillableClaimsPreferredUsername(s *string) *AuthRequestCreate {
	if s != nil {
		arc.SetClaimsPreferredUsername(*s)
	}
	return arc
}

// SetConnectorID sets the "connector_id" field.
func (arc *AuthRequestCreate) SetConnectorID(s string) *AuthRequestCreate {
	arc.mutation.SetConnectorID(s)
	return arc
}

// SetConnectorData sets the "connector_data" field.
func (arc *AuthRequestCreate) SetConnectorData(b []byte) *AuthRequestCreate {
	arc.mutation.SetConnectorData(b)
	return arc
}

// SetExpiry sets the "expiry" field.
func (arc *AuthRequestCreate) SetExpiry(t time.Time) *AuthRequestCreate {
	arc.mutation.SetExpiry(t)
	return arc
}

// SetCodeChallenge sets the "code_challenge" field.
func (arc *AuthRequestCreate) SetCodeChallenge(s string) *AuthRequestCreate {
	arc.mutation.SetCodeChallenge(s)
	return arc
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (arc *AuthRequestCreate) SetNillableCodeChallenge(s *string) *AuthRequestCreate {
	if s != nil {
		arc.SetCodeChallenge(*s)
	}
	return arc
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (arc *AuthRequestCreate) SetCodeChallengeMethod(s string) *AuthRequestCreate {
	arc.mutation.SetCodeChallengeMethod(s)
	return arc
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (arc *AuthRequestCreate) SetNillableCodeChallengeMethod(s *string) *AuthRequestCreate {
	if s != nil {
		arc.SetCodeChallengeMethod(*s)
	}
	return arc
}

// SetHmacKey sets the "hmac_key" field.
func (arc *AuthRequestCreate) SetHmacKey(b []byte) *AuthRequestCreate {
	arc.mutation.SetHmacKey(b)
	return arc
}

// SetID sets the "id" field.
func (arc *AuthRequestCreate) SetID(s string) *AuthRequestCreate {
	arc.mutation.SetID(s)
	return arc
}

// Mutation returns the AuthRequestMutation object of the builder.
func (arc *AuthRequestCreate) Mutation() *AuthRequestMutation {
	return arc.mutation
}

// Save creates the AuthRequest in the database.
func (arc *AuthRequestCreate) Save(ctx context.Context) (*AuthRequest, error) {
	arc.defaults()
	return withHooks(ctx, arc.sqlSave, arc.mutation, arc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AuthRequestCreate) SaveX(ctx context.Context) *AuthRequest {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AuthRequestCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AuthRequestCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AuthRequestCreate) defaults() {
	if _, ok := arc.mutation.ClaimsPreferredUsername(); !ok {
		v := authrequest.DefaultClaimsPreferredUsername
		arc.mutation.SetClaimsPreferredUsername(v)
	}
	if _, ok := arc.mutation.CodeChallenge(); !ok {
		v := authrequest.DefaultCodeChallenge
		arc.mutation.SetCodeChallenge(v)
	}
	if _, ok := arc.mutation.CodeChallengeMethod(); !ok {
		v := authrequest.DefaultCodeChallengeMethod
		arc.mutation.SetCodeChallengeMethod(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AuthRequestCreate) check() error {
	if _, ok := arc.mutation.ClientID(); !ok {
		return &ValidationError{Name: "client_id", err: errors.New(`db: missing required field "AuthRequest.client_id"`)}
	}
	if _, ok := arc.mutation.RedirectURI(); !ok {
		return &ValidationError{Name: "redirect_uri", err: errors.New(`db: missing required field "AuthRequest.redirect_uri"`)}
	}
	if _, ok := arc.mutation.Nonce(); !ok {
		return &ValidationError{Name: "nonce", err: errors.New(`db: missing required field "AuthRequest.nonce"`)}
	}
	if _, ok := arc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`db: missing required field "AuthRequest.state"`)}
	}
	if _, ok := arc.mutation.ForceApprovalPrompt(); !ok {
		return &ValidationError{Name: "force_approval_prompt", err: errors.New(`db: missing required field "AuthRequest.force_approval_prompt"`)}
	}
	if _, ok := arc.mutation.LoggedIn(); !ok {
		return &ValidationError{Name: "logged_in", err: errors.New(`db: missing required field "AuthRequest.logged_in"`)}
	}
	if _, ok := arc.mutation.ClaimsUserID(); !ok {
		return &ValidationError{Name: "claims_user_id", err: errors.New(`db: missing required field "AuthRequest.claims_user_id"`)}
	}
	if _, ok := arc.mutation.ClaimsUsername(); !ok {
		return &ValidationError{Name: "claims_username", err: errors.New(`db: missing required field "AuthRequest.claims_username"`)}
	}
	if _, ok := arc.mutation.ClaimsEmail(); !ok {
		return &ValidationError{Name: "claims_email", err: errors.New(`db: missing required field "AuthRequest.claims_email"`)}
	}
	if _, ok := arc.mutation.ClaimsEmailVerified(); !ok {
		return &ValidationError{Name: "claims_email_verified", err: errors.New(`db: missing required field "AuthRequest.claims_email_verified"`)}
	}
	if _, ok := arc.mutation.ClaimsPreferredUsername(); !ok {
		return &ValidationError{Name: "claims_preferred_username", err: errors.New(`db: missing required field "AuthRequest.claims_preferred_username"`)}
	}
	if _, ok := arc.mutation.ConnectorID(); !ok {
		return &ValidationError{Name: "connector_id", err: errors.New(`db: missing required field "AuthRequest.connector_id"`)}
	}
	if _, ok := arc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`db: missing required field "AuthRequest.expiry"`)}
	}
	if _, ok := arc.mutation.CodeChallenge(); !ok {
		return &ValidationError{Name: "code_challenge", err: errors.New(`db: missing required field "AuthRequest.code_challenge"`)}
	}
	if _, ok := arc.mutation.CodeChallengeMethod(); !ok {
		return &ValidationError{Name: "code_challenge_method", err: errors.New(`db: missing required field "AuthRequest.code_challenge_method"`)}
	}
	if _, ok := arc.mutation.HmacKey(); !ok {
		return &ValidationError{Name: "hmac_key", err: errors.New(`db: missing required field "AuthRequest.hmac_key"`)}
	}
	if v, ok := arc.mutation.ID(); ok {
		if err := authrequest.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`db: validator failed for field "AuthRequest.id": %w`, err)}
		}
	}
	return nil
}

func (arc *AuthRequestCreate) sqlSave(ctx context.Context) (*AuthRequest, error) {
	if err := arc.check(); err != nil {
		return nil, err
	}
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AuthRequest.ID type: %T", _spec.ID.Value)
		}
	}
	arc.mutation.id = &_node.ID
	arc.mutation.done = true
	return _node, nil
}

func (arc *AuthRequestCreate) createSpec() (*AuthRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthRequest{config: arc.config}
		_spec = sqlgraph.NewCreateSpec(authrequest.Table, sqlgraph.NewFieldSpec(authrequest.FieldID, field.TypeString))
	)
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := arc.mutation.ClientID(); ok {
		_spec.SetField(authrequest.FieldClientID, field.TypeString, value)
		_node.ClientID = value
	}
	if value, ok := arc.mutation.Scopes(); ok {
		_spec.SetField(authrequest.FieldScopes, field.TypeJSON, value)
		_node.Scopes = value
	}
	if value, ok := arc.mutation.ResponseTypes(); ok {
		_spec.SetField(authrequest.FieldResponseTypes, field.TypeJSON, value)
		_node.ResponseTypes = value
	}
	if value, ok := arc.mutation.RedirectURI(); ok {
		_spec.SetField(authrequest.FieldRedirectURI, field.TypeString, value)
		_node.RedirectURI = value
	}
	if value, ok := arc.mutation.Nonce(); ok {
		_spec.SetField(authrequest.FieldNonce, field.TypeString, value)
		_node.Nonce = value
	}
	if value, ok := arc.mutation.State(); ok {
		_spec.SetField(authrequest.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := arc.mutation.ForceApprovalPrompt(); ok {
		_spec.SetField(authrequest.FieldForceApprovalPrompt, field.TypeBool, value)
		_node.ForceApprovalPrompt = value
	}
	if value, ok := arc.mutation.LoggedIn(); ok {
		_spec.SetField(authrequest.FieldLoggedIn, field.TypeBool, value)
		_node.LoggedIn = value
	}
	if value, ok := arc.mutation.ClaimsUserID(); ok {
		_spec.SetField(authrequest.FieldClaimsUserID, field.TypeString, value)
		_node.ClaimsUserID = value
	}
	if value, ok := arc.mutation.ClaimsUsername(); ok {
		_spec.SetField(authrequest.FieldClaimsUsername, field.TypeString, value)
		_node.ClaimsUsername = value
	}
	if value, ok := arc.mutation.ClaimsEmail(); ok {
		_spec.SetField(authrequest.FieldClaimsEmail, field.TypeString, value)
		_node.ClaimsEmail = value
	}
	if value, ok := arc.mutation.ClaimsEmailVerified(); ok {
		_spec.SetField(authrequest.FieldClaimsEmailVerified, field.TypeBool, value)
		_node.ClaimsEmailVerified = value
	}
	if value, ok := arc.mutation.ClaimsGroups(); ok {
		_spec.SetField(authrequest.FieldClaimsGroups, field.TypeJSON, value)
		_node.ClaimsGroups = value
	}
	if value, ok := arc.mutation.ClaimsPreferredUsername(); ok {
		_spec.SetField(authrequest.FieldClaimsPreferredUsername, field.TypeString, value)
		_node.ClaimsPreferredUsername = value
	}
	if value, ok := arc.mutation.ConnectorID(); ok {
		_spec.SetField(authrequest.FieldConnectorID, field.TypeString, value)
		_node.ConnectorID = value
	}
	if value, ok := arc.mutation.ConnectorData(); ok {
		_spec.SetField(authrequest.FieldConnectorData, field.TypeBytes, value)
		_node.ConnectorData = &value
	}
	if value, ok := arc.mutation.Expiry(); ok {
		_spec.SetField(authrequest.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	if value, ok := arc.mutation.CodeChallenge(); ok {
		_spec.SetField(authrequest.FieldCodeChallenge, field.TypeString, value)
		_node.CodeChallenge = value
	}
	if value, ok := arc.mutation.CodeChallengeMethod(); ok {
		_spec.SetField(authrequest.FieldCodeChallengeMethod, field.TypeString, value)
		_node.CodeChallengeMethod = value
	}
	if value, ok := arc.mutation.HmacKey(); ok {
		_spec.SetField(authrequest.FieldHmacKey, field.TypeBytes, value)
		_node.HmacKey = value
	}
	return _node, _spec
}

// AuthRequestCreateBulk is the builder for creating many AuthRequest entities in bulk.
type AuthRequestCreateBulk struct {
	config
	err      error
	builders []*AuthRequestCreate
}

// Save creates the AuthRequest entities in the database.
func (arcb *AuthRequestCreateBulk) Save(ctx context.Context) ([]*AuthRequest, error) {
	if arcb.err != nil {
		return nil, arcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AuthRequest, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AuthRequestCreateBulk) SaveX(ctx context.Context) []*AuthRequest {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AuthRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AuthRequestCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}