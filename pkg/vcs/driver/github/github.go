package github

import (
	"net/url"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"

	"github.com/seal-io/walrus/pkg/vcs/options"
)

// NewClientFromURL creates a new github client from url.
func NewClientFromURL(rawURL string, opts ...options.ClientOption) (*scm.Client, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	var client *scm.Client

	switch u.Host {
	case "github.com":
		client = github.NewDefault()

	default:
		client, err = github.New(rawURL)
		if err != nil {
			return nil, err
		}
	}

	options.SetClientOptions(client, opts...)

	return client, nil
}
