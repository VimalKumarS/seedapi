package model

import "net/http"

// HttpHandlerWithError http handler return Error
type HttpHandlerWithError func(http.ResponseWriter, *http.Request) error

//Route : route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HttpHandlerWithError
	Protected   bool
}
