package from

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileHandler implements a JSONHandler interface to
// extract data from a file
type FileHandler struct {
	filename string
}

// File returns a new instance of FileHandler
// struct
func File(filename string) *FileHandler {
	return &FileHandler{
		filename: filename,
	}
}

// Read the file data
func (f *FileHandler) Read() ([]byte, error) {

	file, err := os.Open(f.filename)

	if err != nil {
		return nil, fmt.Errorf("open: %s", err.Error())
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("read all: %s", err.Error())
	}

	return bytes, nil
}

// Write data in a file
func (f *FileHandler) Write(data []byte) error {
	return ioutil.WriteFile(f.filename, data, os.ModePerm)
}
