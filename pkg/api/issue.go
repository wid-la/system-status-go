package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
	"github.com/wid-la/system-status-go/pkg/core"
	"github.com/wid-la/system-status-go/pkg/models"
)

// NewIssue ...
func (api API) NewIssue(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	req := &models.Issue{}
	if err := core.GetJSON(r, &req); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	servID := chi.URLParam(r, "serviceID")
	service := api.Inter.Services[servID]
	service.Issues = append(service.Issues, *req)

	api.Inter.Services[servID] = service

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))
}

// GetIssues ...
func (api API) GetIssues(w http.ResponseWriter, r *http.Request) {
	servID := chi.URLParam(r, "serviceID")
	service := api.Inter.Services[servID]

	jsonRes, err := json.Marshal(service.Issues)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

// GetIssue ...
func (api API) GetIssue(w http.ResponseWriter, r *http.Request) {
	servID := chi.URLParam(r, "serviceID")
	issueID := chi.URLParam(r, "issueID")
	i, _ := strconv.ParseInt(issueID, 10, 64)

	service := api.Inter.Services[servID]

	jsonRes, err := json.Marshal(service.Issues[i])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

// UpdateIssue ...
func (api API) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	servID := chi.URLParam(r, "serviceID")
	issueID := chi.URLParam(r, "issueID")
	i, _ := strconv.ParseInt(issueID, 10, 64)

	// REQUEST
	req := &models.Issue{}
	if err := core.GetJSON(r, &req); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	service := api.Inter.Services[servID]

	service.Issues[i].Title = req.Title
	service.Issues[i].Message = req.Message
	service.Issues[i].Date = req.Date

	api.Inter.Services[servID] = service

	jsonRes, err := json.Marshal(service.Issues[i])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

// DeleteIssue ...
func (api API) DeleteIssue(w http.ResponseWriter, r *http.Request) {
	servID := chi.URLParam(r, "serviceID")
	issueID := chi.URLParam(r, "issueID")
	i, _ := strconv.ParseInt(issueID, 10, 64)

	// REQUEST
	req := &models.Issue{}
	if err := core.GetJSON(r, &req); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	service := api.Inter.Services[servID]
	service.Issues = append(service.Issues[:i], service.Issues[i+1:]...)

	api.Inter.Services[servID] = service

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}
