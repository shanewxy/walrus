package gitee

import (
	"fmt"
	"net/url"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/gitee"

	"github.com/seal-io/walrus/pkg/vcs/options"
)

// NewClientFromURL creates a new gitee client from url.
func NewClientFromURL(rawURL string, opts ...options.ClientOption) (*scm.Client, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	var client *scm.Client

	switch u.Host {
	case "gitee.com":
		client = gitee.NewDefault()
	default:
		client, err = gitee.New(fmt.Sprintf("%s/api/v5", u.Scheme+"://"+u.Host))
		if err != nil {
			return nil, err
		}
	}

	options.SetClientOptions(client, opts...)

	return client, nil
}
