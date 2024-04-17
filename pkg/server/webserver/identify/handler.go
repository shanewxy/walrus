package identify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/seal-io/utils/funcx"
	"github.com/seal-io/utils/httpx"
	"github.com/seal-io/utils/stringx"
	authorization "k8s.io/api/authorization/v1"
	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/utils/ptr"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	applywalrus "github.com/seal-io/walrus/pkg/clients/applyconfiguration/walrus/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubeconfig"
	"github.com/seal-io/walrus/pkg/server/webserver/openapi"
	"github.com/seal-io/walrus/pkg/server/webserver/ui"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemauthz"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

func Route(r *mux.Route) openapi.Decorator {
	p, _ := r.GetPathTemplate()
	sr := r.Subrouter()
	sr.Path("/providers").Methods(http.MethodGet).
		HandlerFunc(providers)
	sr.Path("/login").Methods(http.MethodGet, http.MethodPost).
		HandlerFunc(login)
	sr.Path("/callback/{provider}").Methods(http.MethodGet).
		HandlerFunc(callback)
	sr.Path("/profile").Methods(http.MethodGet, http.MethodPut).
		HandlerFunc(profile)
	sr.Path("/token").Methods(http.MethodGet).
		HandlerFunc(token)
	sr.Path("/rules/{namespace}").Methods(http.MethodGet).
		HandlerFunc(rules)
	sr.Path("/logout").Methods(http.MethodGet).
		HandlerFunc(logout)
	return getOpenapiDecorate(p)
}

type (
	responseProvider struct {
		Name              string `json:"name"`
		Type              string `json:"type"`
		DisplayName       string `json:"displayName"`
		Description       string `json:"description"`
		LoginWithPassword bool   `json:"loginWithPassword"`
	}
	responseProviderList struct {
		Items []responseProvider `json:"items"`
	}
)

// providers is a handler to list all providers.
//
// GET: /providers
func providers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// List.
	subjProvList := new(walrus.SubjectProviderList)
	{
		cli := system.LoopbackCtrlClient.Get()
		err := cli.List(ctx, subjProvList,
			ctrlcli.InNamespace(systemkuberes.SystemNamespaceName))
		if err != nil {
			ui.ResponseErrorWithCode(w, http.StatusInternalServerError, fmt.Errorf("list providers: %w", err))
			return
		}
	}

	// Output.
	resp := responseProviderList{
		Items: make([]responseProvider, 0, len(subjProvList.Items)),
	}
	for i := range subjProvList.Items {
		resp.Items = append(resp.Items, responseProvider{
			Name:              subjProvList.Items[i].Name,
			Type:              subjProvList.Items[i].Spec.Type.String(),
			DisplayName:       subjProvList.Items[i].Spec.DisplayName,
			Description:       subjProvList.Items[i].Spec.Description,
			LoginWithPassword: subjProvList.Items[i].Status.LoginWithPassword,
		})
	}
	httpx.JSON(w, http.StatusOK, resp)
}

type (
	requestLogin struct {
		Provider string `query:"provider"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

// login is a handler to log in.
//
// POST/GET: /login?provider={provider}
func login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse request.
	var req requestLogin
	_ = httpx.BindWith(r, &req, httpx.BindJSON, httpx.BindQuery)

	// Login with username and password.
	if req.Provider == "" || req.Provider == systemkuberes.DefaultSubjectProviderName {
		if r.Method != http.MethodPost {
			ui.ResponseErrorWithCode(w, http.StatusMethodNotAllowed, nil)
			return
		}

		if req.Username == "" {
			ui.ResponseErrorWithCode(w, http.StatusBadRequest, errors.New("username is required"))
			return
		}
		if req.Password == "" {
			ui.ResponseErrorWithCode(w, http.StatusBadRequest, errors.New("password is required"))
			return
		}

		subj := &walrus.Subject{
			ObjectMeta: meta.ObjectMeta{
				Namespace: systemkuberes.SystemNamespaceName,
				Name:      req.Username,
			},
			Spec: walrus.SubjectSpec{
				Credential: ptr.To(req.Password),
			},
		}

		loginSubject(w, r, subj, false)
		return
	}

	// Get provider.
	subjProv, err := getSubjectProvider(ctx, req.Provider)
	if err != nil {
		ui.ResponseErrorWithCode(w, http.StatusInternalServerError, fmt.Errorf("get provider: %w", err))
		return
	}

	// Get connector.
	conn, err := getExternalConnectorFromSubjectProvider(subjProv)
	if err != nil {
		ui.RedirectErrorWithCode(w, r, http.StatusInternalServerError, err)
		return
	}

	switch cn := conn.(type) {
	default:
		ui.ResponseErrorWithCode(w, http.StatusBadRequest, errors.New("unsupported provider type"))
		return
	case ExternalPasswordConnector:
		// Login with password, like LDAP.

		if r.Method != http.MethodPost {
			ui.ResponseErrorWithCode(w, http.StatusMethodNotAllowed, nil)
			return
		}

		if req.Username == "" {
			ui.ResponseErrorWithCode(w, http.StatusBadRequest, errors.New("username is required"))
			return
		}
		if req.Password == "" {
			ui.ResponseErrorWithCode(w, http.StatusBadRequest, errors.New("password is required"))
			return
		}

		id, valid, err := cn.Login(ctx, req.Username, req.Password)
		if err != nil {
			ui.ResponseErrorWithCode(w, http.StatusInternalServerError, err)
			return
		}
		if !valid {
			ui.ResponseErrorWithCode(w, http.StatusUnauthorized, errors.New("login failed"))
			return
		}

		subj, err := convertSubjectFromExternalIdentity(ctx, req.Provider, id)
		if err != nil {
			ui.ResponseErrorWithCode(w, http.StatusInternalServerError, fmt.Errorf("get subject: %w", err))
			return
		}

		loginSubject(w, r, subj, false)
	case ExternalCallbackConnector:
		// Redirect to OAuth login page.

		if r.Method != http.MethodGet {
			ui.ResponseErrorWithCode(w, http.StatusMethodNotAllowed, nil)
			return
		}

		// Get callback URL.
		var callbackUrl string
		{
			u := r.URL.ResolveReference(&url.URL{Path: "callback/" + req.Provider})
			u.RawQuery = ""
			u.RawFragment = ""
			u.Fragment = ""
			su := funcx.NoError(systemsetting.ServeUrl.ValueURL(ctx))
			if su == nil || su.Scheme == "" && su.Host == "" {
				u.Scheme = "https"
				u.Host = "localhost"
			} else {
				u.Scheme = su.Scheme
				u.Host = su.Host
			}
			callbackUrl = u.String()
		}

		// Create state.
		sec := &core.Secret{
			ObjectMeta: meta.ObjectMeta{
				Namespace:    systemkuberes.SystemNamespaceName,
				GenerateName: "walrus-subject-login-callback-",
			},
			Data: map[string][]byte{
				"provider":    []byte(req.Provider),
				"clientID":    []byte(cn.GetClientID()),
				"callbackUrl": []byte(callbackUrl),
			},
		}
		sec, err := kubeclientset.CreateWithCtrlClient(ctx, system.LoopbackCtrlClient.Get(), sec)
		if err != nil {
			ui.ResponseErrorWithCode(w, http.StatusInternalServerError, fmt.Errorf("create state: %w", err))
			return
		}

		// Get login URL.
		loginUrl, err := cn.GetLoginURL(callbackUrl, sec.Name)
		if err != nil {
			ui.ResponseErrorWithCode(w, http.StatusInternalServerError, err)
			return
		}

		http.Redirect(w, r, loginUrl, http.StatusFound)
	}
}

type (
	requestCallback struct {
		Provider string `path:"provider"`
		State    string `query:"state"`
	}
)

// callback is a handler to handle callback.
//
// GET: /callback/{provider}?code={code}&state={state}
func callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse request.
	var req requestCallback
	_ = httpx.BindWith(r, &req, httpx.BindQuery, httpx.BindPath)

	// Get provider.
	subjProv, err := getSubjectProvider(ctx, req.Provider)
	if err != nil {
		ui.RedirectErrorWithCode(w, r, http.StatusInternalServerError, fmt.Errorf("get provider: %w", err))
		return
	}

	// Get connector.
	var cn ExternalCallbackConnector
	{
		conn, err := getExternalConnectorFromSubjectProvider(subjProv)
		if err != nil {
			ui.RedirectErrorWithCode(w, r, http.StatusInternalServerError, err)
			return
		}
		var ok bool
		cn, ok = conn.(ExternalCallbackConnector)
		if !ok {
			http.Error(w, "unsupported provider type", http.StatusBadRequest)
			return
		}
	}

	// Verify state and get callback URL.
	var callbackUrl string
	{
		sec := &core.Secret{
			ObjectMeta: meta.ObjectMeta{
				Namespace: subjProv.Namespace,
				Name:      req.State,
			},
		}
		cli := system.LoopbackCtrlClient.Get()
		err = cli.Get(ctx, ctrlcli.ObjectKeyFromObject(sec), sec)
		if err != nil {
			if !kerrors.IsNotFound(err) {
				ui.RedirectErrorWithCode(w, r, http.StatusInternalServerError, fmt.Errorf("get state: %w", err))
			} else {
				ui.RedirectErrorWithCode(w, r, http.StatusForbidden, errors.New("state not found"))
			}
			return
		}
		_ = cli.Delete(ctx, sec) // Always delete.

		err = func() error {
			switch {
			case string(sec.Data["provider"]) != req.Provider:
				return errors.New("provider mismatch")
			case string(sec.Data["clientID"]) != cn.GetClientID():
				return errors.New("client id mismatch")
			case string(sec.Data["callbackUrl"]) == "":
				return errors.New("callback URL not found")
			case time.Since(sec.CreationTimestamp.Time) > 5*time.Minute:
				return errors.New("state expired")
			}
			return nil
		}()
		if err != nil {
			ui.RedirectErrorWithCode(w, r, http.StatusForbidden, err)
			return
		}

		callbackUrl = string(sec.Data["callbackUrl"])
	}

	// Handle callback.
	id, err := cn.HandleCallback(callbackUrl, r)
	if err != nil {
		ui.RedirectErrorWithCode(w, r, http.StatusInternalServerError, fmt.Errorf("handle callback: %w", err))
		return
	}

	// Get subject.
	subj, err := convertSubjectFromExternalIdentity(ctx, req.Provider, id)
	if err != nil {
		ui.RedirectErrorWithCode(w, r, http.StatusInternalServerError, fmt.Errorf("get subject: %w", err))
		return
	}

	// Login.
	loginSubject(w, r, subj, true)
}

type (
	requestProfile struct {
		DisplayName *string `json:"displayName,omitempty"`
		Email       *string `json:"email,omitempty"`
		Password    *string `json:"password,omitempty"`
	}
	responseProfile struct {
		Name        string   `json:"name"`
		Provider    string   `json:"provider"`
		Role        string   `json:"role"`
		DisplayName string   `json:"displayName,omitempty"`
		Email       string   `json:"email,omitempty"`
		Groups      []string `json:"groups,omitempty"`
	}
)

// profile is a handler to get/update profile.
//
// GET/PUT: /profile
func profile(w http.ResponseWriter, r *http.Request) {
	// Get kube config.
	subjNamespace, subjName, cliCfg := GetSubjectKubeConfig(r)
	if cliCfg == nil {
		ui.ResponseErrorWithCode(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	// Get kube client.
	cli, err := clientset.NewForConfig(cliCfg)
	if err != nil {
		ui.ResponseErrorWithCode(w, http.StatusInternalServerError, fmt.Errorf("get kube client: %w", err))
		return
	}

	if r.Method == http.MethodGet {
		// Get profile.
		subj, err := cli.WalrusV1().Subjects(subjNamespace).
			Get(r.Context(), subjName, meta.GetOptions{ResourceVersion: "0"})
		if err != nil {
			ui.ResponseError(w, fmt.Errorf("get profile: %w", err))
			return
		}

		resp := responseProfile{
			Name:        subj.Name,
			Provider:    subj.Spec.Provider,
			Role:        subj.Spec.Role.String(),
			DisplayName: subj.Spec.DisplayName,
			Email:       subj.Spec.Email,
			Groups:      subj.Spec.Groups,
		}
		httpx.JSON(w, http.StatusOK, resp)
		return
	}

	// Parse request.
	var req requestProfile
	_ = httpx.BindJSON(r, &req)

	// Update profile.
	subjChanged := applywalrus.Subject(subjName, subjNamespace).
		WithSpec(applywalrus.SubjectSpec())
	if req.DisplayName != nil {
		subjChanged.Spec.WithDisplayName(*req.DisplayName)
	}
	if req.Email != nil {
		subjChanged.Spec.WithEmail(*req.Email)
	}
	if req.Password != nil {
		subjChanged.Spec.WithCredential(*req.Password)
	}
	subj, err := cli.WalrusV1().Subjects(subjNamespace).
		Apply(r.Context(), subjChanged, meta.ApplyOptions{FieldManager: "identify"})
	if err != nil {
		ui.ResponseError(w, fmt.Errorf("update profile: %w", err))
		return
	}

	resp := responseProfile{
		Name:        subj.Name,
		Provider:    subj.Spec.Provider,
		Role:        subj.Spec.Role.String(),
		DisplayName: subj.Spec.DisplayName,
		Email:       subj.Spec.Email,
		Groups:      subj.Spec.Groups,
	}
	httpx.JSON(w, http.StatusOK, resp)
}

type (
	requestToken struct {
		ExpirationSeconds *int64 `query:"expirationSeconds"`
	}
)

// token is a handler to get token.
//
// GET: /token?expirationSeconds={expirationSeconds}
func token(w http.ResponseWriter, r *http.Request) {
	// Get kube config.
	subjNamespace, subjName, cliCfg := GetSubjectKubeConfig(r)
	if cliCfg == nil {
		ui.ResponseErrorWithCode(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	// Get kube client.
	cli, err := clientset.NewForConfig(cliCfg)
	if err != nil {
		ui.ResponseErrorWithCode(w, http.StatusInternalServerError, fmt.Errorf("get kube client: %w", err))
		return
	}

	// Parse request.
	var req requestToken
	_ = httpx.BindQuery(r, &req)

	// Create.
	subjToken := &walrus.SubjectToken{
		Spec: walrus.SubjectTokenSpec{
			ExpirationSeconds: req.ExpirationSeconds,
		},
	}
	subjToken, err = cli.WalrusV1().Subjects(subjNamespace).
		CreateToken(r.Context(), subjName, subjToken, meta.CreateOptions{})
	if err != nil {
		ui.ResponseError(w, fmt.Errorf("create token: %w", err))
		return
	}

	resp := subjToken.Status
	httpx.JSON(w, http.StatusOK, resp)
}

type (
	requestRules struct {
		Namespace string `path:"namespace"`
	}
	responseRules struct {
		Items []authorization.ResourceRule `json:"items"`
	}
)

// rules is a handler to get rules.
//
// GET: /rules/{namespace}
func rules(w http.ResponseWriter, r *http.Request) {
	// Get kube config.
	_, _, cliCfg := GetSubjectKubeConfig(r)
	if cliCfg == nil {
		ui.ResponseErrorWithCode(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	// Get kube client.
	cli, err := clientset.NewForConfig(cliCfg)
	if err != nil {
		ui.ResponseErrorWithCode(w, http.StatusInternalServerError, fmt.Errorf("get kube client: %w", err))
		return
	}

	// Parse request.
	var req requestRules
	_ = httpx.BindJSON(r, &req)

	// Get rules.
	rev := &authorization.SelfSubjectRulesReview{
		Spec: authorization.SelfSubjectRulesReviewSpec{
			Namespace: req.Namespace,
		},
	}
	rev, err = cli.AuthorizationV1().SelfSubjectRulesReviews().
		Create(r.Context(), rev, meta.CreateOptions{})
	if err != nil {
		ui.ResponseError(w, fmt.Errorf("create self subject rules reviews: %w", err))
		return
	}

	resp := responseRules{
		Items: make([]authorization.ResourceRule, 0, len(rev.Status.ResourceRules)),
	}
	for i := range rev.Status.ResourceRules {
		var found bool
		for j := range rev.Status.ResourceRules[i].APIGroups {
			if rev.Status.ResourceRules[i].APIGroups[j] != walrus.GroupName {
				continue
			}
			found = true
			break
		}
		if !found {
			continue
		}
		item := rev.Status.ResourceRules[i]
		item.APIGroups = []string{walrus.GroupName}
		resp.Items = append(resp.Items, item)
	}
	httpx.JSON(w, http.StatusOK, resp)
}

// logout is a handler to log out.
//
// GET: /logout
func logout(w http.ResponseWriter, r *http.Request) {
	revertSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}

const _AuthenticationCookie = "walrus_session"

// fetchSession fetches the session token from the request.
func fetchSession(r *http.Request) string {
	if r == nil {
		return ""
	}

	c, err := r.Cookie(_AuthenticationCookie)
	if err != nil {
		return ""
	}

	return c.Value
}

// assignSession assigns a session token to the response writer.
func assignSession(w http.ResponseWriter, token string) {
	exp := funcx.NoError(jwt.Parse(stringx.ToBytes(&token))).Expiration()
	c := &http.Cookie{
		Name:     _AuthenticationCookie,
		Value:    token,
		Path:     "/",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
		Expires:  exp,
		MaxAge:   int(time.Until(exp).Round(time.Second) / time.Second),
	}
	http.SetCookie(w, c)
}

// revertSession reverts the session token from the response writer.
func revertSession(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:     _AuthenticationCookie,
		Value:    "",
		Path:     "/",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
		MaxAge:   -1,
		Expires:  time.Now().Add(-time.Hour),
	}
	http.SetCookie(w, c)
}

// getSubjectProvider gets the subject provider.
func getSubjectProvider(ctx context.Context, provider string) (*walrus.SubjectProvider, error) {
	subjProv := &walrus.SubjectProvider{
		ObjectMeta: meta.ObjectMeta{
			Namespace: systemkuberes.SystemNamespaceName,
			Name:      provider,
		},
	}
	cli := system.LoopbackCtrlClient.Get()
	err := cli.Get(ctx, ctrlcli.ObjectKeyFromObject(subjProv), subjProv)
	if err != nil {
		return nil, err
	}
	return subjProv, nil
}

// loginSubject logs in with the subject.
func loginSubject(w http.ResponseWriter, r *http.Request, subj *walrus.Subject, redirect bool) {
	cli := system.LoopbackKubeClient.Get()

	subjl := &walrus.SubjectLogin{
		ObjectMeta: meta.ObjectMeta{
			Namespace: systemkuberes.SystemNamespaceName,
			Name:      subj.Name,
		},
		Spec: walrus.SubjectLoginSpec{
			Credential: *subj.Spec.Credential,
		},
	}

	subjl, err := cli.WalrusV1().Subjects(systemkuberes.SystemNamespaceName).
		Login(r.Context(), subj.Name, subjl, meta.CreateOptions{})
	if err != nil {
		if redirect {
			ui.RedirectError(w, r, fmt.Errorf("login: %w", err))
		} else {
			ui.ResponseError(w, fmt.Errorf("login: %w", err))
		}
		return
	}

	assignSession(w, subjl.Status.Token)
	if redirect {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// GetSubjectKubeConfig returns subject-specified Kubernetes rest config and subject names according to the request.
//
// GetSubjectKubeConfig returns nil rest config if the request is not authorized.
func GetSubjectKubeConfig(r *http.Request) (subjNamespace, subjName string, cliCfg *rest.Config) {
	loopbackCliCfg := system.LoopbackKubeClientConfig.Get()

	s := fetchSession(r)
	if s == "" {
		user, ok := genericapirequest.UserFrom(r.Context())
		if ok {
			subjNamespace, subjName, ok = systemauthz.ConvertSubjectNamesFromAuthnUser(user)
			if ok {
				cliCfg = kubeconfig.AuthorizeRestConfigWithAuthInfo(loopbackCliCfg,
					clientcmdapi.AuthInfo{
						Impersonate:          user.GetName(),
						ImpersonateUID:       user.GetUID(),
						ImpersonateGroups:    user.GetGroups(),
						ImpersonateUserExtra: user.GetExtra(),
					})
				return subjNamespace, subjName, cliCfg
			}
		}
	} else {
		t, err := jwt.Parse(stringx.ToBytes(&s))
		if err == nil {
			var ok bool
			subjNamespace, subjName, ok = systemauthz.ConvertSubjectNamesFromJwtSubject(t.Subject())
			if ok {
				cliCfg = kubeconfig.AuthorizeRestConfigWithAuthInfo(loopbackCliCfg,
					clientcmdapi.AuthInfo{
						Token: s,
					})
				return subjNamespace, subjName, cliCfg
			}
		}
	}

	subjNamespace, subjName = systemkuberes.SystemNamespaceName, systemkuberes.AdminSubjectName
	if system.DisableAuths.Get() {
		cliCfg = kubeconfig.AuthorizeRestConfigWithAuthInfo(loopbackCliCfg,
			clientcmdapi.AuthInfo{
				Impersonate: systemauthz.ConvertImpersonateUserFromSubjectName(subjNamespace, subjName),
			})
		return subjNamespace, subjName, cliCfg
	}

	return "", "", nil
}
