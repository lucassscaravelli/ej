package from

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type fileSuite struct {
	suite.Suite
}

func (f *fileSuite) TestRead__OpenError() {

	data, err := File("file0.json").Read()

	f.EqualError(err, "open: open file0.json: no such file or directory")
	f.Nil(data)
}

func (f *fileSuite) TestRead__NoError() {
	f.NoError(ioutil.WriteFile("file1.json", []byte(`{"country": "brazil"}`), os.ModePerm))

	data, err := File("file1.json").Read()

	f.Nil(err)
	f.Equal(`{"country": "brazil"}`, string(data))

	f.NoError(os.Remove("file1.json"))
}

func (f *fileSuite) TestWrite() {
	err := File("file2.json").Write([]byte(`{"country": "brazil"}`))

	f.Nil(err)

	file, err := os.Open("file2.json")
	f.Nil(err)

	data, err := ioutil.ReadAll(file)
	f.Nil(err)
	f.Equal(`{"country": "brazil"}`, string(data))

	file.Close()

	f.NoError(os.Remove("file2.json"))
}

func TestFile(t *testing.T) {
	suite.Run(t, new(fileSuite))
}
