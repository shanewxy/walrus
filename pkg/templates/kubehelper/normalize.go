package kubehelper

import (
	"strings"

	"github.com/seal-io/utils/stringx"
)

func NormalizeTemplateName(catalogName, repoName string) string {
	return stringx.Join("-", catalogName, strings.TrimPrefix(repoName, "terraform-"))
}

func NormalizeTemplateVersionSchemaName(templateName, version, suffix string) string {
	return stringx.Join("-", templateName, version, suffix)
}
