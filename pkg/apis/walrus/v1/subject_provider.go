package v1

import (
	"errors"
	"fmt"
	"strings"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// SubjectProvider is the schema for the subject providers API.
//
// +genclient
// +genclient:onlyVerbs=create,get,list,watch,apply,update,patch,delete,deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],shortName=["subjprov"]
type SubjectProvider struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubjectProviderSpec   `json:"spec"`
	Status SubjectProviderStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*SubjectProvider)(nil)

// SubjectProviderType describes the type of subject provider.
// +enum
type SubjectProviderType string

const (
	// SubjectProviderTypeInternal means the subject provider is builtin.
	//
	// It is only support using username/password to authenticate.
	SubjectProviderTypeInternal SubjectProviderType = "Internal"

	// SubjectProviderTypeLDAP means the subject provider is based on LDAP.
	//
	// This type is an external subject provider type,
	// which is able to use standards-compliant LDAP as an authentication source.
	SubjectProviderTypeLDAP SubjectProviderType = "LDAP"
	// SubjectProviderTypeOAuth means the subject provider is based on OAuth 2.0.
	//
	// This type is an external subject provider type,
	// which is able to use standards-compliant OAuth 2.0 provider as an authentication source.
	SubjectProviderTypeOAuth SubjectProviderType = "OAuth"
	// SubjectProviderTypeOIDC means the subject provider is based on OpenID Connect.
	//
	// This type is an external subject provider type,
	// which is able to use standards-compliant OpenID Connect as an authentication source.
	SubjectProviderTypeOIDC SubjectProviderType = "OIDC"

	// SubjectProviderTypeGithub means the subject provider is based on GitHub.
	//
	// This type is an external subject provider type,
	// which is able to use GitHub as an authentication source.
	SubjectProviderTypeGithub SubjectProviderType = "GitHub"
	// SubjectProviderTypeGitlab means the subject provider is based on GitLab.
	//
	// This type is an external subject provider type,
	// which is able to use GitLab as an authentication source.
	SubjectProviderTypeGitlab SubjectProviderType = "GitLab"
	// SubjectProviderTypeBitbucket means the subject provider is based on Bitbucket.
	//
	// This type is an external subject provider type,
	// which is able to use Bitbucket as an authentication source.
	SubjectProviderTypeBitbucket SubjectProviderType = "Bitbucket"
	// SubjectProviderTypeGitea means the subject provider is based on Gitea.
	//
	// This type is an external subject provider type,
	// which is able to use Gitea as an authentication source.
	SubjectProviderTypeGitea SubjectProviderType = "Gitea"
	// SubjectProviderTypeGoogle means the subject provider is based on Google.
	//
	// This type is an external subject provider type,
	// which is able to use Google as an authentication source.
	SubjectProviderTypeGoogle SubjectProviderType = "Google"
	// SubjectProviderTypeMicrosoft means the subject provider is based on Microsoft.
	//
	// This type is an external subject provider type,
	// which is able to use Microsoft as an authentication source.
	SubjectProviderTypeMicrosoft SubjectProviderType = "Microsoft"
)

func (in SubjectProviderType) String() string {
	return string(in)
}

func (in SubjectProviderType) Validate() error {
	switch in {
	case SubjectProviderTypeInternal, SubjectProviderTypeLDAP,
		SubjectProviderTypeOIDC, SubjectProviderTypeOAuth,
		SubjectProviderTypeGithub, SubjectProviderTypeGitlab,
		SubjectProviderTypeBitbucket, SubjectProviderTypeGitea,
		SubjectProviderTypeGoogle, SubjectProviderTypeMicrosoft:
		return nil
	default:
		return errors.New("invalid subject provider type")
	}
}

// SubjectProviderSpec defines the desired state of SubjectProvider.
type SubjectProviderSpec struct {
	// Type is the type of the subject provider,
	// which is immutable.
	//
	// +k8s:validation:enum=["Internal","LDAP","oidc","oauth","github","gitlab","bitbucket","gitea","google","microsoft"]
	Type SubjectProviderType `json:"type"`

	// DisplayName is the display name of the subject provider.
	DisplayName string `json:"displayName,omitempty"`

	// Description is the description of the subject provider.
	Description string `json:"description,omitempty"`

	// ExternalConfig is the configuration of the external subject provider.
	ExternalConfig SubjectProviderExternalConfig `json:"externalConfig,omitempty"`
}

type (
	// SubjectProviderMicrosoftTenant defines the tenant of the Microsoft.
	// +enum
	SubjectProviderMicrosoftTenant string

	// SubjectProviderLDAPUserSearch defines the user search configuration of the LDAP.
	SubjectProviderLDAPUserSearch struct {
		// BaseDN to start the search from.
		BaseDN string `json:"baseDN"`
		// Filter to apply to the search.
		//
		// +default="(objectClass=person)"
		Filter string `json:"filter,omitempty"`

		// NameAttribute is the attribute to use as the username.
		//
		// +default="uid"
		NameAttribute string `json:"nameAttribute,omitempty"`
		// DisplayNameAttribute is the attribute to use as the display name.
		//
		// +default="cn"
		DisplayNameAttribute string `json:"displayNameAttribute,omitempty"`
		// EmailAttribute is the attribute to use as the email.
		//
		// +default="mail"
		EmailAttribute string `json:"emailAttribute,omitempty"`
	}

	// SubjectProviderLdapGroupSearchUserMatcher defines the user matcher of the LDAP group search.
	SubjectProviderLdapGroupSearchUserMatcher struct {
		// GroupAttribute is the attribute of the group.
		GroupAttribute string `json:"groupAttribute"`
		// UserAttribute is the attribute of the user.
		UserAttribute string `json:"userAttribute"`
	}

	// SubjectProviderLDAPGroupSearch defines the group search configuration of the LDAP.
	SubjectProviderLDAPGroupSearch struct {
		// BaseDN to start the search from.
		BaseDN string `json:"baseDN"`
		// Filter to apply to the search.
		//
		// +default="(objectClass=group)"
		Filter string `json:"filter,omitempty"`

		// UserMatchers is the user matcher list for the LDAP group search.
		UserMatchers []SubjectProviderLdapGroupSearchUserMatcher `json:"userMatchers,omitempty"`

		// NameAttribute is the attribute to use as the group name.
		//
		// +default="name"
		NameAttribute string `json:"nameAttribute,omitempty"`
	}

	// SubjectProviderOAuthClaimMapping defines the claim mapping of the OAuth.
	SubjectProviderOAuthClaimMapping struct {
		// NameKey is the key to pick "name" from claim.
		//
		// +default="preferred_username"
		NameKey string `json:"nameKey,omitempty"`
		// DisplayNameKey is the key to pick "displayName" from claim.
		//
		// +default="name"
		DisplayNameKey string `json:"displayNameKey,omitempty"`
		// EmailKey is the key to pick "email" from claim.
		//
		// +default="email"
		EmailKey string `json:"emailKey,omitempty"`
		// GroupsKey is the key to pick "groups" from claim.
		//
		// +default="groups"
		GroupsKey string `json:"groupsKey,omitempty"`
	}

	// SubjectProviderGitGroupsMatcher defines the groups matcher of the Git.
	SubjectProviderGitGroupsMatcher []string
)

const (
	// SubjectProviderMicrosoftTenantCommon means the tenant is common.
	SubjectProviderMicrosoftTenantCommon SubjectProviderMicrosoftTenant = "common"
	// SubjectProviderMicrosoftTenantOrganizations means the tenant is organizations.
	SubjectProviderMicrosoftTenantOrganizations SubjectProviderMicrosoftTenant = "organizations"
	// SubjectProviderMicrosoftTenantConsumers means the tenant is consumers.
	SubjectProviderMicrosoftTenantConsumers SubjectProviderMicrosoftTenant = "consumers"
)

func (in SubjectProviderMicrosoftTenant) String() string {
	return string(in)
}

func (in SubjectProviderMicrosoftTenant) Validate() error {
	switch in {
	case SubjectProviderMicrosoftTenantCommon,
		SubjectProviderMicrosoftTenantOrganizations,
		SubjectProviderMicrosoftTenantConsumers:
		return nil
	default:
		return errors.New("invalid Microsoft tenant")
	}
}

func (in *SubjectProviderOAuthClaimMapping) Default() {
	if in == nil {
		return
	}

	if in.NameKey == "" {
		in.NameKey = "preferred_username"
	}
	if in.DisplayNameKey == "" {
		in.DisplayNameKey = "name"
	}
	if in.EmailKey == "" {
		in.EmailKey = "email"
	}
	if in.GroupsKey == "" {
		in.GroupsKey = "groups"
	}
}

func (in SubjectProviderGitGroupsMatcher) ToMap() map[string][]string {
	m := make(map[string][]string)
	for _, g := range in {
		s := strings.SplitN(g, ":", 2)
		if len(s) != 2 {
			continue
		}
		m[s[0]] = append(m[s[0]], s[1])
	}
	return m
}

type (
	// SubjectProviderExternalConfig defines the configuration of the subject provider.
	SubjectProviderExternalConfig struct {
		// LDAP is the configuration of the LDAP.
		LDAP *SubjectProviderLDAPConfig `json:"ldap,omitempty"`
		// OAuth is the configuration of the OAuth 2.0.
		OAuth *SubjectProviderOAuthConfig `json:"oauth,omitempty"`
		// OIDC is the configuration of the OpenID Connect.
		OIDC *SubjectProviderOIDCConfig `json:"oidc,omitempty"`
		// GitHub is the configuration of the GitHub.
		GitHub *SubjectProviderGitHubConfig `json:"github,omitempty"`
		// GitLab is the configuration of the GitLab.
		GitLab *SubjectProviderGitLabConfig `json:"gitlab,omitempty"`
		// Bitbucket is the configuration of the Bitbucket.
		Bitbucket *SubjectProviderBitbucketConfig `json:"bitbucket,omitempty"`
		// Gitea is the configuration of the Gitea.
		Gitea *SubjectProviderGiteaConfig `json:"gitea,omitempty"`
		// Google is the configuration of the Google.
		Google *SubjectProviderGoogleConfig `json:"google,omitempty"`
		// Microsoft is the configuration of the Microsoft.
		Microsoft *SubjectProviderMicrosoftConfig `json:"microsoft,omitempty"`
	}

	// SubjectProviderLDAPConfig defines the configuration of the LDAP.
	SubjectProviderLDAPConfig struct {
		// Host is the host of the LDAP server,
		// which in form of "hostname:port".
		//
		// If the port is not specified,
		// the default port 636 will be used.
		Host string `json:"host"`
		// SkipInsecureVerify is the flag to skip the insecure verify.
		SkipInsecureVerify bool `json:"skipInsecureVerify,omitempty"`

		// ServiceAccountDN is the distinguished name of the service account of the LDAP server.
		// It is used to search other users.
		ServiceAccountDN string `json:"serviceAccountDN"`
		// ServiceAccountPassword is the password of the service account of the LDAP server.
		ServiceAccountPassword string `json:"serviceAccountPassword"`

		// GroupSearch is the group search configuration of the LDAP.
		//
		// It maps a group to a list of users.
		GroupSearch SubjectProviderLDAPGroupSearch `json:"groupSearch"`
		// UserSearch is the user search configuration of the LDAP.
		//
		// It maps a username and password entered by a user to the LDAP entry.
		UserSearch SubjectProviderLDAPUserSearch `json:"userSearch"`
	}

	// SubjectProviderOAuthConfig defines the configuration of the OAuth 2.0.
	SubjectProviderOAuthConfig struct {
		// AuthorizationEndpoint is the endpoint to authorize.
		AuthorizationEndpoint string `json:"authorizationEndpoint"`
		// TokenEndpoint is the endpoint to get OAuth token.
		TokenEndpoint string `json:"tokenEndpoint"`
		// UserinfoEndpoint is the endpoint to get the user info.
		UserinfoEndpoint string `json:"userinfoEndpoint"`
		// SkipInsecureVerify is the flag to skip the insecure verify.
		SkipInsecureVerify bool `json:"skipInsecureVerify,omitempty"`

		// ClientID is the ID of the OAuth client.
		ClientID string `json:"clientID"`
		// ClientSecret is the secret of the OAuth client.
		ClientSecret string `json:"clientSecret"`

		// ClaimMapping is the claim mapping of the OpenID Connect.
		ClaimMapping *SubjectProviderOAuthClaimMapping `json:"claimMapping,omitempty"`
	}

	// SubjectProviderOIDCConfig defines the configuration of the OpenID Connect.
	SubjectProviderOIDCConfig struct {
		// Issuer is the URL of the OpenID Connect issuer.
		//
		// This field is only used when the provider is set to "custom".
		//
		// Canonical URL of the provider, also used for configuration discovery.
		// This value MUST match the value returned in the provider config discovery.
		//
		// See: https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfig
		Issuer string `json:"issuer"`
		// SkipInsecureVerify is the flag to skip the insecure verify.
		SkipInsecureVerify bool `json:"skipInsecureVerify,omitempty"`

		// ClientID is the client ID of the OpenID Connect client.
		ClientID string `json:"clientID"`
		// ClientSecret is the client secret of the OpenID Connect client.
		ClientSecret string `json:"clientSecret"`

		// ClaimMapping is the claim mapping of the OpenID Connect.
		ClaimMapping *SubjectProviderOAuthClaimMapping `json:"claimMapping,omitempty"`
	}

	// SubjectProviderGitHubConfig defines the configuration of the GitHub.
	SubjectProviderGitHubConfig struct {
		// ClientID is the client ID of the GitHub client.
		ClientID string `json:"clientID"`
		// ClientSecret is the client secret of the GitHub client.
		ClientSecret string `json:"clientSecret"`

		// Groups are used to filter out which groups should be matched.
		//
		// Each item is in the form of "organization:team".
		Groups SubjectProviderGitGroupsMatcher `json:"groups,omitempty"`
	}

	// SubjectProviderGitLabConfig defines the configuration of the GitLab.
	SubjectProviderGitLabConfig struct {
		// ClientID is the client ID of the GitLab client.
		ClientID string `json:"clientID"`
		// ClientSecret is the client secret of the GitLab client.
		ClientSecret string `json:"clientSecret"`

		// Groups are used to filter out which groups should be matched.
		Groups []string `json:"groups,omitempty"`
	}

	// SubjectProviderBitbucketConfig defines the configuration of the Bitbucket.
	SubjectProviderBitbucketConfig struct {
		// ClientID is the client ID of the Bitbucket client.
		ClientID string `json:"clientID"`
		// ClientSecret is the client secret of the Bitbucket client.
		ClientSecret string `json:"clientSecret"`

		// Groups are used to filter out which groups should be matched.
		Groups []string `json:"groups,omitempty"`
	}

	// SubjectProviderGiteaConfig defines the configuration of the Gitea.
	SubjectProviderGiteaConfig struct {
		// ClientID is the client ID of the Gitea client.
		ClientID string `json:"clientID"`
		// ClientSecret is the client secret of the Gitea client.
		ClientSecret string `json:"clientSecret"`

		// Groups is used to filter out which groups should be matched.
		//
		// Each item is in the form of "organization:team".
		Groups SubjectProviderGitGroupsMatcher `json:"groups,omitempty"`
	}

	// SubjectProviderGoogleConfig defines the configuration of the Google.
	SubjectProviderGoogleConfig struct {
		// ClientID is the client ID of the Google client.
		ClientID string `json:"clientID"`
		// ClientSecret is the client secret of the Google client.
		ClientSecret string `json:"clientSecret"`

		// Groups are used to filter out which groups should be matched.
		Groups []string `json:"groups,omitempty"`
	}

	// SubjectProviderMicrosoftConfig defines the configuration of the Microsoft.
	SubjectProviderMicrosoftConfig struct {
		// ClientID is the client ID of the Microsoft client.
		ClientID string `json:"clientID"`
		// ClientSecret is the client secret of the Microsoft client.
		ClientSecret string `json:"clientSecret"`

		// Tenant is the tenant of the Microsoft client.
		//
		// The tenant is the directory that the user is in.
		//
		// +default="common"
		// +k8s:validation:enum=["common","organizations","consumers"]
		Tenant SubjectProviderMicrosoftTenant `json:"tenant,omitempty"`

		// Groups are used to filter out which groups should be matched.
		Groups []string `json:"groups,omitempty"`
	}
)

func (in *SubjectProviderExternalConfig) ValidateWithType(pt SubjectProviderType) error {
	if in == nil {
		return nil
	}

	switch pt {
	case SubjectProviderTypeLDAP:
		if in.LDAP == nil {
			return errors.New("ldap is required")
		}
		in.LDAP.Default()
		return in.LDAP.Validate()
	case SubjectProviderTypeOAuth:
		if in.OAuth == nil {
			return errors.New("oauth is required")
		}
		in.OAuth.Default()
		return in.OAuth.Validate()
	case SubjectProviderTypeOIDC:
		if in.OIDC == nil {
			return errors.New("oidc is required")
		}
		in.OIDC.Default()
		return in.OIDC.Validate()
	case SubjectProviderTypeGithub:
		if in.GitHub == nil {
			return errors.New("github is required")
		}
		return in.GitHub.Validate()
	case SubjectProviderTypeGitlab:
		if in.GitLab == nil {
			return errors.New("gitlab is required")
		}
		return in.GitLab.Validate()
	case SubjectProviderTypeBitbucket:
		if in.Bitbucket == nil {
			return errors.New("bitbucket is required")
		}
		return in.Bitbucket.Validate()
	case SubjectProviderTypeGitea:
		if in.Gitea == nil {
			return errors.New("gitea is required")
		}
		return in.Gitea.Validate()
	case SubjectProviderTypeGoogle:
		if in.Google == nil {
			return errors.New("google is required")
		}
		return in.Google.Validate()
	case SubjectProviderTypeMicrosoft:
		if in.Microsoft == nil {
			return errors.New("microsoft is required")
		}
		in.Microsoft.Default()
		return in.Microsoft.Validate()
	}
	return nil
}

func (in *SubjectProviderLDAPConfig) Default() {
	if in == nil {
		return
	}

	if in.UserSearch.Filter == "" {
		in.UserSearch.Filter = "(objectClass=person)"
	}
	if in.UserSearch.NameAttribute == "" {
		in.UserSearch.NameAttribute = "uid"
	}
	if in.UserSearch.DisplayNameAttribute == "" {
		in.UserSearch.DisplayNameAttribute = "cn"
	}
	if in.UserSearch.EmailAttribute == "" {
		in.UserSearch.EmailAttribute = "mail"
	}

	if in.GroupSearch.Filter == "" {
		in.GroupSearch.Filter = "(objectClass=group)"
	}
	if in.GroupSearch.NameAttribute == "" {
		in.GroupSearch.NameAttribute = "name"
	}
}

func (in *SubjectProviderLDAPConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.Host == "" {
		return errors.New("host is required")
	}
	if in.ServiceAccountDN == "" {
		return errors.New("serviceAccountDN is required")
	}
	if in.ServiceAccountPassword == "" {
		return errors.New("serviceAccountPassword is required")
	}
	if in.GroupSearch.BaseDN == "" {
		return errors.New("groupSearch.baseDN is required")
	}
	for i, um := range in.GroupSearch.UserMatchers {
		if um.GroupAttribute == "" {
			return fmt.Errorf("groupSearch.userMatchers[%d].groupAttribute is required", i)
		}
		if um.UserAttribute == "" {
			return fmt.Errorf("groupSearch.userMatchers[%d].userAttribute is required", i)
		}
	}
	if in.UserSearch.BaseDN == "" {
		return errors.New("userSearch.baseDN is required")
	}
	return nil
}

func (in *SubjectProviderOAuthConfig) Default() {
	if in == nil {
		return
	}

	if in.ClaimMapping == nil {
		in.ClaimMapping = &SubjectProviderOAuthClaimMapping{}
	}
	in.ClaimMapping.Default()
}

func (in *SubjectProviderOAuthConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.AuthorizationEndpoint == "" {
		return errors.New("authorizationEndpoint is required")
	}
	if in.TokenEndpoint == "" {
		return errors.New("tokenEndpoint is required")
	}
	if in.UserinfoEndpoint == "" {
		return errors.New("userinfoEndpoint is required")
	}
	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	return nil
}

func (in *SubjectProviderOIDCConfig) Default() {
	if in == nil {
		return
	}

	if in.ClaimMapping == nil {
		in.ClaimMapping = &SubjectProviderOAuthClaimMapping{}
	}
	in.ClaimMapping.Default()
}

func (in *SubjectProviderOIDCConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.Issuer == "" {
		return errors.New("issuer is required")
	}
	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	return nil
}

func (in *SubjectProviderGitHubConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	return nil
}

func (in *SubjectProviderGitLabConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	return nil
}

func (in *SubjectProviderBitbucketConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	return nil
}

func (in *SubjectProviderGiteaConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	return nil
}

func (in *SubjectProviderGoogleConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	return nil
}

func (in *SubjectProviderMicrosoftConfig) Default() {
	if in == nil {
		return
	}

	in.Tenant = SubjectProviderMicrosoftTenantCommon
}

func (in *SubjectProviderMicrosoftConfig) Validate() error {
	if in == nil {
		return nil
	}

	if in.ClientID == "" {
		return errors.New("clientID is required")
	}
	if in.ClientSecret == "" {
		return errors.New("clientSecret is required")
	}
	if err := in.Tenant.Validate(); err != nil {
		return err
	}
	return nil
}

// SubjectProviderStatus defines the observed state of SubjectProvider.
type SubjectProviderStatus struct {
	// LoginWithPassword is the flag to indicate whether the provider supports login with password.
	LoginWithPassword bool `json:"loginWithPassword"`
}

// SubjectProviderList holds the list of SubjectProvider.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SubjectProviderList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []SubjectProvider `json:"items"`
}

var _ runtime.Object = (*SubjectProviderList)(nil)
