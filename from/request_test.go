package from

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type requestSuite struct {
	suite.Suite
}

func (rs *requestSuite) TestRead__RequestNil() {
	data, err := Request(nil, nil).Read()

	rs.Nil(data)
	rs.EqualError(err, errNilRequest.Error())
}

func (rs *requestSuite) TestRead__NoError() {

	json := `{"game": "csgo"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := Request(w, r).Read()

		rs.NoError(err)
		rs.Equal([]byte(json), data)
	}))

	req, _ := http.NewRequest("POST", server.URL, bytes.NewReader([]byte(json)))
	client := &http.Client{}

	_, err := client.Do(req)

	rs.NoError(err)
}

func (rs *requestSuite) TestWrite__ResponseWriterNil() {
	err := Request(nil, nil).Write([]byte{})
	rs.EqualError(err, errNilResponseWriter.Error())
}

func (rs *requestSuite) TestWrite__NoError() {
	json := `{"name": "lucas"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := Request(w, r).Write([]byte(json))
		rs.NoError(err)
	}))

	req, _ := http.NewRequest("POST", server.URL, bytes.NewReader([]byte(json)))
	client := &http.Client{}

	response, err := client.Do(req)

	rs.NoError(err)

	dataResponse, err := ioutil.ReadAll(response.Body)

	rs.NoError(err)
	rs.Equal([]byte(json), dataResponse)
}

func TestRequest(t *testing.T) {
	suite.Run(t, new(requestSuite))
}
