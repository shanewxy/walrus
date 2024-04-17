package walrus

import (
	"context"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
)

// ConnectorHandler handles v1.Connector objects.
//
// ConnectorHandler proxies the v1.Connector objects to the walrus core.
type ConnectorHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations
}

func (h *ConnectorHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err = fi.IndexField(ctx, &walruscore.Connector{}, "metadata.name",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			return []string{obj.GetName()}
		})
	if err != nil {
		return schema.GroupVersionResource{}, nil, err
	}

	// Declare GVR.
	gvr = walrus.SchemeGroupVersionResource("connectors")

	// Create table convertor to pretty the kubectl's output.
	var tc rest.TableConvertor
	{
		tc, err = extensionapi.NewJSONPathTableConvertor(
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Category",
					Type: "string",
				},
				JSONPath: ".spec.category",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Type",
					Type: "string",
				},
				JSONPath: ".spec.type",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Applicable Environment Type",
					Type: "string",
				},
				JSONPath: ".spec.applicableEnvironmentType",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Project",
					Type: "string",
				},
				JSONPath: ".status.project",
			})
		if err != nil {
			return gvr, nil, err
		}
	}

	// As storage.
	h.ObjectInfo = &walrus.Connector{}
	h.CurdOperations = extensionapi.WithCurdProxy[
		*walrus.Connector, *walrus.ConnectorList, *walruscore.Connector, *walruscore.ConnectorList,
	](tc, h, opts.Manager.GetClient().(ctrlcli.WithWatch), opts.Manager.GetAPIReader())

	return gvr, nil, nil
}

var (
	_ rest.Storage           = (*ConnectorHandler)(nil)
	_ rest.Creater           = (*ConnectorHandler)(nil)
	_ rest.Lister            = (*ConnectorHandler)(nil)
	_ rest.Watcher           = (*ConnectorHandler)(nil)
	_ rest.Getter            = (*ConnectorHandler)(nil)
	_ rest.Updater           = (*ConnectorHandler)(nil)
	_ rest.Patcher           = (*ConnectorHandler)(nil)
	_ rest.GracefulDeleter   = (*ConnectorHandler)(nil)
	_ rest.CollectionDeleter = (*ConnectorHandler)(nil)
)

func (h *ConnectorHandler) Destroy() {
}

func (h *ConnectorHandler) New() runtime.Object {
	return &walrus.Connector{}
}

func (h *ConnectorHandler) NewList() runtime.Object {
	return &walrus.ConnectorList{}
}

func (h *ConnectorHandler) NewListForProxy() runtime.Object {
	return &walruscore.ConnectorList{}
}

func (h *ConnectorHandler) CastObjectTo(do *walrus.Connector) (uo *walruscore.Connector) {
	return (*walruscore.Connector)(do)
}

func (h *ConnectorHandler) CastObjectFrom(uo *walruscore.Connector) (do *walrus.Connector) {
	return (*walrus.Connector)(uo)
}
