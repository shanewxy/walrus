package gitlab

import (
	"net/url"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/gitlab"

	"github.com/seal-io/walrus/pkg/vcs/options"
)

// NewClientFromURL creates a new gitlab client from url.
func NewClientFromURL(rawURL string, opts ...options.ClientOption) (*scm.Client, error) {
	_, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	client, err := gitlab.New(rawURL)
	if err != nil {
		return nil, err
	}

	options.SetClientOptions(client, opts...)

	return client, nil
}
