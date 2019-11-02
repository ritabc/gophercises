package main

import (
	"encoding/json"
	"fmt"
	"gophercises/03-cyoa/cyoa"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Open JSON File
	jsonFile, err := os.Open("../gopher.json")
	if err != nil {
		fmt.Println("Error opening JSON file")
	}
	defer jsonFile.Close()

	// Read JSON file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file")
	}

	// Unmarshal JSON data
	var arc cyoa.Arc = make(map[cyoa.ShortTitle]cyoa.Scene)
	err = json.Unmarshal(byteValue, &arc)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data")
	}

	// Listen and Serve
	cyoaHandleFunc := cyoaHandler(arc)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", cyoaHandleFunc)
}

func cyoaHandler(arc cyoa.Arc) http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		arc, found := arc["intro"]
		if found {
			fmt.Fprintln(resWriter, arc.FullTitle)
		}
	}
}
