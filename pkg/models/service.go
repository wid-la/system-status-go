package models

// Service ...
type Service struct {
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	Status string  `json:"status"`
	Issues []Issue `json:"issues"`
}
