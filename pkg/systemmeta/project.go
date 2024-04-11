package systemmeta

import (
	"github.com/seal-io/walrus/pkg/system"
)

func GetProjectName(namespace string) string {
	switch namespace {
	default:
		return namespace
	case system.NamespaceName:
		return ""
	}
}
