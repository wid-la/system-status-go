package api

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/wid-la/system-status-go/pkg/interactors"
)

// API ...
type API struct {
	Inter *interactors.Inter
}

func (api API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.router().ServeHTTP(w, r)
}

func (api API) router() http.Handler {
	r := chi.NewRouter()

	r.Use(api.Authenticate)

	r.Route("/services", func(r chi.Router) {
		r.Post("/", api.NewService)
		r.Get("/", api.GetServices)
		r.Route("/:serviceID", func(r chi.Router) {
			r.Get("/", api.GetService)
			r.Put("/", api.UpdateService)
			r.Route("/issues", func(r chi.Router) {
				r.Post("/", api.NewIssue)
				r.Get("/", api.GetIssues)
				r.Route("/:issueID", func(r chi.Router) {
					r.Get("/", api.GetIssue)
					r.Put("/", api.UpdateIssue)
					r.Delete("/", api.DeleteIssue)
				})
			})
		})
	})

	return r
}
