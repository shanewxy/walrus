package walrus

import (
	"context"
	"strings"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apiserver/pkg/registry/rest"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/extensionapi"
	"github.com/seal-io/walrus/pkg/systemmeta"
	"github.com/seal-io/walrus/pkg/templates/kubehelper"
)

// SchemaHandler handles v1.Schema objects.
type SchemaHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations

	client ctrlcli.Client
}

func (h *SchemaHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Declare GVR.
	gvr = walrus.SchemeGroupVersionResource("schemas")

	// Create table convertor to pretty the kubectl's output.
	var tc rest.TableConvertor
	{
		tc, err = extensionapi.NewJSONPathTableConvertor(
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Project",
					Type: "string",
				},
				JSONPath: ".status.project",
			},
		)
		if err != nil {
			return
		}
	}

	// As storage.
	h.ObjectInfo = &walrus.Schema{}
	h.CurdOperations = extensionapi.WithCurdProxy[
		*walrus.Schema, *walrus.SchemaList, *walruscore.Schema, *walruscore.SchemaList,
	](tc, h, opts.Manager.GetClient().(ctrlcli.WithWatch))

	// Set client.
	h.client = opts.Manager.GetClient()

	return
}

var (
	_ rest.Storage           = (*SchemaHandler)(nil)
	_ rest.Creater           = (*SchemaHandler)(nil)
	_ rest.Lister            = (*SchemaHandler)(nil)
	_ rest.Watcher           = (*SchemaHandler)(nil)
	_ rest.Getter            = (*SchemaHandler)(nil)
	_ rest.Updater           = (*SchemaHandler)(nil)
	_ rest.Patcher           = (*SchemaHandler)(nil)
	_ rest.GracefulDeleter   = (*SchemaHandler)(nil)
	_ rest.CollectionDeleter = (*SchemaHandler)(nil)
)

func (h *SchemaHandler) Destroy() {
}

func (h *SchemaHandler) New() runtime.Object {
	return &walrus.Schema{}
}

func (h *SchemaHandler) NewList() runtime.Object {
	return &walrus.SchemaList{}
}

func (h *SchemaHandler) CastObjectTo(do *walrus.Schema) (uo *walruscore.Schema) {
	uo = (*walruscore.Schema)(do)
	uo.APIVersion = walruscore.SchemeGroupVersion.String()
	return uo
}

func (h *SchemaHandler) CastObjectFrom(uo *walruscore.Schema) (do *walrus.Schema) {
	do = (*walrus.Schema)(uo)
	do.APIVersion = walrus.SchemeGroupVersion.String()
	return do
}

func (h *SchemaHandler) CastObjectListTo(dol *walrus.SchemaList) (uol *walruscore.SchemaList) {
	uol = (*walruscore.SchemaList)(dol)
	for i := range uol.Items {
		uol.Items[i].APIVersion = walruscore.SchemeGroupVersion.String()
	}
	return uol
}

func (h *SchemaHandler) CastObjectListFrom(uol *walruscore.SchemaList) (dol *walrus.SchemaList) {
	dol = (*walrus.SchemaList)(uol)
	for i := range dol.Items {
		dol.Items[i].APIVersion = walrus.SchemeGroupVersion.String()
	}
	return dol
}

func (h *SchemaHandler) OnGet(ctx context.Context, key types.NamespacedName, opts ctrlcli.GetOptions) (runtime.Object, error) {
	uo := h.CastObjectTo(h.New().(*walrus.Schema))
	err := h.client.Get(ctx, key, uo, &opts)
	return h.CastObjectFrom(uo), err
}

func (h *SchemaHandler) OnUpdate(ctx context.Context, obj, _ runtime.Object, opts ctrlcli.UpdateOptions) (runtime.Object, error) {
	s := obj.(*walrus.Schema)

	if apistatus.SchemaStatusReset.IsTrue(obj) {
		var (
			ouis    walruscore.Schema
			ouisKey = ctrlcli.ObjectKey{
				Namespace: s.Namespace,
				Name:      strings.TrimSuffix(s.Name, walruscore.NameSuffixUISchema) + walruscore.NameSuffixOriginalUISchema,
			}
		)
		err := h.client.Get(ctx, ouisKey, &ouis)
		if err != nil {
			return nil, err
		}

		apistatus.SchemaStatusReset.False(s, "", "")
		systemmeta.UnnoteResource(s)
		s.Status.Value = ouis.Status.Value
	} else {
		systemmeta.NoteResource(s, "", map[string]string{kubehelper.SchemaUserEditedNote: "true"})
	}

	uo := h.CastObjectTo(s)
	err := h.client.Update(ctx, uo, &opts)
	return h.CastObjectFrom(uo), err
}
