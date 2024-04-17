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

// TemplateHandler handles v1.Template objects.
//
// TemplateHandler proxies the v1.Template objects to the walrus core.
type TemplateHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations

	client ctrlcli.Client
}

func (h *TemplateHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Declare GVR.
	gvr = walrus.SchemeGroupVersionResource("templates")

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
	h.ObjectInfo = &walrus.Template{}
	h.CurdOperations = extensionapi.WithCurdProxy[
		*walrus.Template, *walrus.TemplateList, *walruscore.Template, *walruscore.TemplateList,
	](tc, h, opts.Manager.GetClient().(ctrlcli.WithWatch))

	// Set client.
	h.client = opts.Manager.GetClient()
	return gvr, srs, nil
}

var (
	_ rest.Storage           = (*TemplateHandler)(nil)
	_ rest.Creater           = (*TemplateHandler)(nil)
	_ rest.Lister            = (*TemplateHandler)(nil)
	_ rest.Watcher           = (*TemplateHandler)(nil)
	_ rest.Getter            = (*TemplateHandler)(nil)
	_ rest.Updater           = (*TemplateHandler)(nil)
	_ rest.Patcher           = (*TemplateHandler)(nil)
	_ rest.GracefulDeleter   = (*TemplateHandler)(nil)
	_ rest.CollectionDeleter = (*TemplateHandler)(nil)
)

func (h *TemplateHandler) Destroy() {
}

func (h *TemplateHandler) New() runtime.Object {
	return &walrus.Template{}
}

func (h *TemplateHandler) NewList() runtime.Object {
	return &walrus.TemplateList{}
}

func (h *TemplateHandler) NewListForProxy() runtime.Object {
	return &walruscore.TemplateList{}
}

func (h *TemplateHandler) CastObjectTo(do *walrus.Template) (uo *walruscore.Template) {
	return (*walruscore.Template)(do)
}

func (h *TemplateHandler) CastObjectFrom(uo *walruscore.Template) (do *walrus.Template) {
	return (*walrus.Template)(uo)
}
