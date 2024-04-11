package vcssource

import (
	"context"
	"regexp"

	"github.com/drone/go-scm/scm"
	"github.com/seal-io/utils/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemmeta"
	"github.com/seal-io/walrus/pkg/systemsetting"
	"github.com/seal-io/walrus/pkg/templates/kubehelper"
	"github.com/seal-io/walrus/pkg/vcs"
	"github.com/seal-io/walrus/pkg/vcs/options"
)

func New() *Lister {
	return &Lister{}
}

// Lister is a vcs source list implement.
type Lister struct{}

// List returns a list of templates from the given catalog.
func (l *Lister) List(ctx context.Context, c *walruscore.Catalog) ([]walruscore.Template, error) {
	logger := klog.Background().WithName("vcs").WithName("lister")

	var (
		source        = c.Spec.VCSSource
		repos         []*scm.Repository
		filteredRepos []*scm.Repository
		tmpls         []walruscore.Template
	)

	// List.
	{
		opts, err := l.listOptions(ctx)
		if err != nil {
			return nil, err
		}

		repos, err = vcs.GetOrgRepos(ctx, source.Platform, source.URL, opts...)
		if err != nil {
			return nil, err
		}
		logger.Infof("found %d repositories in %s/%s before filtered", len(repos), c.Namespace, c.Name)
	}

	// Filtering.
	{
		var (
			includeReg *regexp.Regexp
			excludeReg *regexp.Regexp
			err        error
		)

		if filters := c.Spec.Filters; filters != nil {
			if filters.IncludeExpression != "" {
				includeReg, err = regexp.Compile(filters.IncludeExpression)
				if err != nil {
					return nil, err
				}
			}

			if filters.ExcludeExpression != "" {
				excludeReg, err = regexp.Compile(filters.ExcludeExpression)
				if err != nil {
					return nil, err
				}
			}
		}

		for i := range repos {
			repo := repos[i]

			if includeReg != nil && !includeReg.MatchString(repo.Name) {
				continue
			}

			if excludeReg != nil && excludeReg.MatchString(repo.Name) {
				continue
			}
			filteredRepos = append(filteredRepos, repo)
		}
	}

	// Generate Templates.
	{
		tmpls = make([]walruscore.Template, len(filteredRepos))
		for i := range filteredRepos {
			repo := filteredRepos[i]

			t := walruscore.Template{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: c.Namespace,
					Name:      kubehelper.NormalizeTemplateName(c.Name, repo.Name),
					Labels:    GenWalrusBuiltinLabels(repo.Topics, c.Name),
				},
				Spec: walruscore.TemplateSpec{
					TemplateFormat: c.Spec.TemplateFormat,
					Description:    repo.Description,
					VCSRepository: &walruscore.VCSRepository{
						Platform: c.Spec.VCSSource.Platform,
						URL:      repo.Link,
					},
				},
				Status: walruscore.TemplateStatus{
					OriginalName: repo.Name,
					Project:      systemmeta.GetProjectName(c.Namespace),
				},
			}

			kubemeta.ControlOn(&t, c, walrus.SchemeGroupVersion.WithKind("Catalog"))

			tmpls[i] = t
		}

		logger.Infof("found %d repositories in %s/%s after filtered", len(tmpls), c.Namespace, c.Name)
	}

	return tmpls, nil
}

func (l *Lister) listOptions(ctx context.Context) ([]options.ClientOption, error) {
	opts := make([]options.ClientOption, 0)

	sid, err := systemsetting.ServeIdentify.Value(ctx)
	if err != nil {
		return nil, err
	}
	ua := version.GetUserAgent() + "; uuid=" + sid
	opts = append(opts, options.WithUserAgent(ua))

	tlsVerify, err := systemsetting.EnableRemoteTlsVerify.ValueBool(ctx)
	if err != nil {
		return nil, err
	}

	if !tlsVerify {
		opts = append(opts, options.WithInsecureSkipVerify())
	}
	return opts, nil
}
