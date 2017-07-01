package api

import (
	"net/http"
)

// Authenticate ...
func (api API) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Auth-Token")

		if authToken != api.Inter.Deps.Token {
			http.Error(w, http.StatusText(403), 403)
			return
		}

		next.ServeHTTP(w, r)
	})
}
