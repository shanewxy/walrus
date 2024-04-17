package vcsrepo

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"k8s.io/klog/v2"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/loader"
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
	"github.com/seal-io/walrus/pkg/vcs"
)

// getValidVersions get valid repository versions.
func getVersions(
	t *walruscore.Template,
	r *git.Repository,
) ([]string, map[string]*api.SchemaGroup, error) {
	su, err := sourceurl.ParseURLToSourceURL(t.Spec.VCSRepository.URL)
	if err != nil {
		return nil, nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, nil, err
	}

	var versions []string
	switch {
	default:
		versions, err = vcs.ListLocalRepoVersions(r)
		if err != nil {
			return nil, nil, err
		}
	case su.Reference != "":
		versions = []string{su.Reference}
	}

	validVersions := make([]string, 0, len(versions))
	versionSchema := make(map[string]*api.SchemaGroup)

	for i := range versions {
		tag := versions[i]
		klog.V(5).Infof("getting schema of template \"%s/%s\", version \"%s\"", t.Namespace, t.Name, tag)

		// Checkout tag.
		{
			resetRef, err := vcs.LookupLocalRepoRef(r, tag)
			if err != nil {
				klog.Warningf("failed to get \"%s/%s\" git reference \"%s\": %v",
					t.Namespace, t.Name, tag, err)
				continue
			}

			hash := resetRef.Hash()

			// If tag is not a commit hash, get commit hash from tag object target.
			object, err := r.TagObject(hash)
			if err == nil {
				hash = object.Target
			}

			err = w.Reset(&git.ResetOptions{
				Commit: hash,
				Mode:   git.HardReset,
			})
			if err != nil {
				klog.Warningf("failed to set git reference to %s of template \"%s/%s\" : %v", tag, t.Namespace, t.Name, err)
				continue
			}
		}

		// Load schema.
		{
			dir := w.Filesystem.Root()
			if su.SubPath != "" {
				dir = filepath.Join(dir, su.SubPath)
			}

			sg, err := loader.LoadSchema(dir, t.Name, t.Spec.TemplateFormat)
			if err != nil {
				klog.Warningf("failed to get schema of template \"%s:%s\", version \"%s\": %v", t.Namespace, t.Name, tag, err)
				continue
			}

			validVersions = append(validVersions, tag)
			versionSchema[tag] = sg
		}
	}

	return validVersions, versionSchema, nil
}
