package models

type JsonData struct {
	Get *[]Get `json:"get,omitempty"`
	Set *[]Set `json:"set,omitempty"`
}

type Get struct {
	Path  string `json:"path"`
	Query *Query `json:"query,omitempty"`
}

type Set struct {
	Path  string `json:"path"`
	Value any    `json:"value"`
}

type Query struct {
	Depth  int    `json:"depth,omitempty"`
	RegExp string `json:"regExp,omitempty"`
}
