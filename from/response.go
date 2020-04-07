package from

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	errNilResponse     = errors.New("response is nil")
	errWriteNotAllowed = errors.New("write is not allowed")
)

// ResponseHandler implements a JSONHandler interface
// to extract data from a http response
type ResponseHandler struct {
	response *http.Response
}

// Response returns a new instance of ResponseHandler
// struct
func Response(r *http.Response) *ResponseHandler {
	return &ResponseHandler{
		response: r,
	}
}

// Read the data from http response body
func (r *ResponseHandler) Read() ([]byte, error) {

	if r.response == nil {
		return nil, errNilResponse
	}

	if r.response.Body == nil {
		return nil, errNilBody
	}

	data, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %s", err.Error())
	}

	return data, nil
}

// Write ...
func (r *ResponseHandler) Write(data []byte) error {
	return errWriteNotAllowed
}
