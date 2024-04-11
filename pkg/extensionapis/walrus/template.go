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
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
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

func (h *TemplateHandler) CastObjectTo(do *walrus.Template) (uo *walruscore.Template) {
	uo = (*walruscore.Template)(do)
	uo.APIVersion = walruscore.SchemeGroupVersion.String()
	return uo
}

func (h *TemplateHandler) CastObjectFrom(uo *walruscore.Template) (do *walrus.Template) {
	do = (*walrus.Template)(uo)
	do.APIVersion = walrus.SchemeGroupVersion.String()
	return do
}

func (h *TemplateHandler) CastObjectListTo(dol *walrus.TemplateList) (uol *walruscore.TemplateList) {
	uol = (*walruscore.TemplateList)(dol)
	for i := range uol.Items {
		uol.Items[i].APIVersion = walruscore.SchemeGroupVersion.String()
	}
	return uol
}

func (h *TemplateHandler) CastObjectListFrom(uol *walruscore.TemplateList) (dol *walrus.TemplateList) {
	dol = (*walrus.TemplateList)(uol)
	for i := range dol.Items {
		dol.Items[i].APIVersion = walrus.SchemeGroupVersion.String()
	}
	return dol
}

func (h *TemplateHandler) OnCreate(ctx context.Context, obj runtime.Object, opts ctrlcli.CreateOptions) (runtime.Object, error) {
	// Validate.
	t := obj.(*walrus.Template)
	_, err := sourceurl.ParseURLToSourceURL(t.Spec.VCSRepository.URL)
	if err != nil {
		return nil, err
	}

	uo := h.CastObjectTo(t)
	err = h.client.Create(ctx, uo, &opts)
	return h.CastObjectFrom(uo), err
}
