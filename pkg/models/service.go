package models

import (
	"time"
)

// Service ...
type Service struct {
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	Desc        string  `json:"desc"`
	Status      string  `json:"status"`
	LastUpdated string  `json:"lastUpdated"`
	TimeOut     int     `json:"timeOut"`
	Issues      []Issue `json:"issues,omitempty"`
}

// CheckStatus is ...
func (o *Service) CheckStatus() {
	t1, _ := time.Parse(
		time.RFC3339,
		o.LastUpdated)

	now := time.Now()
	diff := now.Sub(t1)

	if o.Status == "UP" && int(diff.Seconds()) > o.TimeOut {
		o.Status = "DOWN"
	}
}
