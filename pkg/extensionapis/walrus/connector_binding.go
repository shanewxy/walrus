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

// ConnectorBindingHandler handles v1.ConnectorBinding objects.
//
// ConnectorBindingHandler proxies the v1.ConnectorBinding objects to the walrus core.
type ConnectorBindingHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations
}

func (h *ConnectorBindingHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err = fi.IndexField(ctx, &walruscore.ConnectorBinding{}, "metadata.name",
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
	gvr = walrus.SchemeGroupVersionResource("connectorbindings")

	// Create table convertor to pretty the kubectl's output.
	var tc rest.TableConvertor
	{
		tc, err = extensionapi.NewJSONPathTableConvertor(
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Connector",
					Type: "string",
				},
				JSONPath: ".spec.connector.name",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Connector Namespace",
					Type: "string",
				},
				JSONPath: ".spec.connector.namespace",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Connector Type",
					Type: "string",
				},
				JSONPath: ".status.Type",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Connector Category",
					Type: "string",
				},
				JSONPath: ".status.Category",
			},
		)
		if err != nil {
			return gvr, nil, err
		}
	}

	// As storage.
	h.ObjectInfo = &walrus.ConnectorBinding{}
	h.CurdOperations = extensionapi.WithCurdProxy[
		*walrus.ConnectorBinding, *walrus.ConnectorBindingList, *walruscore.ConnectorBinding, *walruscore.ConnectorBindingList,
	](tc, h, opts.Manager.GetClient().(ctrlcli.WithWatch), opts.Manager.GetAPIReader())

	return gvr, nil, nil
}

var (
	_ rest.Storage           = (*ConnectorBindingHandler)(nil)
	_ rest.Creater           = (*ConnectorBindingHandler)(nil)
	_ rest.Lister            = (*ConnectorBindingHandler)(nil)
	_ rest.Getter            = (*ConnectorBindingHandler)(nil)
	_ rest.Updater           = (*ConnectorBindingHandler)(nil)
	_ rest.Patcher           = (*ConnectorBindingHandler)(nil)
	_ rest.Watcher           = (*ConnectorBindingHandler)(nil)
	_ rest.CollectionDeleter = (*ConnectorBindingHandler)(nil)
	_ rest.GracefulDeleter   = (*ConnectorBindingHandler)(nil)
)

func (h *ConnectorBindingHandler) New() runtime.Object {
	return &walrus.ConnectorBinding{}
}

func (h *ConnectorBindingHandler) Destroy() {
}

func (h *ConnectorBindingHandler) NewList() runtime.Object {
	return &walrus.ConnectorBindingList{}
}

func (h *ConnectorBindingHandler) NewListForProxy() runtime.Object {
	return &walruscore.ConnectorBindingList{}
}

func (h *ConnectorBindingHandler) CastObjectTo(do *walrus.ConnectorBinding) *walruscore.ConnectorBinding {
	return (*walruscore.ConnectorBinding)(do)
}

func (h *ConnectorBindingHandler) CastObjectFrom(uo *walruscore.ConnectorBinding) *walrus.ConnectorBinding {
	return (*walrus.ConnectorBinding)(uo)
}
