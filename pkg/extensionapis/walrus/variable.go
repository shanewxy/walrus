package walrus

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/seal-io/utils/pools/gopool"
	"golang.org/x/exp/maps"
	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/utils/ptr"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// VariableHandler handles v1.Variable objects.
//
// VariableHandler maps all v1.Variable objects to a Kubernetes Secret resource,
// which is named as "${namespace}/walrus-variables".
//
// Each v1.Variable object records as a key-value pair in the Secret's Data field.
type VariableHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations

	Client    ctrlcli.Client
	APIReader ctrlcli.Reader
}

func (h *VariableHandler) SetupHandler(
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
	gvr = walrus.SchemeGroupVersionResource("variables")

	// Create table convertor to pretty the kubectl's output.
	var tc rest.TableConvertor
	{
		tc, err = extensionapi.NewJSONPathTableConvertor(
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Value",
					Type: "string",
				},
				JSONPath: ".status.value",
			},
			extensionapi.JSONPathTableColumnDefinition{
				TableColumnDefinition: meta.TableColumnDefinition{
					Name: "Scope",
					Type: "string",
				},
				JSONPath: ".status.scope",
			})
		if err != nil {
			return
		}
	}

	// As storage.
	h.ObjectInfo = &walrus.Variable{}
	h.CurdOperations = extensionapi.WithCurd(tc, h)

	// Set client.
	h.Client = opts.Manager.GetClient()
	h.APIReader = opts.Manager.GetAPIReader()

	return
}

var (
	_ rest.Storage           = (*VariableHandler)(nil)
	_ rest.Creater           = (*VariableHandler)(nil)
	_ rest.Lister            = (*VariableHandler)(nil)
	_ rest.Watcher           = (*VariableHandler)(nil)
	_ rest.Getter            = (*VariableHandler)(nil)
	_ rest.Updater           = (*VariableHandler)(nil)
	_ rest.Patcher           = (*VariableHandler)(nil)
	_ rest.GracefulDeleter   = (*VariableHandler)(nil)
	_ rest.CollectionDeleter = (*VariableHandler)(nil)
)

func (h *VariableHandler) New() runtime.Object {
	return &walrus.Variable{}
}

func (h *VariableHandler) Destroy() {}

func (h *VariableHandler) OnCreate(ctx context.Context, obj runtime.Object, opts ctrlcli.CreateOptions) (runtime.Object, error) {
	// Validate.
	vra := obj.(*walrus.Variable)
	{
		var errs field.ErrorList
		if vra.Spec.Value == nil {
			errs = field.ErrorList{
				field.Required(field.NewPath("spec.value"), "variable value is required"),
			}
		}
		ns, err := systemmeta.ReflectNamespace(ctx, h.Client, vra.Namespace)
		if err != nil {
			errs = field.ErrorList{
				field.Invalid(field.NewPath("metadata.namespace"), vra.Namespace, err.Error()),
			}
		} else {
			vra.Status.Scope = walrus.VariableScope(ns.Kind())
		}
		if len(errs) > 0 {
			return nil, kerrors.NewInvalid(walrus.SchemeKind("variables"), vra.Name, errs)
		}
	}

	// Update or Create.
	eSec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: vra.Namespace,
			Name:      systemkuberes.VariablesDelegatedSecretName,
		},
		Data: map[string][]byte{
			vra.Name: []byte(*vra.Spec.Value),
		},
	}
	eResType := "variables"
	eNotes := map[string]string{
		"scope":                 string(vra.Status.Scope),
		vra.Name + "-uid":       uuid.NewString(),
		vra.Name + "-create-at": time.Now().Format(time.RFC3339),
		vra.Name + "-sensitive": strconv.FormatBool(vra.Spec.Sensitive),
	}
	systemmeta.NoteResource(eSec, eResType, eNotes)
	alignFn := func(aSec *core.Secret) (*core.Secret, bool, error) {
		// Validate.
		if aSec.Data == nil {
			aSec.Data = make(map[string][]byte)
		}
		if _, ok := aSec.Data[vra.Name]; ok {
			return nil, true, kerrors.NewAlreadyExists(walrus.SchemeResource("variables"), vra.Name)
		}
		// Align data.
		aSec.Data[vra.Name] = eSec.Data[vra.Name]
		// Align delegated info.
		_, aNotes := systemmeta.DescribeResource(aSec)
		eNotes := maps.Clone(eNotes)
		maps.Copy(eNotes, aNotes)
		systemmeta.NoteResource(aSec, eResType, eNotes)
		return aSec, false, nil
	}

	sec, err := kubeclientset.UpdateWithCtrlClient(ctx, h.Client, eSec,
		kubeclientset.WithUpdateAlign(alignFn),
		kubeclientset.WithCreateIfNotExisted[*core.Secret]())
	if err != nil {
		return nil, err
	}

	// Convert.
	vra = convertVariableFromSecret(sec, vra.Name)
	return vra, nil
}

func (h *VariableHandler) NewList() runtime.Object {
	return &walrus.VariableList{}
}

func (h *VariableHandler) OnList(ctx context.Context, opts ctrlcli.ListOptions) (runtime.Object, error) {
	// List.
	if opts.Namespace == "" {
		secList := new(core.SecretList)
		err := h.Client.List(ctx, secList,
			convertSecretListOptsFromVariableListOpts(opts))
		if err != nil {
			return nil, err
		}

		// Convert.
		return convertVariableListFromSecretList(secList, opts), nil
	}

	// Get.
	sec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: opts.Namespace,
			Name:      systemkuberes.VariablesDelegatedSecretName,
		},
	}
	err := h.APIReader.Get(ctx, ctrlcli.ObjectKeyFromObject(sec), sec)
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return nil, err
		}
		// We return an empty list if the secret is not found.
		return &walrus.VariableList{}, nil
	}

	scp := systemmeta.DescribeResourceNote(sec, "scope")
	if scp == "" {
		return &walrus.VariableList{}, nil
	}
	ns, err := systemmeta.ReflectNamespace(ctx, h.Client, sec.Namespace)
	if err != nil {
		return &walrus.VariableList{}, kerrors.NewInternalError(err)
	}

	secList := &core.SecretList{
		Items: []core.Secret{*sec},
	}
	switch ns.Kind() {
	case systemmeta.NamespaceKindEnvironment:
		projSec := &core.Secret{
			ObjectMeta: meta.ObjectMeta{
				Namespace: ns.OwnerName(),
				Name:      systemkuberes.VariablesDelegatedSecretName,
			},
		}
		err = h.APIReader.Get(ctx, ctrlcli.ObjectKeyFromObject(projSec), projSec)
		if err != nil {
			if !kerrors.IsNotFound(err) {
				return nil, kerrors.NewInternalError(fmt.Errorf("get project variables: %w", err))
			}
		}
		ns, err = ns.RetrieveOwner(ctx)
		if err != nil {
			return nil, kerrors.NewInternalError(fmt.Errorf("get project owner: %w", err))
		}
		secList.Items = append(secList.Items, *projSec)
		fallthrough
	case systemmeta.NamespaceKindProject:
		systemSec := &core.Secret{
			ObjectMeta: meta.ObjectMeta{
				Namespace: ns.OwnerName(),
				Name:      systemkuberes.VariablesDelegatedSecretName,
			},
		}
		err = h.APIReader.Get(ctx, ctrlcli.ObjectKeyFromObject(systemSec), systemSec)
		if err != nil {
			if !kerrors.IsNotFound(err) {
				return nil, kerrors.NewInternalError(fmt.Errorf("get system variables: %w", err))
			}
		}
		secList.Items = append(secList.Items, *systemSec)
	}

	// Merge.
	vList := mergeVariableListFromSecretList(secList, opts)
	return vList, nil
}

func (h *VariableHandler) OnWatch(ctx context.Context, opts ctrlcli.ListOptions) (watch.Interface, error) {
	namespaceFilter := sets.New[string]()
	if opts.Namespace != "" {
		ns, err := systemmeta.ReflectNamespace(ctx, h.Client, opts.Namespace)
		if err != nil {
			return nil, kerrors.NewInternalError(err)
		}
		namespaceFilter.Insert(opts.Namespace)
		switch ns.Kind() {
		case systemmeta.NamespaceKindEnvironment:
			namespaceFilter.Insert(ns.OwnerName())
			ns, err = ns.RetrieveOwner(ctx)
			if err != nil {
				return nil, kerrors.NewInternalError(fmt.Errorf("get project owner: %w", err))
			}
			fallthrough
		case systemmeta.NamespaceKindProject:
			namespaceFilter.Insert(ns.OwnerName())
		}
	}

	// Index.
	vraIndexer := map[string]walrus.Variable{}         // [vnamespace/vname] -> vra
	vraReverseIndexer := map[string]sets.Set[string]{} // [vname] -> [vnamespace, ...]
	{
		listObj, err := h.OnList(ctx, opts)
		if err != nil {
			return nil, err
		}
		vList := listObj.(*walrus.VariableList)
		for i := range vList.Items {
			vraIndexKey := kubemeta.GetNamespacedNameKey(&vList.Items[i])
			vraIndexer[vraIndexKey] = vList.Items[i]
			if _, ok := vraReverseIndexer[vList.Items[i].Name]; !ok {
				vraReverseIndexer[vList.Items[i].Name] = sets.New[string]()
			}
			vraReverseIndexer[vList.Items[i].Name].Insert(vList.Items[i].Namespace)
		}
	}

	// Watch.
	wopts := convertSecretListOptsFromVariableListOpts(opts)
	wopts.Namespace = ""
	uw, err := h.Client.(ctrlcli.WithWatch).Watch(ctx, new(core.SecretList), wopts)
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

				// Nothing to do
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
					e.Object = &walrus.Variable{ObjectMeta: sec.ObjectMeta}
					c <- e
					continue
				}

				// Disallow if not matched.
				if opts.Namespace != "" && !namespaceFilter.Has(sec.Namespace) {
					continue
				}

				vraL1IndexKey := sec.Namespace + "/"
				vraIndexKeySet := sets.New[string]()

				// Send.
				for name := range sec.Data {
					// Ignore if not be selected by `kubectl get --field-selector=metadata.namespace=...,metadata.name=...`.
					if fs := opts.FieldSelector; fs != nil &&
						!fs.Matches(fields.Set{"metadata.namespace": sec.Namespace, "metadata.name": name}) {
						continue
					}

					// Convert.
					vra := convertVariableFromSecret(sec, name)
					if vra == nil {
						continue
					}

					vraIndexKey := kubemeta.GetNamespacedNameKey(vra)
					vraIndexKeySet.Insert(vraIndexKey)

					prevVra, ok := vraIndexer[vraIndexKey]
					switch {
					default:
						// Ignore if the same as previous.
						continue
					case !ok:
						// Add or update if not exist.
						vraIndexer[vraIndexKey] = *vra
						if vraReverseIndexer[vra.Name] == nil {
							vraReverseIndexer[vra.Name] = sets.New[string]()
							vraReverseIndexer[vra.Name].Insert(vra.Namespace)
							e2 := e.DeepCopy()
							e2.Object = vra
							e2.Type = watch.Added
							c <- *e2
							continue
						}
						vraReverseIndexer[vra.Name].Insert(vra.Namespace)
						if getHighestPriorityVariableNamespace(vraReverseIndexer[vra.Name]) != vra.Namespace {
							continue
						}
						e2 := e.DeepCopy()
						e2.Object = vra
						e2.Type = watch.Modified
						c <- *e2
					case !vra.Equal(&prevVra):
						// Update if changed.
						vraIndexer[vraIndexKey] = *vra
						if opts.Namespace != "" && vra.Namespace != opts.Namespace {
							continue
						}
						e2 := e.DeepCopy()
						e2.Object = vra
						e2.Type = watch.Modified
						c <- *e2
					}
				}

				// Clean up and get notified.
				var notifiedVras []walrus.Variable
				for vraIndexKey := range vraIndexer {
					if !strings.HasPrefix(vraIndexKey, vraL1IndexKey) {
						continue
					}
					switch {
					default:
						continue
					case e.Type == watch.Deleted:
					case !vraIndexKeySet.Has(vraIndexKey):
					}

					// Delete if not exist.
					vra := vraIndexer[vraIndexKey]
					delete(vraIndexer, vraIndexKey)
					if vraReverseIndexer[vra.Name] != nil {
						vraReverseIndexer[vra.Name].Delete(vra.Namespace)
						if vraReverseIndexer[vra.Name].Len() == 0 {
							vraReverseIndexer[vra.Name] = nil
						}
					}

					// Ignore if not be selected by `kubectl get --field-selector=metadata.namespace=...,metadata.name=...`.
					if fs := opts.FieldSelector; fs != nil &&
						!fs.Matches(fields.Set{"metadata.namespace": vra.Namespace, "metadata.name": vra.Name}) {
						continue
					}

					vra.ResourceVersion = sec.ResourceVersion
					vra.Generation = sec.Generation
					notifiedVras = append(notifiedVras, vra)
				}
				if len(notifiedVras) == 0 {
					continue
				}

				// Dispatch notified event.
				for i := range notifiedVras {
					e2 := e.DeepCopy()
					vra := &notifiedVras[i]
					if opts.Namespace != "" && vraReverseIndexer[vra.Name] != nil {
						ns := getHighestPriorityVariableNamespace(vraReverseIndexer[vra.Name])
						if ns == opts.Namespace {
							continue
						}
						vra = ptr.To(vraIndexer[ns+"/"+vra.Name])
						e2.Type = watch.Modified
					} else {
						vra.DeletionTimestamp = ptr.To(meta.NewTime(time.Now()))
						vra.DeletionGracePeriodSeconds = ptr.To[int64](0)
						e2.Type = watch.Deleted
					}
					vra.ResourceVersion = sec.ResourceVersion
					vra.Generation = sec.Generation
					e2.Object = vra
					c <- *e2
				}
			}
		}
	})

	return dw, nil
}

func (h *VariableHandler) OnGet(ctx context.Context, name types.NamespacedName, opts ctrlcli.GetOptions) (runtime.Object, error) {
	// Get.
	sec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: name.Namespace,
			Name:      systemkuberes.VariablesDelegatedSecretName,
		},
	}
	err := h.APIReader.Get(ctx, ctrlcli.ObjectKeyFromObject(sec), sec)
	if err != nil {
		return nil, err
	}

	// Convert.
	vra := convertVariableFromSecret(sec, name.Name)
	if vra == nil {
		return nil, kerrors.NewNotFound(walrus.SchemeResource("variables"), name.Name)
	}
	return vra, nil
}

func (h *VariableHandler) OnUpdate(ctx context.Context, obj, _ runtime.Object, opts ctrlcli.UpdateOptions) (runtime.Object, error) {
	// Validate.
	vra := obj.(*walrus.Variable)
	if vra.Spec.Value == nil {
		errs := field.ErrorList{
			field.Required(field.NewPath("spec.value"), "variable value is required"),
		}
		return nil, kerrors.NewInvalid(walrus.SchemeKind("variables"), vra.Name, errs)
	}

	// Update.
	eSec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: vra.Namespace,
			Name:      systemkuberes.VariablesDelegatedSecretName,
		},
		Data: map[string][]byte{
			vra.Name: []byte(*vra.Spec.Value),
		},
	}
	eNotes := map[string]string{
		vra.Name + "-sensitive": strconv.FormatBool(vra.Spec.Sensitive),
	}
	systemmeta.NoteResource(eSec, "variables", eNotes)
	alignFn := func(aSec *core.Secret) (*core.Secret, bool, error) {
		// Validate.
		if aSec.Data == nil || aSec.Data[vra.Name] == nil {
			return nil, true, kerrors.NewNotFound(walrus.SchemeResource("variables"), vra.Name)
		}
		// Align data.
		aSec.Data[vra.Name] = eSec.Data[vra.Name]
		// Align delegated info.
		systemmeta.NoteResource(aSec, "variables", eNotes)
		return aSec, false, nil
	}

	sec, err := kubeclientset.UpdateWithCtrlClient(ctx, h.Client, eSec,
		kubeclientset.WithUpdateAlign(alignFn))
	if err != nil {
		return nil, err
	}

	// Convert.
	vra = convertVariableFromSecret(sec, vra.Name)
	if vra == nil {
		return nil, kerrors.NewNotFound(walrus.SchemeResource("variables"), vra.Name)
	}
	return vra, nil
}

func (h *VariableHandler) OnDelete(ctx context.Context, obj runtime.Object, opts ctrlcli.DeleteOptions) error {
	vra := obj.(*walrus.Variable)

	// Update.
	eSec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: vra.Namespace,
			Name:      systemkuberes.VariablesDelegatedSecretName,
		},
	}
	alignFn := func(aSec *core.Secret) (*core.Secret, bool, error) {
		// Validate.
		if aSec.Data == nil || aSec.Data[vra.Name] == nil {
			return nil, true, kerrors.NewNotFound(walrus.SchemeResource("variables"), vra.Name)
		}
		// Align data.
		delete(aSec.Data, vra.Name)
		// Align delegated info.
		systemmeta.PopResourceNotes(aSec, []string{
			vra.Name + "-uid",
			vra.Name + "-create-at",
			vra.Name + "-sensitive",
		})
		return aSec, false, nil
	}

	_, err := kubeclientset.UpdateWithCtrlClient(ctx, h.Client, eSec,
		kubeclientset.WithUpdateAlign(alignFn))
	return err
}

func convertSecretListOptsFromVariableListOpts(in ctrlcli.ListOptions) (out *ctrlcli.ListOptions) {
	// Lock field selector.
	in.FieldSelector = fields.SelectorFromSet(fields.Set{
		"metadata.name": systemkuberes.VariablesDelegatedSecretName,
	})

	// Add necessary label selector.
	if lbs := systemmeta.GetResourcesLabelSelectorOfType("variables"); in.LabelSelector == nil {
		in.LabelSelector = lbs
	} else {
		reqs, _ := lbs.Requirements()
		in.LabelSelector = in.LabelSelector.DeepCopySelector().Add(reqs...)
	}

	return &in
}

func convertVariableFromSecret(sec *core.Secret, name string) *walrus.Variable {
	resType, notes := systemmeta.DescribeResource(sec)
	if resType != "variables" {
		return nil
	}

	// Filter out.
	if _, ok := sec.Data[name]; !ok {
		return nil
	}

	uid := sec.UID
	if uidS := notes[name+"-uid"]; len(uidS) != 0 {
		uid = types.UID(uidS)
	}
	createAt := sec.CreationTimestamp
	if createS := notes[name+"-create-at"]; len(createS) != 0 {
		if createAt_, err := time.Parse(time.RFC3339, createS); err == nil {
			createAt = meta.NewTime(createAt_)
		}
	}
	sensitive := notes[name+"-sensitive"] == "true"
	var (
		value  = []byte("")
		value_ = sec.Data[name]
	)
	if len(value_) != 0 && sensitive {
		value = []byte("(sensitive)")
	} else if len(value_) != 0 {
		value = value_
	}

	vra := &walrus.Variable{
		ObjectMeta: meta.ObjectMeta{
			Namespace:         sec.Namespace,
			Name:              name,
			UID:               uid,
			ResourceVersion:   sec.ResourceVersion,
			CreationTimestamp: createAt,
			DeletionTimestamp: sec.DeletionTimestamp,
		},
		Spec: walrus.VariableSpec{
			Sensitive: sensitive,
		},
		Status: walrus.VariableStatus{
			Scope:  walrus.VariableScope(notes["scope"]),
			Value:  string(value),
			Value_: string(value_),
		},
	}

	kubemeta.ConfigureLastAppliedAnnotation(vra)
	return vra
}

func convertVariableListFromSecret(sec *core.Secret, opts ctrlcli.ListOptions) *walrus.VariableList {
	resType, notes := systemmeta.DescribeResource(sec)
	if resType != "variables" {
		return &walrus.VariableList{}
	}

	vList := &walrus.VariableList{
		Items: make([]walrus.Variable, 0, len(sec.Data)),
	}

	for _, name := range sets.KeySet(sec.Data).UnsortedList() {
		// Ignore if not be selected by `kubectl get --field-selector=metadata.namespace=...,metadata.name=...`.
		if fs := opts.FieldSelector; fs != nil &&
			!fs.Matches(fields.Set{"metadata.namespace": sec.Namespace, "metadata.name": name}) {
			continue
		}

		uid := sec.UID
		if uidS := notes[name+"-uid"]; len(uidS) != 0 {
			uid = types.UID(uidS)
		}
		createAt := sec.CreationTimestamp
		if createS := notes[name+"-create-at"]; len(createS) != 0 {
			if createAt_, err := time.Parse(time.RFC3339, createS); err == nil {
				createAt = meta.NewTime(createAt_)
			}
		}
		sensitive := notes[name+"-sensitive"] == "true"
		var (
			value  = []byte("")
			value_ = sec.Data[name]
		)
		if len(value_) != 0 && sensitive {
			value = []byte("(sensitive)")
		} else if len(value_) != 0 {
			value = value_
		}

		vra := &walrus.Variable{
			ObjectMeta: meta.ObjectMeta{
				Namespace:         sec.Namespace,
				Name:              name,
				UID:               uid,
				ResourceVersion:   sec.ResourceVersion,
				CreationTimestamp: createAt,
				DeletionTimestamp: sec.DeletionTimestamp,
			},
			Spec: walrus.VariableSpec{
				Sensitive: sensitive,
			},
			Status: walrus.VariableStatus{
				Scope:  walrus.VariableScope(notes["scope"]),
				Value:  string(value),
				Value_: string(value_),
			},
		}

		kubemeta.OverwriteLastAppliedAnnotation(vra)
		vList.Items = append(vList.Items, *vra)
	}

	return vList
}

func mergeVariableListFromSecretList(secList *core.SecretList, opts ctrlcli.ListOptions) *walrus.VariableList {
	var vList *walrus.VariableList
	{
		vListCount := 0
		for i := range secList.Items {
			vListCount += len(secList.Items[i].Data)
		}
		vList = &walrus.VariableList{
			Items: make([]walrus.Variable, 0, vListCount),
		}
	}

	names := sets.New[string]()
	for i := range secList.Items {
		svList := convertVariableListFromSecret(&secList.Items[i], opts)
		if svList == nil {
			continue
		}
		for j := range svList.Items {
			if names.Has(svList.Items[j].Name) {
				continue
			}
			names.Insert(svList.Items[j].Name)
			vList.Items = append(vList.Items, svList.Items[j])
		}
	}

	// Sort by scope and name.
	sort.SliceStable(vList.Items, func(i, j int) bool {
		l, r := vList.Items[i].Status.Scope, vList.Items[j].Status.Scope
		if l != r {
			return l.Priority() < r.Priority()
		}
		return vList.Items[i].Name < vList.Items[j].Name
	})

	return vList
}

func convertVariableListFromSecretList(secList *core.SecretList, opts ctrlcli.ListOptions) *walrus.VariableList {
	// Sort by resource version.
	sort.SliceStable(secList.Items, func(i, j int) bool {
		l, r := secList.Items[i].ResourceVersion, secList.Items[j].ResourceVersion
		return len(l) < len(r) ||
			(len(l) == len(r) && l < r)
	})

	var vList *walrus.VariableList
	{
		vListCount := 0
		for i := range secList.Items {
			vListCount += len(secList.Items[i].Data)
		}
		vList = &walrus.VariableList{
			Items: make([]walrus.Variable, 0, vListCount),
		}
	}

	for i := range secList.Items {
		svList := convertVariableListFromSecret(&secList.Items[i], opts)
		if svList == nil {
			continue
		}
		vList.Items = append(vList.Items, svList.Items...)
	}

	return vList
}

func getHighestPriorityVariableNamespace(nsSet sets.Set[string]) string {
	nss := nsSet.UnsortedList()
	switch len(nss) {
	case 0:
		return ""
	case 1:
		return nss[0]
	}
	sort.Sort(sort.Reverse(sort.StringSlice(nss)))
	if nss[0] == systemkuberes.SystemNamespaceName {
		return nss[1]
	}
	return nss[0]
}
