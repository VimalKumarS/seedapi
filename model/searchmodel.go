package model

import "time"

// Search : search test model
type Search struct {
	ID     int               `json:"id,string,omitempty"`
	Name   string            `json:"name,string"`
	Date   time.Time         `json:"date,string"`
	Errors map[string]string `json:"-"`
}

// NewModel create a new search model
func NewModel(name string, date time.Time) *Search {
	return &Search{
		Name: name,
		Date: date,
	}
}

// Validate a Search
func (k *Search) Validate() bool {
	k.Errors = make(map[string]string)

	if k.Name == "" {
		k.Errors["name"] = "name can not be empty"
	}

	if len(k.Errors) > 0 {
		return false
	}

	return true
}
