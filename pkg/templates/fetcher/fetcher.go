package fetcher

import (
	"context"
	"fmt"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/templates/fetchers/vcsrepo"
)

type Fetcher interface {
	Fetch(ctx context.Context, template *walruscore.Template) (*walruscore.Template, error)
}

// Fetch sync the template and it's versions.
func Fetch(ctx context.Context, tmpl *walruscore.Template) (*walruscore.Template, error) {
	var e Fetcher
	switch {
	default:
		return nil, fmt.Errorf("unsupport template format %s", tmpl.Spec.TemplateFormat)
	case tmpl.Spec.VCSRepository != nil:
		e = vcsrepo.New()
	}

	return e.Fetch(ctx, tmpl)
}
