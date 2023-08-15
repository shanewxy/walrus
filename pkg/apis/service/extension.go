package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/zclconf/go-cty/cty"

	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/service"
	"github.com/seal-io/seal/pkg/dao/model/serviceresource"
	"github.com/seal-io/seal/pkg/dao/model/servicerevision"
	"github.com/seal-io/seal/pkg/dao/model/templateversion"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/object"
	"github.com/seal-io/seal/pkg/dao/types/property"
	"github.com/seal-io/seal/pkg/dao/types/status"
	"github.com/seal-io/seal/pkg/operator/k8s/intercept"
	optypes "github.com/seal-io/seal/pkg/operator/types"
	pkgservice "github.com/seal-io/seal/pkg/service"
	pkgresource "github.com/seal-io/seal/pkg/serviceresources"
	tfparser "github.com/seal-io/seal/pkg/terraform/parser"
	"github.com/seal-io/seal/pkg/topic/datamessage"
	"github.com/seal-io/seal/utils/strs"
	"github.com/seal-io/seal/utils/topic"
	"github.com/seal-io/seal/utils/validation"
)

func (h Handler) RouteUpgrade(req RouteUpgradeRequest) error {
	entity := req.Model()

	// Update service, mark status from deploying.
	status.ServiceStatusDeployed.Reset(entity, "Upgrading")
	entity.Status.SetSummary(status.WalkService(&entity.Status))

	err := h.modelClient.WithTx(req.Context, func(tx *model.Tx) (err error) {
		entity, err = tx.Services().UpdateOne(entity).
			Set(entity).
			SaveE(req.Context, dao.ServiceDependenciesEdgeSave)

		return err
	})
	if err != nil {
		return err
	}

	dp, err := h.getDeployer(req.Context)
	if err != nil {
		return err
	}

	// Apply service.
	applyOpts := pkgservice.Options{
		TlsCertified: h.tlsCertified,
		Tags:         req.RemarkTags,
	}

	return pkgservice.Apply(
		req.Context,
		h.modelClient,
		dp,
		entity,
		applyOpts)
}

func (h Handler) RouteRollback(req RouteRollbackRequest) error {
	rev, err := h.modelClient.ServiceRevisions().Query().
		Where(
			servicerevision.ID(req.RevisionID),
			servicerevision.ServiceID(req.ID)).
		WithService().
		Only(req.Context)
	if err != nil {
		return err
	}

	tv, err := h.modelClient.TemplateVersions().Query().
		Where(
			templateversion.Name(rev.TemplateName),
			templateversion.Version(rev.TemplateVersion)).
		Only(req.Context)
	if err != nil {
		return err
	}

	entity := rev.Edges.Service

	entity.Attributes = rev.Attributes
	entity.TemplateID = tv.ID
	status.ServiceStatusDeployed.Reset(entity, "Rolling back")
	entity.Status.SetSummary(status.WalkService(&entity.Status))

	entity, err = h.modelClient.Services().UpdateOne(entity).
		Set(entity).
		SaveE(req.Context, dao.ServiceDependenciesEdgeSave)
	if err != nil {
		return err
	}

	dp, err := h.getDeployer(req.Context)
	if err != nil {
		return err
	}

	applyOpts := pkgservice.Options{
		TlsCertified: h.tlsCertified,
	}

	return pkgservice.Apply(
		req.Context,
		h.modelClient,
		dp,
		entity,
		applyOpts)
}

func (h Handler) RouteGetAccessEndpoints(req RouteGetAccessEndpointsRequest) (RouteGetAccessEndpointsResponse, error) {
	if stream := req.Stream; stream != nil {
		t, err := topic.Subscribe(datamessage.ServiceRevision)
		if err != nil {
			return nil, err
		}

		defer func() { t.Unsubscribe() }()

		for {
			var event topic.Event

			event, err = t.Receive(stream)
			if err != nil {
				return nil, err
			}

			dm, ok := event.Data.(datamessage.Message[object.ID])
			if !ok {
				continue
			}

			if dm.Type == datamessage.EventDelete {
				continue
			}

			for _, id := range dm.Data {
				ar, err := h.modelClient.ServiceRevisions().Query().
					Where(servicerevision.ID(id)).
					Only(stream)
				if err != nil {
					return nil, err
				}

				if ar.ServiceID != req.ID {
					continue
				}

				eps, err := h.getAccessEndpoints(stream, req.ID)
				if err != nil {
					return nil, err
				}

				if len(eps) == 0 {
					continue
				}

				var resp *runtime.ResponseCollection

				switch dm.Type {
				case datamessage.EventCreate:
					// While create new service revision,
					// the previous endpoints from outputs and resources need to be deleted.
					resp = runtime.TypedResponse(datamessage.EventDelete.String(), eps)
				case datamessage.EventUpdate:
					// While the service revision status is succeeded,
					// the endpoints is updated to the current revision.
					if ar.Status != status.ServiceRevisionStatusSucceeded {
						continue
					}

					resp = runtime.TypedResponse(datamessage.EventUpdate.String(), eps)
				}

				if err = stream.SendJSON(resp); err != nil {
					return nil, err
				}
			}
		}
	}

	return h.getAccessEndpoints(req.Context, req.ID)
}

func (h Handler) getAccessEndpoints(ctx context.Context, id object.ID) ([]AccessEndpoint, error) {
	// Endpoints from output.
	eps, err := h.getEndpointsFromOutput(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(eps) != 0 {
		return eps, nil
	}

	// Endpoints from resources.
	return h.getEndpointsFromResources(ctx, id)
}

func (h Handler) getEndpointsFromOutput(ctx context.Context, id object.ID) ([]AccessEndpoint, error) {
	outputs, err := h.getServiceOutputs(ctx, id, true)
	if err != nil {
		return nil, err
	}

	var (
		invalidTypeErr = runtime.Error(http.StatusBadRequest,
			"element type of output endpoints should be string")
		endpoints = make([]AccessEndpoint, 0, len(outputs))
	)

	for _, v := range outputs {
		if !strings.HasPrefix(v.Name, "endpoint") {
			continue
		}

		prop := property.Property{
			Type:  v.Type,
			Value: v.Value,
		}

		switch {
		case v.Type == cty.String:
			ep, _, err := prop.GetString()
			if err != nil {
				return nil, err
			}

			if err := validation.IsValidEndpoint(ep); err != nil {
				return nil, runtime.Error(http.StatusBadRequest, err)
			}

			endpoints = append(endpoints, AccessEndpoint{
				Endpoints: []string{ep},
				Name:      v.Name,
			})
		case v.Type.IsListType() || v.Type.IsSetType() || v.Type.IsTupleType():
			if v.Type.IsTupleType() {
				// For tuple: each element has its own type.
				for _, tp := range v.Type.TupleElementTypes() {
					if tp != cty.String {
						return nil, invalidTypeErr
					}
				}
			} else if v.Type.ElementType() != cty.String {
				// For list and set: all elements are the same type.
				return nil, invalidTypeErr
			}

			eps, _, err := property.GetSlice[string](prop)
			if err != nil {
				return nil, err
			}

			if err := validation.IsValidEndpoints(eps); err != nil {
				return nil, runtime.Error(http.StatusBadRequest, err)
			}

			endpoints = append(endpoints, AccessEndpoint{
				Endpoints: eps,
				Name:      v.Name,
			})
		}
	}

	return endpoints, nil
}

func (h Handler) getEndpointsFromResources(ctx context.Context, id object.ID) ([]AccessEndpoint, error) {
	res, err := h.modelClient.ServiceResources().Query().
		Where(
			serviceresource.ServiceID(id),
			serviceresource.Mode(types.ServiceResourceModeManaged),
			serviceresource.TypeIn(intercept.TFEndpointsTypes...)).
		Select(
			serviceresource.FieldConnectorID,
			serviceresource.FieldType,
			serviceresource.FieldName,
			serviceresource.FieldStatus,
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	var endpoints []AccessEndpoint

	for _, v := range res {
		for _, eps := range v.Status.ResourceEndpoints {
			endpoints = append(endpoints, AccessEndpoint{
				Name:      strs.Join("/", eps.EndpointType, v.Name),
				Endpoints: eps.Endpoints,
			})
		}
	}

	return endpoints, nil
}

func (h Handler) RouteGetOutputs(req RouteGetOutputsRequest) (RouteGetOutputsResponse, error) {
	if stream := req.Stream; stream != nil {
		t, err := topic.Subscribe(datamessage.ServiceRevision)
		if err != nil {
			return nil, err
		}

		defer func() { t.Unsubscribe() }()

		for {
			var event topic.Event

			event, err = t.Receive(stream)
			if err != nil {
				return nil, err
			}

			dm, ok := event.Data.(datamessage.Message[object.ID])
			if !ok {
				continue
			}

			if dm.Type == datamessage.EventDelete {
				continue
			}

			for _, id := range dm.Data {
				ar, err := h.modelClient.ServiceRevisions().Query().
					Where(servicerevision.ID(id)).
					Only(stream)
				if err != nil {
					return nil, err
				}

				if ar.ServiceID != req.ID {
					continue
				}

				outs, err := h.getServiceOutputs(stream, ar.ServiceID, false)
				if err != nil {
					return nil, err
				}

				if len(outs) == 0 {
					continue
				}

				var resp *runtime.ResponseCollection

				switch dm.Type {
				case datamessage.EventCreate:
					// While create new service revision,
					// the outputs of new revision is the previous outputs.
					resp = runtime.TypedResponse(datamessage.EventDelete.String(), outs)
				case datamessage.EventUpdate:
					// While the service revision status is succeeded,
					// the outputs is updated to the current revision.
					if ar.Status != status.ServiceRevisionStatusSucceeded {
						continue
					}

					resp = runtime.TypedResponse(datamessage.EventUpdate.String(), outs)
				}

				if err = stream.SendJSON(resp); err != nil {
					return nil, err
				}
			}
		}
	}

	return h.getServiceOutputs(req.Context, req.ID, true)
}

func (h Handler) getServiceOutputs(ctx context.Context, id object.ID, onlySuccess bool) ([]types.OutputValue, error) {
	sr, err := h.modelClient.ServiceRevisions().Query().
		Where(servicerevision.ServiceID(id)).
		Select(
			servicerevision.FieldOutput,
			servicerevision.FieldTemplateName,
			servicerevision.FieldTemplateVersion,
			servicerevision.FieldAttributes,
			servicerevision.FieldStatus,
		).
		WithService(func(sq *model.ServiceQuery) {
			sq.Select(service.FieldName)
		}).
		Order(model.Desc(servicerevision.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, fmt.Errorf("error getting the latest service revision")
	}

	if sr == nil {
		return nil, nil
	}

	if onlySuccess && sr.Status != status.ServiceRevisionStatusSucceeded {
		return nil, nil
	}

	o, err := tfparser.ParseStateOutput(sr)
	if err != nil {
		return nil, fmt.Errorf("error get outputs: %w", err)
	}

	return o, nil
}

func (h Handler) RouteGetGraph(req RouteGetGraphRequest) (*RouteGetGraphResponse, error) {
	fields := []string{
		serviceresource.FieldServiceID,
		serviceresource.FieldDeployerType,
		serviceresource.FieldType,
		serviceresource.FieldID,
		serviceresource.FieldName,
		serviceresource.FieldMode,
		serviceresource.FieldShape,
		serviceresource.FieldClassID,
		serviceresource.FieldCreateTime,
		serviceresource.FieldUpdateTime,
		serviceresource.FieldStatus,
	}

	// Fetch entities.
	query := h.modelClient.ServiceResources().Query().
		Select(fields...).
		Order(model.Desc(serviceresource.FieldCreateTime)).
		Where(
			serviceresource.ServiceID(req.ID),
			serviceresource.Mode(types.ServiceResourceModeManaged),
			serviceresource.Shape(types.ServiceResourceShapeClass),
		)

	entities, err := dao.ServiceResourceShapeClassQuery(query, fields...).
		All(req.Context)
	if err != nil {
		return nil, err
	}

	// Calculate capacity for allocation.
	var verticesCap, edgesCap int
	{
		// Count the number of ServiceResource.
		verticesCap = len(entities)
		for i := 0; i < len(entities); i++ {
			// Count the vertex size of sub ServiceResource,
			// and the edge size from ServiceResource to sub ServiceResource.
			verticesCap += len(entities[i].Edges.Components)
			edgesCap += len(entities[i].Edges.Components)
		}
	}

	// Construct response.
	var (
		vertices  = make([]GraphVertex, 0, verticesCap)
		edges     = make([]GraphEdge, 0, edgesCap)
		operators = make(map[object.ID]optypes.Operator)
	)

	pkgresource.SetKeys(req.Context, entities, operators)
	vertices, edges = pkgresource.GetVerticesAndEdges(entities, vertices, edges)

	return &RouteGetGraphResponse{
		Vertices: vertices,
		Edges:    edges,
	}, nil
}