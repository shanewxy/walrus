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

// CatalogHandler handles v1.Catalog objects.
//
// CatalogHandler proxies the v1.Catalog objects to the walrus core.
type CatalogHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations

	client ctrlcli.Client
}

func (h *CatalogHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Declare GVR.
	gvr = walrus.SchemeGroupVersionResource("catalogs")

	// Create table convertor to pretty the kubectl's output.
	var tc rest.TableConvertor
	{
		tc, err = extensionapi.NewJSONPathTableConvertor(
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Format",
					Type: "string",
				},
				JSONPath: ".spec.templateFormat",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "URL",
					Type: "string",
				},
				JSONPath: ".status.url",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Project",
					Type: "string",
				},
				JSONPath: ".status.project",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Status",
					Type: "string",
				},
				JSONPath: ".status.phase",
			},
		)
		if err != nil {
			return gvr, srs, err
		}
	}

	// As storage.
	h.ObjectInfo = &walrus.Catalog{}
	h.CurdOperations = extensionapi.WithCurdProxy[
		*walrus.Catalog, *walrus.CatalogList, *walruscore.Catalog, *walruscore.CatalogList,
	](tc, h, opts.Manager.GetClient().(ctrlcli.WithWatch))

	// Set client.
	h.client = opts.Manager.GetClient()
	return gvr, srs, nil
}

var (
	_ rest.Storage           = (*CatalogHandler)(nil)
	_ rest.Creater           = (*CatalogHandler)(nil)
	_ rest.Lister            = (*CatalogHandler)(nil)
	_ rest.Watcher           = (*CatalogHandler)(nil)
	_ rest.Getter            = (*CatalogHandler)(nil)
	_ rest.Updater           = (*CatalogHandler)(nil)
	_ rest.Patcher           = (*CatalogHandler)(nil)
	_ rest.GracefulDeleter   = (*CatalogHandler)(nil)
	_ rest.CollectionDeleter = (*CatalogHandler)(nil)
)

func (h *CatalogHandler) Destroy() {
}

func (h *CatalogHandler) New() runtime.Object {
	return &walrus.Catalog{}
}

func (h *CatalogHandler) NewList() runtime.Object {
	return &walrus.CatalogList{}
}

func (h *CatalogHandler) CastObjectTo(do *walrus.Catalog) (uo *walruscore.Catalog) {
	return (*walruscore.Catalog)(do)
}

func (h *CatalogHandler) CastObjectFrom(uo *walruscore.Catalog) (do *walrus.Catalog) {
	do = (*walrus.Catalog)(uo)
	do.APIVersion = walrus.SchemeGroupVersion.String()
	return do
}

func (h *CatalogHandler) CastObjectListTo(dol *walrus.CatalogList) (uol *walruscore.CatalogList) {
	uol = (*walruscore.CatalogList)(dol)
	for i := range dol.Items {
		dol.Items[i].APIVersion = walruscore.SchemeGroupVersion.String()
	}
	return uol
}

func (h *CatalogHandler) CastObjectListFrom(uol *walruscore.CatalogList) (dol *walrus.CatalogList) {
	dol = (*walrus.CatalogList)(uol)
	for i := range dol.Items {
		dol.Items[i].APIVersion = walrus.SchemeGroupVersion.String()
	}
	return dol
}
