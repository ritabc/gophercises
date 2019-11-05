package main

import (
	"fmt"
	"gophercises/03-cyoa/cyoa"
	"net/http"
	"os"
)

func main() {
	// Open JSON File
	jsonFile, err := os.Open("../../gopher.json")
	if err != nil {
		fmt.Println("Error opening JSON file")
	}
	defer jsonFile.Close()

	arc, err := cyoa.JSONStory(jsonFile)
	if err != nil {
		fmt.Println("Error Unmarshalling JSON")
	}

	h := cyoa.NewHandler(arc)
	fmt.Println("Starting the server on :8080")
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		fmt.Println("Something went wrong: ", err.Error())
	}
}
