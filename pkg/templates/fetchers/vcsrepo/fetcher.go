package vcsrepo

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/seal-io/utils/stringx"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"k8s.io/utils/set"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemmeta"
	"github.com/seal-io/walrus/pkg/systemsetting"
	"github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/kubehelper"
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
	"github.com/seal-io/walrus/pkg/vcs"
)

type Fetcher struct{}

func New() *Fetcher {
	return &Fetcher{}
}

// Fetch the template info and fill in to template status it.
func (l *Fetcher) Fetch(ctx context.Context, obj *walruscore.Template) (*walruscore.Template, error) {
	tempDir := filepath.Join(os.TempDir(), "seal-template-"+stringx.RandomHex(10))
	defer os.RemoveAll(tempDir)

	source := obj.Spec.VCSRepository

	// Clone.
	{
		tlsVerify, err := systemsetting.EnableRemoteTlsVerify.ValueBool(ctx)
		if err != nil {
			return nil, err
		}

		opts := vcs.GitCloneOptions{
			URL:             source.URL,
			InsecureSkipTLS: !tlsVerify,
		}

		cloneCtx, cancel := context.WithTimeout(ctx, 10*time.Minute)
		defer cancel()

		_, err = vcs.GitClone(cloneCtx, tempDir, opts)
		if err != nil {
			return nil, err
		}
	}

	r, err := git.PlainOpen(tempDir)
	if err != nil {
		return nil, err
	}

	// Icon.
	iconURL, err := gitRepoIconURL(r, source.URL)
	if err != nil {
		return nil, err
	}

	// Versions.
	versions, versionSchema, err := getVersions(obj, r)
	if err != nil {
		return nil, err
	}

	// Skip template without valid version.
	if len(versions) == 0 {
		return nil, nil
	}

	// Ensure template before create schema.
	obj, err = getOrCreateTemplate(ctx, obj)
	if err != nil {
		return nil, err
	}

	templateVersions, err := genTemplateVersions(ctx, obj, versions, versionSchema)
	if err != nil {
		return nil, err
	}

	// Set template status.
	apistatus.TemplateConditionReady.True(obj, "", "")
	obj.Status.Project = systemmeta.GetProjectName(obj.Namespace)
	obj.Status.LastSyncTime = meta.Now()
	obj.Status.URL = source.URL
	obj.Status.Icon = iconURL
	obj.Status.Versions = templateVersions

	err = updateTemplateStatus(ctx, obj)
	return obj, err
}

func updateTemplateStatus(ctx context.Context, obj *walruscore.Template) error {
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	existed, err := loopbackKubeClient.WalruscoreV1().Templates(obj.Namespace).Get(ctx, obj.Name, meta.GetOptions{})
	if err != nil {
		return err
	}

	newSet := set.Set[string]{}
	for _, v := range obj.Status.Versions {
		newSet.Insert(v.Version)
	}

	// Set of versions are removed.
	var removed []walruscore.TemplateVersion
	for i, v := range existed.Status.Versions {
		if !newSet.Has(v.Version) {
			existed.Status.Versions[i].Removed = true
			removed = append(removed, existed.Status.Versions[i])
		}
	}

	// Update template status.
	existed.Status = obj.Status
	existed.Status.ConditionSummary = *apistatus.WalkTemplate(&existed.Status.StatusDescriptor)
	existed.Status.Versions = append(existed.Status.Versions, removed...)
	_, err = loopbackKubeClient.WalruscoreV1().Templates(existed.Namespace).UpdateStatus(ctx, existed, meta.UpdateOptions{})
	return err
}

// getOrCreateTemplate get or create template.
func getOrCreateTemplate(ctx context.Context, obj *walruscore.Template) (*walruscore.Template, error) {
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	existed, err := loopbackKubeClient.WalruscoreV1().Templates(obj.Namespace).Get(ctx, obj.Name, meta.GetOptions{})
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return nil, err
		}

		// Create template.
		existed, err = loopbackKubeClient.WalruscoreV1().Templates(obj.Namespace).Create(ctx, obj, meta.CreateOptions{})
		if err != nil && !kerrors.IsAlreadyExists(err) {
			return nil, err
		}

		return existed, nil
	}
	return existed, nil
}

// genTemplateVersionsFromGitRepo retrieves template versions from a git repository.
func genTemplateVersions(
	ctx context.Context,
	obj *walruscore.Template,
	versions []string,
	versionSchema map[string]*api.SchemaGroup,
) ([]walruscore.TemplateVersion, error) {
	var (
		logger = klog.NewStandardLogger("WARNING")
		tvs    = make([]walruscore.TemplateVersion, 0, len(versionSchema))
	)

	su, err := sourceurl.ParseURLToSourceURL(obj.Spec.VCSRepository.URL)
	if err != nil {
		return nil, err
	}

	for i := range versions {
		version := versions[i]
		schema, ok := versionSchema[version]
		if !ok {
			logger.Printf("%s/%s version: %s version schema not found", obj.Namespace, obj.Name, version)
			continue
		}

		var u string
		{
			link, err := url.Parse(su.Link)
			if err != nil {
				return nil, err
			}
			link.RawQuery = url.Values{"ref": []string{version}}.Encode()
			u = link.String()
		}

		// Generate template version.
		tv, err := kubehelper.GenTemplateVersion(ctx, u, version, obj, schema)
		if err != nil {
			return nil, err
		}

		tvs = append(tvs, *tv)
	}

	return tvs, nil
}
