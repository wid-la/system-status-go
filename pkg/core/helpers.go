package core

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	postMethod = "POST"
	getMethod  = "GET"
)

// GetJSON is ...
func GetJSON(r *http.Request, target interface{}) error {
	var body io.Reader
	if r.Method == "PUT" || r.Method == postMethod {
		body = r.Body
	} else {
		obj := make(map[string]string)
		u, _ := url.Parse(r.RequestURI)
		m, _ := url.ParseQuery(u.RawQuery)

		for k, v := range m {
			obj[k] = v[0]
		}

		mapObj, _ := json.Marshal(obj)
		body = bytes.NewReader(mapObj)
	}
	return json.NewDecoder(body).Decode(&target)
}

func envString(env string, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func envInt(env string, fallback int) int {
	e := envString(env, "")
	if e == "" {
		return fallback
	}
	res, err := strconv.Atoi(e)
	if err != nil {
		return fallback
	}
	return res
}
