package openapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/golang/groupcache/singleflight"
	"github.com/gorilla/mux"
	"github.com/seal-io/utils/httpx"
	"github.com/seal-io/utils/json"
	"github.com/seal-io/utils/stringx"
	"k8s.io/apimachinery/pkg/util/sets"
	openspec3 "k8s.io/kube-openapi/pkg/spec3"
	openvalidatespec "k8s.io/kube-openapi/pkg/validation/spec"
)

type Decorator func(spec *openspec3.OpenAPI) *openspec3.OpenAPI

func Route(r *mux.Route, d ...Decorator) {
	p, _ := r.GetPathTemplate()
	r.Handler(http.StripPrefix(p, index(d...)))
}

func index(d ...Decorator) http.Handler {
	in := interceptor{
		d: d,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := in.Get(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(bs)
	})
}

type interceptor struct {
	d []Decorator
	v atomic.Value
	g singleflight.Group
}

func (l *interceptor) Get(ctx context.Context) ([]byte, error) {
	if v := l.v.Load(); v != nil {
		return v.([]byte), nil
	}

	r, err := l.g.Do("", func() (any, error) {
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		req, err := httpx.NewGetRequestWithContext(ctx, "https://localhost/openapi/v3/apis/walrus.seal.io/v1")
		if err != nil {
			return nil, fmt.Errorf("new request: %w", err)
		}

		resp, err := httpx.DefaultInsecureClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("do request: %w", err)
		}
		defer httpx.Close(resp)

		s := new(openspec3.OpenAPI)
		err = json.NewDecoder(resp.Body).Decode(s)
		if err != nil {
			return nil, fmt.Errorf("decode openapi: %w", err)
		}

		s = decorate(s)
		for i := range l.d {
			if l.d[i] == nil {
				continue
			}
			s = l.d[i](s)
		}

		v, err := json.Marshal(s)
		if err != nil {
			return nil, fmt.Errorf("encode openapi: %w", err)
		}

		l.v.Store(v)
		return v, nil
	})
	if err != nil {
		return nil, err
	}

	return r.([]byte), nil
}

func decorate(spec *openspec3.OpenAPI) *openspec3.OpenAPI {
	decorateInfo(spec)
	decoratePaths(spec)
	decorateComponents(spec)

	return spec
}

func decorateInfo(spec *openspec3.OpenAPI) {
	if spec.Info == nil {
		spec.Info = &openvalidatespec.Info{}
	}

	spec.Info.Description = "Restful APIs to access Walrus."
}

func decoratePaths(spec *openspec3.OpenAPI) {
	if spec.Paths == nil {
		spec.Paths = &openspec3.Paths{
			Paths: map[string]*openspec3.Path{},
		}
	}

	pResM := map[string]string{}
	for _, path := range sets.List(sets.KeySet(spec.Paths.Paths)) {
		// Clean legacy paths.
		switch {
		case !strings.Contains(path, "/namespaces/{namespace}/"):
			delete(spec.Paths.Paths, path)
			continue
		case strings.Contains(path, "/watch/"):
			delete(spec.Paths.Paths, path)
			continue
		default:
		}

		pathV := spec.Paths.Paths[path]
		var pRes string
		{
			_, pRes, _ = strings.Cut(path, "/namespaces/{namespace}/")
			pRes, _, _ = strings.Cut(pRes, "/")
		}
		decoratePathOperation(pathV.Get, pResM, pRes)
		decoratePathOperation(pathV.Put, pResM, pRes)
		decoratePathOperation(pathV.Post, pResM, pRes)
		decoratePathOperation(pathV.Delete, pResM, pRes)
		decoratePathOperation(pathV.Options, pResM, pRes)
		decoratePathOperation(pathV.Head, pResM, pRes)
		decoratePathOperation(pathV.Patch, pResM, pRes)
		decoratePathOperation(pathV.Trace, pResM, pRes)
		decoratePathParameters(pathV.Parameters)
	}
}

func decoratePathOperation(spec *openspec3.Operation, pResM map[string]string, pRes string) {
	if spec == nil {
		return
	}

	// Secure operations.
	spec.SecurityRequirement = []map[string][]string{
		{"BearerAuth": {}},
	}

	// Replace request body.
	if spec.RequestBody != nil && spec.RequestBody.Content != nil {
		if v, ok := spec.RequestBody.Content["*/*"]; ok {
			spec.RequestBody.Content["application/json"] = v
			delete(spec.RequestBody.Content, "*/*")
		}
	}

	// Parameters.
	decoratePathParameters(spec.Parameters)

	// Retag operations.
	if kind := pResM[pRes]; kind != "" {
		spec.Tags = []string{kind}
		return
	}
	{
		var gvk _GroupVersionKind
		if spec.Extensions["x-kubernetes-group-version-kind"] != nil {
			_ = spec.Extensions.GetObject("x-kubernetes-group-version-kind", &gvk)
			if gvk.Validate() {
				kind := gvk.Kind
				res := strings.ToLower(stringx.Pluralize(kind))
				if res == pRes {
					pResM[pRes] = kind
					spec.Tags = []string{kind}
					return
				}
			}
		}
	}
	spec.Tags = []string{stringx.Capitalize(pRes)}
}

func decoratePathParameters(params []*openspec3.Parameter) {
	for i := range params {
		if params[i] == nil || params[i].Schema == nil {
			continue
		}
		param := params[i]
		switch param.Name {
		default:
			continue
		case "fieldManager":
			param.Schema.Default = "walrus-swagger-ui"
		case "fieldValidation":
			param.Schema.Default = "Ignore"
			param.Schema.Enum = []any{
				"Ignore",
				"Warn",
				"Strict",
			}
		case "propagationPolicy":
			param.Schema.Default = "Background"
			param.Schema.Enum = []any{
				"Background",
				"Foreground",
				"Orphan",
			}
		}
	}
}

func decorateComponents(spec *openspec3.OpenAPI) {
	if spec.Components == nil {
		spec.Components = &openspec3.Components{}
	}

	decorateComponentsSchemas(spec.Components)
	decorateComponentsSecuritySchemes(spec.Components)
}

func decorateComponentsSchemas(spec *openspec3.Components) {
	if spec.Schemas == nil {
		spec.Schemas = map[string]*openvalidatespec.Schema{}
	}

	for loc := range spec.Schemas {
		if !strings.HasPrefix(loc, "com.github.seal-io.walrus.") {
			continue
		}
		spec := spec.Schemas[loc]
		var gvkl _GroupVersionKindList
		{
			if spec.Extensions["x-kubernetes-group-version-kind"] == nil {
				continue
			}
			_ = spec.Extensions.GetObject("x-kubernetes-group-version-kind", &gvkl)
			if !gvkl.Validate() {
				continue
			}
		}
		if s, ok := spec.Properties["apiVersion"]; ok {
			s.Default = gvkl[0].APIVersion()
			spec.Properties["apiVersion"] = s
		}
		if s, ok := spec.Properties["kind"]; ok {
			s.Default = gvkl[0].Kind
			spec.Properties["kind"] = s
		}
	}
}

func decorateComponentsSecuritySchemes(spec *openspec3.Components) {
	if spec.SecuritySchemes == nil {
		spec.SecuritySchemes = map[string]*openspec3.SecurityScheme{}
	}

	spec.SecuritySchemes["BearerAuth"] = &openspec3.SecurityScheme{
		SecuritySchemeProps: openspec3.SecuritySchemeProps{
			Type:        "http",
			In:          "header",
			Scheme:      "bearer",
			Description: "Bearer Authentication, the token must be a valid Walrus token.",
		},
	}
}

type _GroupVersionKind struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

func (gvk _GroupVersionKind) APIVersion() string {
	if gvk.Group == "" || gvk.Group == "core" {
		return gvk.Version
	}
	return gvk.Group + "/" + gvk.Version
}

func (gvk _GroupVersionKind) Validate() bool {
	return gvk.Group != "" && gvk.Version != "" && gvk.Kind != ""
}

type _GroupVersionKindList []_GroupVersionKind

func (gvkl _GroupVersionKindList) Validate() bool {
	return len(gvkl) == 1 && gvkl[0].Validate()
}
