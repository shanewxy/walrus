package system

import (
	"github.com/seal-io/utils/varx"
	"k8s.io/apimachinery/pkg/util/sets"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

var (
	// BootstrapPassword is the password for bootstrapping the system.
	BootstrapPassword = varx.NewOnce("")

	// DisableAuths is a flag to disable authentication.
	DisableAuths = varx.NewOnce(false)

	// DisableApplications is a set of applications that are not allowed to be installed.
	DisableApplications = varx.NewOnce(sets.New[string]())

	// BuiltinCatalogVCSPlatform is the builtin VCS platform for the catalog.
	BuiltinCatalogVCSPlatform = varx.NewOnce(walruscore.VCSPlatformGitHub)
)

// ConfigureControl configures the function of the system.
func ConfigureControl(
	bootstrapPassword string,
	disableAuths bool,
	disableApps []string,
	builtinCatalogVCSPlatform string,
) {
	BootstrapPassword.Configure(bootstrapPassword)
	DisableAuths.Configure(disableAuths)
	DisableApplications.Configure(sets.New[string](disableApps...))
	BuiltinCatalogVCSPlatform.Configure(walruscore.VCSPlatform(builtinCatalogVCSPlatform))
}
