package webserver

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/seal-io/walrus/pkg/server/webserver/clis"
	"github.com/seal-io/walrus/pkg/server/webserver/identify"
	"github.com/seal-io/walrus/pkg/server/webserver/loopback"
	"github.com/seal-io/walrus/pkg/server/webserver/openapi"
	"github.com/seal-io/walrus/pkg/server/webserver/swagger"
	"github.com/seal-io/walrus/pkg/server/webserver/ui"
)

func Index() http.Handler {
	r := mux.NewRouter()

	// Extension routes.
	clisOpenapi := clis.Route(r.PathPrefix("/clis").Methods(http.MethodGet))
	identifyOpenapi := identify.Route(r.PathPrefix("/identify"))
	loopbackOpenapi := loopback.Route(r.PathPrefix("/loopback"))
	openapi.Route(r.PathPrefix("/openapi").Methods(http.MethodGet),
		clisOpenapi, identifyOpenapi, loopbackOpenapi)

	// Swagger UI.
	r.Path("/swagger").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})
	swagger.Route(r.PathPrefix("/swagger").Methods(http.MethodGet))

	// UI.
	r.NotFoundHandler = ui.Index()

	return r
}
