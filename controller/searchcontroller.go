package controller

import (
	"net/http"
	"seedapi/serviceinterface"
)

// SearchController struct
type SearchController struct {
	Controller
	searchInterfacae serviceinterface.SearchInterface
}

// SearchControllerInstance instance
func SearchControllerInstance(s serviceinterface.SearchInterface) *SearchController {
	return &SearchController{
		Controller:       Controller{},
		searchInterfacae: s,
	}
}

// Index func return all search in database
func (c *SearchController) Index(w http.ResponseWriter, r *http.Request) {
	k, err := c.searchInterfacae.GetAll()

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}
