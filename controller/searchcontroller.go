package controller

import (
	"net/http"
	"seedapi/serviceinterface"
	errors "seedapi/util"
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
func (c *SearchController) Index(w http.ResponseWriter, r *http.Request) error {
	k, err := c.searchInterfacae.GetAll()

	// if c.HandleError(err, w) {
	// 	return
	// }

	if err != nil {
		err = errors.BadRequest.Wrapf(err, "Custom error or validation error")
		err = errors.AddErrorContext(err, "id", "wrong id format, should be an integer")
		return err
	}

	c.SendJSON(w, &k, http.StatusOK)
	return nil
}

// Index func return all search in database
func (c *SearchController) CustomError(w http.ResponseWriter, r *http.Request) error {
	var err error
	err = errors.New("New error created")
	err = errors.BadRequest.Wrapf(err, "Custom error or validation error")
	err = errors.AddErrorContext(err, "id", "wrong id format, should be an integer")
	err = errors.AddErrorContext(err, "id1", "wrong id format, should be an integer")
	return err

}

// Index func return all search in database
func (c *SearchController) PanicRequest(w http.ResponseWriter, r *http.Request) error {
	panic("Custom error")
	return nil
}
