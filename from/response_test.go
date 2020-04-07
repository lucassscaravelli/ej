package from

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucassscaravelli/ej"

	"github.com/stretchr/testify/suite"
)

type testData struct {
	Count int
	Txt   string
}

type responseSuite struct {
	suite.Suite
}

func (r *responseSuite) TestRead__NilResponse() {
	r.Equal(
		fmt.Errorf("read: %s", errNilResponse.Error()),
		ej.JSON(Response(nil)).ParseToData(nil),
	)
}

func (r *responseSuite) TestRead__NoError() {
	expectedData := testData{
		Count: 123,
		Txt:   "space-engineers",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ej.JSON(Request(w, r)).Write(&expectedData)
	}))

	var responseData testData

	response, err := http.Get(server.URL)
	r.Nil(err)
	r.Nil(ej.JSON(Response(response)).ParseToData(&responseData))
	r.Equal(expectedData, responseData)
}

func (r *responseSuite) TestWrite__WriteNotAllowed() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	response, err := http.Get(server.URL)

	r.Nil(err)
	r.Equal(
		fmt.Errorf("write: %s", errWriteNotAllowed.Error()),
		ej.JSON(Response(response)).Write(nil),
	)
}

func TestResponse(t *testing.T) {
	suite.Run(t, new(responseSuite))
}
