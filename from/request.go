package from

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	errNilResponseWriter = errors.New("response writer is nil")
	errNilRequest        = errors.New("request is nil")
	errNilBody           = errors.New("request body is nil")
)

// RequestHandler implements a JSONHandler interface 
// to extract data from a http request
type RequestHandler struct {
	responseWriter http.ResponseWriter
	request        *http.Request
}

// Request returns a new instance of RequestHandler
// struct
func Request(w http.ResponseWriter, r *http.Request) *RequestHandler {
	return &RequestHandler{
		responseWriter: w,
		request:        r,
	}
}

// Read the data from http request body
func (f *RequestHandler) Read() ([]byte, error) {

	if f.request == nil {
		return nil, errNilRequest
	}

	if f.request.Body == nil {
		return nil, errNilBody
	}

	data, err := ioutil.ReadAll(f.request.Body)

	if err != nil {
		return nil, fmt.Errorf("read body: %s", err.Error())
	}

	return data, nil
}

// Write data in http response writer
func (f *RequestHandler) Write(data []byte) error {

	if f.responseWriter == nil {
		return errNilResponseWriter
	}

	_, err := f.responseWriter.Write(data)

	if err != nil {
		return fmt.Errorf("response write: %s", err.Error())
	}

	return err
}
