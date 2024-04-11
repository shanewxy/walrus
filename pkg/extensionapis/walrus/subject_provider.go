package walrus

import (
	"context"
	"slices"
	"sort"
	"strings"

	"github.com/seal-io/utils/json"
	"github.com/seal-io/utils/pools/gopool"
	"github.com/seal-io/utils/stringx"
	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// SubjectProviderHandler handles v1.SubjectProvider objects.
//
// SubjectProviderHandler maps the v1.SubjectProvider object to a Kubernetes Secret resource,
// which is named as "walrus-subjectprovider-<name>".
type SubjectProviderHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations

	Client    ctrlcli.Client
	APIReader ctrlcli.Reader
}

func (h *SubjectProviderHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err = fi.IndexField(ctx, &core.Secret{}, "metadata.name",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			return []string{obj.GetName()}
		})
	if err != nil {
		return
	}

	// Declare GVR.
	gvr = walrus.SchemeGroupVersionResource("subjectproviders")

	// Create table convertor to pretty the kubectl's output.
	var tc rest.TableConvertor
	{
		tc, err = extensionapi.NewJSONPathTableConvertor(
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Type",
					Type: "string",
				},
				JSONPath: ".spec.type",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Password Login",
					Type: "boolean",
				},
				JSONPath: ".status.loginWithPassword",
			})
		if err != nil {
			return
		}
	}

	// As storage.
	h.ObjectInfo = &walrus.SubjectProvider{}
	h.CurdOperations = extensionapi.WithCurd(tc, h)

	// Set client.
	h.Client = opts.Manager.GetClient()
	h.APIReader = opts.Manager.GetAPIReader()

	return
}

var (
	_ rest.Storage           = (*SubjectProviderHandler)(nil)
	_ rest.Creater           = (*SubjectProviderHandler)(nil)
	_ rest.Lister            = (*SubjectProviderHandler)(nil)
	_ rest.Watcher           = (*SubjectProviderHandler)(nil)
	_ rest.Getter            = (*SubjectProviderHandler)(nil)
	_ rest.Updater           = (*SubjectProviderHandler)(nil)
	_ rest.Patcher           = (*SubjectProviderHandler)(nil)
	_ rest.GracefulDeleter   = (*SubjectProviderHandler)(nil)
	_ rest.CollectionDeleter = (*SubjectProviderHandler)(nil)
)

func (h *SubjectProviderHandler) New() runtime.Object {
	return &walrus.SubjectProvider{}
}

func (h *SubjectProviderHandler) Destroy() {
}

func (h *SubjectProviderHandler) OnCreate(ctx context.Context, obj runtime.Object, opts ctrlcli.CreateOptions) (runtime.Object, error) {
	// Validate.
	subjProv := obj.(*walrus.SubjectProvider)
	{
		var errs field.ErrorList
		if subjProv.Namespace != systemkuberes.SystemNamespaceName {
			errs = append(errs, field.Invalid(
				field.NewPath("metadata.namespace"), subjProv.Namespace, "subject provider namespace must be "+systemkuberes.SystemNamespaceName))
		}
		if stringx.StringWidth(subjProv.Name) > 30 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("metadata.name"), stringx.StringWidth(subjProv.Name), 30))
		}
		switch {
		case subjProv.Name != systemkuberes.DefaultSubjectProviderName && subjProv.Spec.Type == walrus.SubjectProviderTypeInternal:
			errs = append(errs, field.Invalid(
				field.NewPath("spec.type"), subjProv.Spec.Type, "internal subject provider must be named as "+systemkuberes.DefaultSubjectProviderName))
		case subjProv.Name == systemkuberes.DefaultSubjectProviderName && subjProv.Spec.Type != walrus.SubjectProviderTypeInternal:
			errs = append(errs, field.Invalid(
				field.NewPath("spec.type"), subjProv.Spec.Type, "default subject provider must be internal"))
		}
		if err := subjProv.Spec.Type.Validate(); err != nil {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.type"), subjProv.Spec.Type, err.Error()))
		}
		if stringx.StringWidth(subjProv.Spec.DisplayName) > 30 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.displayName"), stringx.StringWidth(subjProv.Spec.DisplayName), 30))
		}
		if stringx.StringWidth(subjProv.Spec.Description) > 50 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.description"), stringx.StringWidth(subjProv.Spec.Description), 50))
		}
		if err := subjProv.Spec.ExternalConfig.ValidateWithType(subjProv.Spec.Type); err != nil {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.externalConfig"), subjProv.Spec.ExternalConfig, err.Error()))
		}
		if len(errs) > 0 {
			return nil, kerrors.NewInvalid(walrus.SchemeKind("subjectproviders"), subjProv.Name, errs)
		}
	}

	// Create.
	{
		sec := convertSecretFromSubjectProvider(subjProv)
		err := h.Client.Create(ctx, sec, &opts)
		if err != nil {
			return nil, err
		}
		subjProv = convertSubjectProviderFromSecret(sec)
	}

	return subjProv, nil
}

func (h *SubjectProviderHandler) NewList() runtime.Object {
	return &walrus.SubjectProviderList{}
}

func (h *SubjectProviderHandler) OnList(ctx context.Context, opts ctrlcli.ListOptions) (runtime.Object, error) {
	// List.
	secList := new(core.SecretList)
	err := h.APIReader.List(ctx, secList,
		convertSecretListOptsFromSubjectProviderListOpts(opts))
	if err != nil {
		return nil, err
	}

	// Convert.
	spList := convertSubjectProviderListFromSecretList(secList, opts)
	return spList, nil
}

func (h *SubjectProviderHandler) OnWatch(ctx context.Context, opts ctrlcli.ListOptions) (watch.Interface, error) {
	// Watch.
	uw, err := h.Client.(ctrlcli.WithWatch).Watch(ctx, new(core.SecretList),
		convertSecretListOptsFromSubjectProviderListOpts(opts))
	if err != nil {
		return nil, err
	}

	c := make(chan watch.Event)
	dw := watch.NewProxyWatcher(c)
	gopool.Go(func() {
		defer close(c)
		defer uw.Stop()

		for {
			select {
			case <-ctx.Done():
				// Cancel by context.
				return
			case <-dw.StopChan():
				// Stop by downstream.
				return
			case e, ok := <-uw.ResultChan():
				if !ok {
					// Close by upstream.
					return
				}

				// Nothing to do.
				if e.Object == nil {
					c <- e
					continue
				}

				// Type assert.
				sec, ok := e.Object.(*core.Secret)
				if !ok {
					c <- e
					continue
				}

				// Process bookmark.
				if e.Type == watch.Bookmark {
					e.Object = &walrus.SubjectProvider{ObjectMeta: sec.ObjectMeta}
					c <- e
					continue
				}

				// Convert.
				subj := safeConvertSubjectProviderFromSecret(sec, opts.Namespace)
				if subj == nil {
					continue
				}

				// Dispatch.
				e.Object = subj
				c <- e
			}
		}
	})

	return dw, nil
}

func (h *SubjectProviderHandler) OnGet(ctx context.Context, key types.NamespacedName, opts ctrlcli.GetOptions) (runtime.Object, error) {
	// Get.
	sec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: key.Namespace,
			Name:      _SubjectProviderDelegatedSecretNamePrefix + key.Name,
		},
	}
	err := h.APIReader.Get(ctx, ctrlcli.ObjectKeyFromObject(sec), sec, &opts)
	if err != nil {
		return nil, err
	}

	// Convert.
	proj := convertSubjectProviderFromSecret(sec)
	if proj == nil {
		return nil, kerrors.NewNotFound(walrus.SchemeResource("subjectproviders"), key.Name)
	}
	return proj, nil
}

func (h *SubjectProviderHandler) OnUpdate(ctx context.Context, obj, oldObj runtime.Object, opts ctrlcli.UpdateOptions) (runtime.Object, error) {
	// Validate.
	subjProv := obj.(*walrus.SubjectProvider)
	{
		oldSubjProvider := oldObj.(*walrus.SubjectProvider)
		var errs field.ErrorList
		if subjProv.Spec.Type != oldSubjProvider.Spec.Type {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.type"), subjProv.Spec.Type, "type is immutable"))
		}
		if stringx.StringWidth(subjProv.Spec.DisplayName) > 30 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.displayName"), stringx.StringWidth(subjProv.Spec.DisplayName), 30))
		}
		if stringx.StringWidth(subjProv.Spec.Description) > 50 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.description"), stringx.StringWidth(subjProv.Spec.Description), 50))
		}
		if err := subjProv.Spec.ExternalConfig.ValidateWithType(subjProv.Spec.Type); err != nil {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.externalConfig"), subjProv.Spec.ExternalConfig, err.Error()))
		}
		if len(errs) > 0 {
			return nil, kerrors.NewInvalid(walrus.SchemeKind("subjectproviders"), subjProv.Name, errs)
		}
	}

	// Update.
	{
		sec := convertSecretFromSubjectProvider(subjProv)
		err := h.Client.Update(ctx, sec, &opts)
		if err != nil {
			return nil, err
		}
		subjProv = convertSubjectProviderFromSecret(sec)
	}

	return subjProv, nil
}

func (h *SubjectProviderHandler) OnDelete(ctx context.Context, obj runtime.Object, opts ctrlcli.DeleteOptions) error {
	subjProv := obj.(*walrus.SubjectProvider)

	// Validate.
	{
		// Prevent deleting default subject provider.
		if subjProv.Name == systemkuberes.DefaultSubjectProviderName {
			return kerrors.NewBadRequest("default subject provider is reserved")
		}
	}

	// Delete.
	sec := convertSecretFromSubjectProvider(subjProv)
	return h.Client.Delete(ctx, sec, &opts)
}

func convertSecretListOptsFromSubjectProviderListOpts(in ctrlcli.ListOptions) (out *ctrlcli.ListOptions) {
	if in.Namespace != systemkuberes.SystemNamespaceName {
		return &in
	}

	// Ignore name selector
	if in.FieldSelector != nil {
		reqs := slices.DeleteFunc(in.FieldSelector.Requirements(), func(req fields.Requirement) bool {
			return req.Field == "metadata.name"
		})
		if len(reqs) == 0 {
			in.FieldSelector = nil
		} else {
			in.FieldSelector = kubemeta.FieldSelectorFromRequirements(reqs)
		}
	}

	// Add necessary label selector.
	if lbs := systemmeta.GetResourcesLabelSelectorOfType("subjectproviders"); in.LabelSelector == nil {
		in.LabelSelector = lbs
	} else {
		reqs, _ := lbs.Requirements()
		in.LabelSelector = in.LabelSelector.DeepCopySelector().Add(reqs...)
	}

	return &in
}

const _SubjectProviderDelegatedSecretNamePrefix = "walrus-subjectprovider-"

func convertSecretFromSubjectProvider(subjProv *walrus.SubjectProvider) *core.Secret {
	sec := &core.Secret{
		ObjectMeta: subjProv.ObjectMeta,
	}
	systemmeta.NoteResource(sec, "subjectproviders", map[string]string{
		"type":        subjProv.Spec.Type.String(),
		"displayName": subjProv.Spec.DisplayName,
		"description": subjProv.Spec.Description,
	})
	sec.Data = map[string][]byte{
		"externalConfig": json.MustMarshal(subjProv.Spec.ExternalConfig),
	}
	sec.Name = _SubjectProviderDelegatedSecretNamePrefix + subjProv.Name
	return sec
}

func convertSubjectProviderFromSecret(sec *core.Secret) *walrus.SubjectProvider {
	if sec == nil {
		return nil
	}

	resType, notes := systemmeta.UnnoteResource(sec)
	if resType != "subjectproviders" {
		return nil
	}
	if !strings.HasPrefix(sec.Name, _SubjectProviderDelegatedSecretNamePrefix) {
		return nil
	}

	subjProv := &walrus.SubjectProvider{
		ObjectMeta: sec.ObjectMeta,
		Spec: walrus.SubjectProviderSpec{
			Type:        walrus.SubjectProviderType(notes["type"]),
			DisplayName: notes["displayName"],
			Description: notes["description"],
		},
	}
	if sec.Data != nil && sec.Data["externalConfig"] != nil {
		json.ShouldUnmarshal(sec.Data["externalConfig"], &subjProv.Spec.ExternalConfig)
	}
	subjProv.Name = strings.TrimPrefix(sec.Name, _SubjectProviderDelegatedSecretNamePrefix)
	switch subjProv.Spec.Type {
	case walrus.SubjectProviderTypeInternal, walrus.SubjectProviderTypeLDAP:
		subjProv.Status.LoginWithPassword = true
	}
	return subjProv
}

func safeConvertSubjectProviderFromSecret(sec *core.Secret, reqNamespace string) *walrus.SubjectProvider {
	subjProv := convertSubjectProviderFromSecret(sec)
	if subjProv != nil && reqNamespace != "" && reqNamespace != subjProv.Namespace {
		// NB(thxCode): sanitize if the subject provider's namespace doesn't match requested namespace.
		subjProv = nil
	}
	return subjProv
}

func convertSubjectProviderListFromSecretList(secList *core.SecretList, opts ctrlcli.ListOptions) *walrus.SubjectProviderList {
	if secList == nil {
		return &walrus.SubjectProviderList{}
	}

	// Sort by resource version.
	sort.SliceStable(secList.Items, func(i, j int) bool {
		l, r := secList.Items[i].ResourceVersion, secList.Items[j].ResourceVersion
		return len(l) < len(r) ||
			(len(l) == len(r) && l < r)
	})

	spList := &walrus.SubjectProviderList{
		Items: make([]walrus.SubjectProvider, 0, len(secList.Items)),
	}

	for i := range secList.Items {
		subjProv := safeConvertSubjectProviderFromSecret(&secList.Items[i], opts.Namespace)
		if subjProv == nil {
			continue
		}
		// Ignore if not be selected by `kubectl get --field-selector=metadata.namespace=...`.
		if fs := opts.FieldSelector; fs != nil &&
			!fs.Matches(fields.Set{"metadata.namespace": subjProv.Namespace, "metadata.name": subjProv.Name}) {
			continue
		}
		spList.Items = append(spList.Items, *subjProv)
	}

	return spList
}
