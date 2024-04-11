package systemkuberes

import (
	"context"
	"fmt"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

const (
	BuiltinCatalogName = "builtin"
)

func InstallBuiltinCatalog(ctx context.Context, cli clientset.Interface, vcsPlatform walruscore.VCSPlatform) error {
	enableBuiltinCatalog, err := systemsetting.EnableBuiltInCatalog.ValueBool(ctx)
	if err != nil {
		return err
	}
	if !enableBuiltinCatalog {
		return nil
	}

	catalog := &walruscore.Catalog{
		ObjectMeta: meta.ObjectMeta{
			Namespace: SystemNamespaceName,
			Name:      BuiltinCatalogName,
		},
		Spec: walruscore.CatalogSpec{
			Builtin:        true,
			TemplateFormat: walruscore.TemplateFormatTerraform,
			Description:    "Walrus Builtin Catalog.",
			VCSSource: &walruscore.VCSSource{
				Platform: walruscore.VCSPlatformGitHub,
				URL:      "https://catalog.seal.io/walrus-catalog",
			},
		},
	}

	switch vcsPlatform {
	case walruscore.VCSPlatformGitHub:
		catalog.Spec.VCSSource.Platform = walruscore.VCSPlatformGitHub
	case walruscore.VCSPlatformGitee:
		catalog.Spec.VCSSource.Platform = walruscore.VCSPlatformGitee
	default:
		return fmt.Errorf("invalid builtin catalog vcs platform: %s", vcsPlatform)
	}

	catalogCli := cli.WalruscoreV1().Catalogs(SystemNamespaceName)

	_, err = kubeclientset.Create(ctx, catalogCli, catalog)
	if err != nil {
		return fmt.Errorf("install builtin catalog: %w", err)
	}

	return nil
}
