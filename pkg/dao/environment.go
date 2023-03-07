package dao

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/utils/strs"
)

// WrappedEnvironmentCreate is a wrapper for model.EnvironmentCreate
// to process the relationship with model.Connector.
// TODO(thxCode): generate this with entc.
type WrappedEnvironmentCreate struct {
	entity   *model.Environment
	delegate *model.EnvironmentCreate
}

func (ec *WrappedEnvironmentCreate) Save(ctx context.Context) (created *model.Environment, err error) {
	var mc = ec.delegate.Mutation().Client()

	// save entity.
	created, err = ec.delegate.Save(ctx)
	if err != nil {
		return
	}

	// construct relationships.
	var newRss = ec.entity.Edges.Connectors
	var createRss = make([]*model.EnvironmentConnectorRelationshipCreate, len(newRss))
	for i, rs := range newRss {
		if rs == nil {
			return nil, errors.New("invalid input: nil relationship")
		}

		// required.
		var c = mc.EnvironmentConnectorRelationships().Create().
			SetEnvironmentID(created.ID).
			SetConnectorID(rs.ConnectorID)

		createRss[i] = c
	}

	// save relationships.
	newRss, err = mc.EnvironmentConnectorRelationships().CreateBulk(createRss...).
		Save(ctx)
	if err != nil {
		return
	}
	created.Edges.Connectors = newRss
	return
}

func (ec *WrappedEnvironmentCreate) Exec(ctx context.Context) error {
	var _, err = ec.Save(ctx)
	return err
}

func EnvironmentCreates(mc model.ClientSet, input ...*model.Environment) ([]*WrappedEnvironmentCreate, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid input: empty list")
	}

	var rrs = make([]*WrappedEnvironmentCreate, len(input))
	for i, r := range input {
		if r == nil {
			return nil, errors.New("invalid input: nil entity")
		}

		// required.
		var c = mc.Environments().Create().
			SetName(r.Name)

		// optional.
		c.SetDescription(r.Description)
		if r.Labels != nil {
			c.SetLabels(r.Labels)
		}
		rrs[i] = &WrappedEnvironmentCreate{
			entity:   input[i],
			delegate: c,
		}
	}
	return rrs, nil
}

// WrappedEnvironmentUpdate is a wrapper for model.EnvironmentUpdate
// to process the relationship with model.Connector.
// TODO(thxCode): generate this with entc.
type WrappedEnvironmentUpdate struct {
	entity           *model.Environment
	entityPredicates []predicate.Environment
	delegate         *model.EnvironmentUpdate
}

func (eu *WrappedEnvironmentUpdate) Save(ctx context.Context) (updated int, err error) {
	var mc = eu.delegate.Mutation().Client()

	if len(eu.delegate.Mutation().Fields()) != 0 {
		// update entity.
		updated, err = eu.delegate.Save(ctx)
		if err != nil {
			return
		}
	}
	if eu.entity.Edges.Connectors == nil {
		return
	}

	// get old relationships.
	oldEntity, err := mc.Environments().Query().
		Where(eu.entityPredicates...).
		Select(environment.FieldID).
		WithConnectors(func(eq *model.EnvironmentConnectorRelationshipQuery) {
			eq.Select(
				environmentconnectorrelationship.FieldEnvironmentID,
				environmentconnectorrelationship.FieldConnectorID,
			)
		}).
		Only(ctx)
	if err != nil {
		return
	}

	// create new relationship or update relationship.
	var environmentID = oldEntity.ID
	var newRsKeys = sets.New[string]()
	var newRss = eu.entity.Edges.Connectors
	for _, rs := range newRss {
		newRsKeys.Insert(strs.Join("/", string(environmentID), string(rs.ConnectorID)))

		// required.
		var c = mc.EnvironmentConnectorRelationships().Create().
			SetEnvironmentID(environmentID).
			SetConnectorID(rs.ConnectorID)

		err = c.OnConflict(
			sql.ConflictColumns(
				environmentconnectorrelationship.FieldEnvironmentID,
				environmentconnectorrelationship.FieldConnectorID,
			)).
			DoNothing().
			Exec(ctx)
		if err != nil {
			return
		}
	}

	// delete stale relationship.
	var oldRss = oldEntity.Edges.Connectors
	for _, rs := range oldRss {
		if newRsKeys.Has(strs.Join("/", string(rs.EnvironmentID), string(rs.ConnectorID))) {
			continue
		}

		_, err = mc.EnvironmentConnectorRelationships().Delete().
			Where(
				environmentconnectorrelationship.EnvironmentID(rs.EnvironmentID),
				environmentconnectorrelationship.ConnectorID(rs.ConnectorID),
			).
			Exec(ctx)
		if err != nil {
			return
		}
	}

	return
}

func (eu *WrappedEnvironmentUpdate) Exec(ctx context.Context) error {
	var _, err = eu.Save(ctx)
	return err
}

func EnvironmentUpdates(mc model.ClientSet, input ...*model.Environment) ([]*WrappedEnvironmentUpdate, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid input: empty list")
	}

	var rrs = make([]*WrappedEnvironmentUpdate, len(input))
	for i, r := range input {
		if r == nil {
			return nil, errors.New("invalid input: nil entity")
		}

		// predicated.
		var ps []predicate.Environment
		switch {
		case r.ID.IsNaive():
			ps = append(ps, environment.ID(r.ID))
		case r.Name != "":
			ps = append(ps, environment.Name(r.Name))
		}
		if len(ps) == 0 {
			return nil, errors.New("invalid input: illegal predicates")
		}

		// conditional.
		var c = mc.Environments().Update().
			Where(ps...).
			SetDescription(r.Description)
		if r.Name != "" {
			c.SetName(r.Name)
		}
		if r.Labels != nil {
			c.SetLabels(r.Labels)
		}
		rrs[i] = &WrappedEnvironmentUpdate{
			entity:           input[i],
			entityPredicates: ps,
			delegate:         c,
		}
	}
	return rrs, nil
}