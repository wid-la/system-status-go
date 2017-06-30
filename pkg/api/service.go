package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pressly/chi"
	"github.com/wid-la/system-status-go/pkg/core"
	"github.com/wid-la/system-status-go/pkg/models"
)

// NewService ...
func (api API) NewService(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	req := &models.Service{}
	if err := core.GetJSON(r, &req); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	req.LastUpdated = time.Now().Format(time.RFC3339)
	if req.TimeOut == 0 {
		req.TimeOut = api.Inter.Deps.Timeout
	}

	api.Inter.Services[req.Key] = *req

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))
}

// GetServices ...
func (api API) GetServices(w http.ResponseWriter, r *http.Request) {
	jsonRes, err := json.Marshal(api.Inter.Services)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

// GetService ...
func (api API) GetService(w http.ResponseWriter, r *http.Request) {
	servID := chi.URLParam(r, "serviceID")

	service := api.Inter.Services[servID]

	jsonRes, err := json.Marshal(service)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

// UpdateService ...
func (api API) UpdateService(w http.ResponseWriter, r *http.Request) {
	servID := chi.URLParam(r, "serviceID")

	// REQUEST
	req := &models.Service{}
	if err := core.GetJSON(r, &req); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	service := api.Inter.Services[servID]
	service.Key = servID
	service.Name = req.Name
	service.Desc = req.Desc
	service.Status = req.Status
	service.LastUpdated = time.Now().Format(time.RFC3339)

	api.Inter.Services[servID] = service

	jsonRes, err := json.Marshal(service)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
