package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Catalog is the schema for the catalogs API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type Catalog struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   CatalogSpec   `json:"spec"`
	Status CatalogStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Catalog)(nil)

// CatalogSpec defines the desired state of Catalog.
type CatalogSpec struct {
	// Builtin indicate the catalog is builtin catalog.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Builtin bool `json:"builtin,omitempty"`

	// TemplateFormat of the catalog.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	TemplateFormat string `json:"templateFormat"`

	// Description of the catalog.
	Description string `json:"description,omitempty"`

	// Filters specifies the filtering rules for the catalog.
	Filters *Filters `json:"filters,omitempty"`

	// VCSSource specifies the vcs source configure, should update to optional after we support more storage source.
	VCSSource *VCSSource `json:"vcsSource"`
}

// CatalogStatus defines the observed state of Catalog.
type CatalogStatus struct {
	// StatusDescriptor defines the status of the catalog.
	StatusDescriptor `json:",inline"`

	// LastSyncTime record the last sync catalog time.
	LastSyncTime meta.Time `json:"lastSyncTime,omitempty"`

	// TemplateCount is the count of templates.
	TemplateCount int64 `json:"templateCount,omitempty"`

	//  URL of the catalog.
	URL string `json:"url,omitempty"`

	// Project is the project to which the catalog belongs.
	Project string `json:"project,omitempty"`
}

// CatalogList holds the list of Catalog.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CatalogList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Catalog `json:"items"`
}

var _ runtime.Object = (*CatalogList)(nil)

// VCSPlatform is the platform of the version control source.
//
// +enum
type VCSPlatform string

const (
	VCSPlatformGitHub VCSPlatform = "GitHub"
	VCSPlatformGitLab VCSPlatform = "GitLab"
	VCSPlatformGitee  VCSPlatform = "Gitee"
)

const (
	TemplateFormatTerraform string = "Terraform"
)

// Filters specifies the filtering rules for filtering templates imported from the catalog.
type Filters struct {
	// IncludeExpression specifies the regular expression used to match the names of included templates.
	IncludeExpression string `json:"includeExpression,omitempty"`

	// ExcludeExpression specifies the regular expression used to match the names of excluded templates.
	ExcludeExpression string `json:"excludeExpression,omitempty"`
}

// VCSSource specifies the version control source configure.
type VCSSource struct {
	// Platform of the source.
	//
	// +k8s:validation:enum=["GitHub","GitLab","Gitee"]
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Platform VCSPlatform `json:"platform"`

	// URL of the source address, a valid URL contains at least a protocol and host.
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	URL string `json:"url"`
}
