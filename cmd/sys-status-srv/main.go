package main

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"github.com/wid-la/system-status-go/pkg/api"
	"github.com/wid-la/system-status-go/pkg/core"
)

var coreDeps *core.Dependency

func main() {
	coreDeps = core.NewDependency()

	// inter := interactors.Inter{Deps: coreDeps}
	router := initRouter()

	fmt.Println(coreDeps.FieldsReport())

	// go inter.Process(5)
	http.ListenAndServe(coreDeps.Port, router)
}

func initRouter() *chi.Mux {
	api := api.API{Deps: coreDeps}

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Mount("/v1", api)
	})

	return router
}
