package lister

import (
	"context"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/templates/listers/vcssource"
)

type Lister interface {
	List(ctx context.Context, catalog *walruscore.Catalog) ([]walruscore.Template, error)
}

func List(ctx context.Context, catalog *walruscore.Catalog) ([]walruscore.Template, error) {
	switch {
	default:
		panic("unsupported catalog lister type")
	case catalog.Spec.VCSSource != nil:
		l := vcssource.New()
		return l.List(ctx, catalog)
	}
}
