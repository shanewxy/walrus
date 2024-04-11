package vcs

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/drone/go-scm/scm"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/vcs/driver"
	"github.com/seal-io/walrus/pkg/vcs/options"
)

// GetOrgRepos returns full repositories list from the given org.
func GetOrgRepos(ctx context.Context, platform walruscore.VCSPlatform, url string, opts ...options.ClientOption) ([]*scm.Repository, error) {
	client, err := driver.NewClientFromURL(platform, url, opts...)
	if err != nil {
		return nil, err
	}

	orgName, err := getOrgNameFromURL(url)
	if err != nil {
		return nil, err
	}

	var (
		list     []*scm.Repository
		listOpts = scm.ListOptions{Size: 100}
	)

	for {
		repos, meta, err := client.Organizations.ListRepositories(ctx, orgName, listOpts)
		if err != nil {
			return nil, err
		}

		for _, src := range repos {
			if src != nil && !src.Archived {
				list = append(list, src)
			}
		}

		listOpts.Page = meta.Page.Next
		listOpts.URL = meta.Page.NextURL

		if listOpts.Page == 0 && listOpts.URL == "" {
			break
		}
	}

	return list, nil
}

// getOrgNameFromURL parses the organization name from the given URL.
func getOrgNameFromURL(source string) (string, error) {
	u, err := url.Parse(source)
	if err != nil {
		return "", err
	}

	parts := strings.Split(u.Path, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid git url")
	}

	return strings.TrimPrefix(u.Path, "/"), nil
}
