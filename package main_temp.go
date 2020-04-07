package main

import "net/http"

type testData struct {
	Count int
	Txt   string
}

func main() {
	expectedData := testData{
		Count: 123,
		Txt:   "space-engineers",
	}

	var responseData testData
	response, err := http.Get("http://<url>/any")
	if err != nil {
		log.Fatal(err)
	}

	if err := ej.JSON(Response(response)).ParseToData(&responseData); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("dataRead: %+v\n", dataRead)
}
