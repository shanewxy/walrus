package vcs

import (
	"fmt"
	"sort"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/hashicorp/go-version"
	"k8s.io/klog/v2"
)

// LookupLocalRepoRef returns a reference from a git repository.
func LookupLocalRepoRef(r *git.Repository, name string) (*plumbing.Reference, error) {
	if ref, err := r.Reference(plumbing.NewTagReferenceName(name), true); err == nil {
		return ref, nil
	}

	if ref, err := r.Reference(plumbing.NewBranchReferenceName(name), true); err == nil {
		return ref, nil
	}

	if ref, err := r.Reference(plumbing.NewRemoteReferenceName("origin", name), true); err == nil {
		return ref, nil
	}

	if ref, err := r.Reference(plumbing.NewNoteReferenceName(name), true); err == nil {
		return ref, nil
	}

	if revision, err := r.ResolveRevision(plumbing.Revision(name)); err == nil {
		return plumbing.NewHashReference(plumbing.ReferenceName(name), *revision), nil
	}

	return nil, fmt.Errorf("failed to get reference: %s", name)
}

// ListLocalRepoVersions returns all versions of a git repository in descending order.
func ListLocalRepoVersions(r *git.Repository) ([]string, error) {
	tagRefs, err := r.Tags()
	if err != nil {
		return nil, err
	}

	var versions []*version.Version

	err = tagRefs.ForEach(func(ref *plumbing.Reference) error {
		v, verr := version.NewVersion(ref.Name().Short())
		if verr != nil {
			klog.Warningf("failed to parse tag %s: %v", ref.Name().Short(), verr)
		}

		if v != nil {
			versions = append(versions, v)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].LessThan(versions[j])
	})

	versionStrings := make([]string, len(versions))
	for i, v := range versions {
		versionStrings[i] = v.Original()
	}
	return versionStrings, nil
}
