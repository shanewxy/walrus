package kubeconfig

import (
	"errors"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// LoadRestConfigNonInteractive loads a rest config according to the following rules.
//
//  1. assume that running as a Pod and try to connect to
//     the Kubernetes cluster with the mounted ServiceAccount.
//  2. load from recommended home file if none of the above conditions are met.
func LoadRestConfigNonInteractive() (cfgPath string, restCfg *rest.Config, inside bool, err error) {
	// Try the in-cluster config.
	restCfg, err = rest.InClusterConfig()
	switch {
	case err == nil:
		return "", restCfg, true, nil
	case !errors.Is(err, rest.ErrNotInCluster):
		return "", nil, false, err
	}

	// Try the recommended config.
	var (
		ld = &clientcmd.ClientConfigLoadingRules{
			Precedence: []string{clientcmd.RecommendedHomeFile},
		}
		od = &clientcmd.ConfigOverrides{}
	)
	restCfg, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(ld, od).ClientConfig()
	return clientcmd.RecommendedHomeFile, restCfg, false, err
}

// LoadClientConfig loads a client config from the specified path,
// the given path must exist.
func LoadClientConfig(path string) (clientcmd.ClientConfig, error) {
	if path == "" {
		return nil, errors.New("blank kubeconfig path")
	}

	var (
		ld = &clientcmd.ClientConfigLoadingRules{
			ExplicitPath: path,
		}
		od = &clientcmd.ConfigOverrides{}
	)

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(ld, od), nil
}

// LoadRestConfig loads a rest config from the specified path,
// the given path must exist.
func LoadRestConfig(path string) (*rest.Config, error) {
	cc, err := LoadClientConfig(path)
	if err != nil {
		return nil, err
	}

	return cc.ClientConfig()
}

// AuthorizeRestConfigWithAuthInfo decorates the given rest config with the given authentication info.
func AuthorizeRestConfigWithAuthInfo(restCfg rest.Config, authInfo clientcmdapi.AuthInfo) *rest.Config {
	restCfg.TLSClientConfig = *restCfg.TLSClientConfig.DeepCopy()

	switch {
	case authInfo.Username != "" && authInfo.Password != "":
		restCfg.Username = authInfo.Username
		restCfg.Password = authInfo.Password
		restCfg.BearerTokenFile = ""
		restCfg.BearerToken = ""
		restCfg.CertFile = ""
		restCfg.CertData = nil
		restCfg.KeyFile = ""
		restCfg.KeyData = nil
	case authInfo.Token != "":
		restCfg.BearerToken = authInfo.Token
		restCfg.BearerTokenFile = ""
		restCfg.Username = ""
		restCfg.Password = ""
		restCfg.CertFile = ""
		restCfg.CertData = nil
		restCfg.KeyFile = ""
		restCfg.KeyData = nil
	}

	if authInfo.Impersonate != "" {
		restCfg.Impersonate = rest.ImpersonationConfig{
			UserName: authInfo.Impersonate,
			UID:      authInfo.ImpersonateUID,
			Groups:   authInfo.ImpersonateGroups,
			Extra:    authInfo.ImpersonateUserExtra,
		}
	}

	return &restCfg
}

// UnauthorizeRestConfig erases the authentication info from the  given rest config.
func UnauthorizeRestConfig(restCfg rest.Config) *rest.Config {
	restCfg.TLSClientConfig = *restCfg.TLSClientConfig.DeepCopy()

	restCfg.Impersonate = rest.ImpersonationConfig{
		UserName: "system:anonymous",
		Groups:   []string{"system:unauthenticated"},
	}

	return &restCfg
}
