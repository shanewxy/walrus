package vcsrepo

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"k8s.io/klog/v2"

	"github.com/seal-io/walrus/pkg/templates/sourceurl"
)

// gitRepoIconURL retrieves template icon from a git repository and return icon URL.
func gitRepoIconURL(r *git.Repository, url string) (string, error) {
	su, err := sourceurl.ParseURLToSourceURL(url)
	if err != nil {
		return "", err
	}

	// Get icon path.
	p, err := gitRepoIconFilePath(r, su.SubPath)
	if err != nil {
		klog.Errorf("failed to get icon url: %v", err)
		return "", err
	}

	u, err := su.FileRawURL(p)
	if err != nil {
		return "", err
	}
	return u, nil
}

// gitRepoIconFileName retrieves template icon from a git repository and return icon path.
func gitRepoIconFilePath(repoLocal *git.Repository, subPath string) (string, error) {
	var (
		err error
		// Valid icon files.
		icons = []string{
			"icon.png",
			"icon.jpg",
			"icon.jpeg",
			"icon.svg",
		}
	)

	w, err := repoLocal.Worktree()
	if err != nil {
		return "", err
	}

	// Get icon URL.
	for _, icon := range icons {
		if subPath != "" {
			icon = filepath.Join(subPath, icon)
		}
		// If icon exists, get icon rawURL.
		if _, err := w.Filesystem.Stat(icon); err == nil {
			return icon, nil
		}
	}

	return "", nil
}
