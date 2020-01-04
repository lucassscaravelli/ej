# Easy Json (EJ)

[![Go Report Card](https://goreportcard.com/badge/github.com/lucassscaravelli/ej)](https://goreportcard.com/report/github.com/lucassscaravelli/ej)
[![Build Status](https://travis-ci.org/lucassscaravelli/ej.svg?branch=master)](https://travis-ci.org/lucassscaravelli/ej)
[![Coverage Status](https://coveralls.io/repos/github/lucassscaravelli/ej/badge.svg?branch=master)](https://coveralls.io/github/lucassscaravelli/ej?branch=master)
[![GoDoc](https://godoc.org/github.com/lucassscaravelli/ej?status.svg)](https://godoc.org/github.com/lucassscaravelli/ej)

Package ```ej``` implements a JSON handler to write and read json succinctly from different sources like files and http requests.

## Install

```
go get -u github.com/lucassscaravelli/ej
```

## Examples

### File

```go
package main

type exData struct {
	Hello int
	World []string
}

func main() {

	dataWrite := &exData{
		Hello: 1,
		World: []string{"h", "e", "l", "l", "o"},
	}

	var dataRead exData
	
	// marshal the content of "dataWrite" to JSON and write in "ex.json" file
	if err := ej.JSON(from.File("ex.json")).Write(&dataWrite); err != nil {
		log.Fatal(err)
	}
	
	// read the data of "ex.json" file and unmarshal the JSON to "dataRead" content
	if err := ej.JSON(from.File("ex.json")).ParseToData(&dataRead); err != nil {
		log.Fatal(err)
	}
	
	// equal content
	fmt.Printf("dataWrite: %+v\n", dataWrite)
	fmt.Printf("dataRead: %+v\n", dataRead)
}

```

### HTTP Request

```go 
package main

import (
	"log"
	"net/http"

	"github.com/lucassscaravelil/ej"
	"github.com/lucassscaravelil/ej/from"
)

type requestPayload struct {
	NumberToFind int
	Numbers      []int
}

type responseErrorPayload struct {
	StatusCode int
	ErrorTxt   string
}

type responsePayload struct {
	Found  bool
	Number int
}

func writeError(jsonHandler *ej.EJ, status int, err error) {
	jsonHandler.Write(&responseErrorPayload{
		StatusCode: status,
		ErrorTxt:   err.Error(),
	})
}

func main() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		var bodyData requestPayload

		jsonHandler := ej.JSON(from.Request(w, r))

		if err := jsonHandler.ParseToData(&bodyData); err != nil {
			writeError(jsonHandler, http.StatusBadRequest, err)
			return
		}

		found := false
		foundNumber := 0

		for _, number := range bodyData.Numbers {
			if number == bodyData.NumberToFind {
				found = true
				foundNumber = number
				break
			}
		}

		jsonHandler.Write(&responsePayload{
			Found:  found,
			Number: foundNumber,
		})

		return
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
