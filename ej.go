package ej

import (
	"encoding/json"
	"fmt"
)

// A JSONHandler read and write bytes
// from any source.
type JSONHandler interface {
	Read() ([]byte, error)
	Write([]byte) error
}

// EJ - Easy JSON struct
type EJ struct {
	jsonHandler JSONHandler
}

// JSON receives a JSONHandler and returns
// a new JSONHandler instance
func JSON(JSONh JSONHandler) *EJ {
	return &EJ{
		jsonHandler: JSONh,
	}
}

// ParseToData parses the JSON-encoded data and stores the result
// in the value pointed to by "data"
func (e *EJ) ParseToData(data interface{}) error {

	value, err := e.jsonHandler.Read()

	if err != nil {
		return fmt.Errorf("read: %s", err.Error())
	}

	if err := json.Unmarshal(value, data); err != nil {
		return fmt.Errorf("json unmarshal: %s", err.Error())
	}

	return nil
}

// Write the JSON encoding of "data" in the 
// source of JSONHandler
func (e *EJ) Write(data interface{}) error {

	value, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("json marshal: %s", err.Error())
	}

	if err := e.jsonHandler.Write(value); err != nil {
		return fmt.Errorf("write: %s", err.Error())
	}

	return nil
}
