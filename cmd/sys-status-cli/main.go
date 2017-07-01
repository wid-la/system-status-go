package systemStatus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/wid-la/system-status-go/pkg/models"
)

// SystemStatus ...
type SystemStatus struct {
	URL     string
	Token   string
	Service models.Service
}

// NewSystemStatusCli ...
func NewSystemStatusCli(url string, token string) SystemStatus {
	return SystemStatus{URL: url, Token: token}
}

// CreateService ...
func (s *SystemStatus) CreateService(key string, name string, desc string, timeOut int) {
	s.Service = models.Service{
		Key:     key,
		Name:    name,
		Desc:    desc,
		TimeOut: timeOut,
	}

	s.callEndpoint("POST")
}

// UpdateService ...
func (s *SystemStatus) UpdateService() {
	s.callEndpoint("PUT")
}

func (s *SystemStatus) callEndpoint(method string) {
	serviceURL := fmt.Sprintf("https://%s/v1/services", s.URL)

	var URL *url.URL
	URL, err := url.Parse(serviceURL)
	if err != nil {
		fmt.Printf("Error Parsing URL : %v \n", serviceURL)
		os.Exit(-1)
	}

	data, err := json.Marshal(s.Service)
	req, err := http.NewRequest(method, URL.String(), bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error Creating Request : %v \n", err.Error())
		os.Exit(-1)
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("Auth-Token", s.Token)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error Calling Client : %v \n", err.Error())
		fmt.Println(req)
		os.Exit(-1)
	}

	fmt.Println(res.Status)
}
