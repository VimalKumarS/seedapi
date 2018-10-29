package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// JSONResponse tries to convert the given data into a json.
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// ReadJSONBody reads the body of the given request.
func ReadJSONBody(r *http.Request, data interface{}) error {
	body, err := ReadBody(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

// ReadBody reads the body of the given http.Request.
func ReadBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()

	// replace the body in the request so that it can be red again
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
