package api

import (
	"encoding/json"
	"net/http"
)

// NewService ...
func (api API) NewService(w http.ResponseWriter, r *http.Request) {

}

// GetServices ...
func (api API) GetServices(w http.ResponseWriter, r *http.Request) {
	jsonRes, err := json.Marshal(api.Services)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

// GetService ...
func (api API) GetService(w http.ResponseWriter, r *http.Request) {

}

// UpdateService ...
func (api API) UpdateService(w http.ResponseWriter, r *http.Request) {

}
