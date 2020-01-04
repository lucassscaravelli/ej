package ej

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var errTest = errors.New("test error")

type dataTest struct {
	Number int
	Txt    string
}

type jsonHandlerMock struct {
	mock.Mock
}

func (j *jsonHandlerMock) Read() ([]byte, error) {
	args := j.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (j *jsonHandlerMock) Write(data []byte) error {
	return j.Called(data).Error(0)
}

type ejSuite struct {
	suite.Suite
}

func (e *ejSuite) TestParseToData__ReadError() {

	jsonHandler := new(jsonHandlerMock)
	jsonHandler.On("Read").Return([]byte{}, errTest)

	e.EqualError(JSON(jsonHandler).ParseToData(nil), "read: test error")
}

func (e *ejSuite) TestParseToData__UnmarshalError() {

	jsonHandler := new(jsonHandlerMock)
	jsonHandler.On("Read").Return([]byte("{{{"), nil)

	e.EqualError(JSON(jsonHandler).ParseToData(nil),
		"json unmarshal: invalid character '{' looking for beginning of object key string")
}

func (e *ejSuite) TestParseToData__NoError() {

	expectedData := dataTest{
		Number: 2,
		Txt:    "hello",
	}

	data := dataTest{}

	jsonHandler := new(jsonHandlerMock)
	jsonHandler.On("Read").Return([]byte(`{"number": 2, "txt": "hello"}`), nil)

	e.NoError(JSON(jsonHandler).ParseToData(&data))
	e.Equal(expectedData, data)
}

func (e *ejSuite) TestWrite__WriteError() {

	expectedData := dataTest{
		Number: 15,
		Txt:    "__hi__",
	}

	data, _ := json.Marshal(expectedData)

	jsonHandler := new(jsonHandlerMock)
	jsonHandler.On("Write", data).Return(errTest)

	e.EqualError(JSON(jsonHandler).Write(expectedData), "write: test error")
}

func (e *ejSuite) TestWrite__NoError() {
	expectedData := dataTest{
		Number: 15,
		Txt:    "__hi__",
	}

	data, _ := json.Marshal(expectedData)

	jsonHandler := new(jsonHandlerMock)
	jsonHandler.On("Write", data).Return(nil)

	e.NoError(JSON(jsonHandler).Write(expectedData))
}

func TestEJ(t *testing.T) {
	suite.Run(t, new(ejSuite))
}
