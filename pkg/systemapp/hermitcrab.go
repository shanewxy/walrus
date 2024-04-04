package systemapp

import (
	"context"
	"fmt"
	"path/filepath"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemapp/helm"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

func installHermitCrab(ctx context.Context, cli *helm.Client, globalValuesContext map[string]any, disable sets.Set[string]) error {
	// NB: please update the following files if changed.
	// - hack/mirror/walrus-images.txt.
	// - pack/walrus/image/Dockerfile.
	// - github.com/seal-io/helm-charts/charts/walrus.

	name := "hermitcrab"
	version := "0.1.4"
	if disable.Has(name) {
		return nil
	}

	namespace := cli.Namespace()
	release := "walrus-terraform-mirror"
	file := filepath.Join(system.SubLibDir("charts"), fmt.Sprintf("%s.tgz", name))
	download := fmt.Sprintf("https://github.com/seal-io/helm-charts/releases/download/%[1]s-%[2]s/%[1]s-%[2]s.tgz", name, version)
	valuesTemplate := `
global:
  imageRegistry: "{{ .ImageRegistry }}"

fullnameOverride: "{{ .Release }}"
namespaceOverride: "{{ .Namespace }}"

commonAnnotations: 
  {{ .ManagedLabel }}: "true"

hermitcrab:
  image:
    repository: "sealio/hermitcrab"
    tag: "v0.1.4"
{{ if .Env }}
  env: {{ toYaml .Env | nindent 2 }}
{{ end }}
`
	valuesContext := globalValuesContext
	valuesContext["Release"] = release
	valuesContext["Namespace"] = namespace

	chart := &helm.Chart{
		Name:            name,
		Version:         version,
		Release:         release,
		File:            file,
		FileDownloadURL: download,
		Values: helm.YamlTemplateChartValues{
			Template: valuesTemplate,
			Context:  valuesContext,
		},
	}
	_, err := cli.Install(ctx, chart)
	if err != nil {
		return err
	}

	host := fmt.Sprintf("%s-hermitcrab.%s", release, namespace)

	endpoint := fmt.Sprintf("https://%s/v1/providers/", host)
	return systemsetting.TerraformDeployerNetworkMirrorUrl.Configure(ctx, endpoint)
}
