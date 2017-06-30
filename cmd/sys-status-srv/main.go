package main

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"github.com/wid-la/system-status-go/pkg/api"
	"github.com/wid-la/system-status-go/pkg/core"
	"github.com/wid-la/system-status-go/pkg/interactors"
	"github.com/wid-la/system-status-go/pkg/models"
)

func main() {
	coreDeps := core.NewDependency()

	inter := &interactors.Inter{
		Deps:     coreDeps,
		Services: make(map[string]models.Service),
	}
	router := initRouter(inter)

	fmt.Println(coreDeps.FieldsReport())

	go inter.Process(coreDeps.Interval)
	http.ListenAndServe(coreDeps.Port, router)
}

func initRouter(inter *interactors.Inter) *chi.Mux {
	api := api.API{Inter: inter}

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Mount("/v1", api)
	})

	return router
}
