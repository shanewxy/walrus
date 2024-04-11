package driver

import (
	"fmt"
	"net/url"

	"github.com/drone/go-scm/scm"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/vcs/driver/gitee"
	"github.com/seal-io/walrus/pkg/vcs/driver/github"
	"github.com/seal-io/walrus/pkg/vcs/driver/gitlab"
	"github.com/seal-io/walrus/pkg/vcs/options"
)

func NewClientFromURL(platform walruscore.VCSPlatform, rawURL string, opts ...options.ClientOption) (*scm.Client, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	// TODO support reverse proxy for self-hosted server.
	server := u.Scheme + "://" + u.Host

	switch platform {
	case walruscore.VCSPlatformGitHub:
		return github.NewClientFromURL(server, opts...)
	case walruscore.VCSPlatformGitLab:
		return gitlab.NewClientFromURL(server, opts...)
	case walruscore.VCSPlatformGitee:
		return gitee.NewClientFromURL(server, opts...)
	}

	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unsupported VCS platform %q", platform)
}
