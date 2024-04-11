package vcssource

import (
	"strings"
)

const (
	// ConnectorTypeLabel indicates the type of the connector.
	ConnectorTypeLabel string = "walrus.seal.io/connector.type"
	// ResourceDefinitionTypeLabel indicates the type of the resource definition.
	ResourceDefinitionTypeLabel string = "walrus.seal.io/resource-definition.type"
	// CatalogNameLabel indicates the name of the catalog.
	CatalogNameLabel string = "walrus.seal.io/catalog.name"
)

// GenWalrusBuiltinLabels generates the builtin labels from repository topics.
func GenWalrusBuiltinLabels(topics []string, catalog string) map[string]string {
	labels := map[string]string{
		CatalogNameLabel: catalog,
	}

	for _, topic := range topics {
		switch {
		case strings.HasPrefix(topic, "c-"):
			labels[ConnectorTypeLabel] = strings.TrimPrefix(topic, "c-")
		case strings.HasPrefix(topic, "t-"):
			labels[ResourceDefinitionTypeLabel] = strings.TrimPrefix(topic, "t-")
		}
	}

	return labels
}
