package main

import (
	"fmt"
	"gophercises/03-cyoa/cyoa"
	"net/http"
	"os"
)

// type dataForTemplate struct {
// 	Path  cyoa.ShortTitle
// 	Scene cyoa.Scene
// }

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
	http.ListenAndServe(":8080", h)

	// Listen and Serve
	// for shortTitle, scene := range arc {
	// 	path := fmt.Sprintf("/%s", shortTitle)
	// 	currentScene := scene
	// 	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
	// 		t, err := template.ParseFiles("../../template.html")
	// 		if err != nil {
	// 			fmt.Printf("Error parsing template: %s", err.Error())
	// 		}
	// 		err = t.Execute(w, currentScene)
	// 		if err != nil {
	// 			fmt.Printf("Error executing template: %s", err.Error())
	// 		}
	// 	})
	// }
}

// func cyoaHandler(arc cyoa.Arc) *http.ServeMux {
// 	mux := http.NewServeMux()
// 	for path, scene := range arc {
// 		mux.HandleFunc(fmt.Sprintf("/%s", path), func(resWriter http.ResponseWriter, req *http.Request) {
// 			fmt.Fprintf(resWriter, scene.FullTitle)
// 		})
// 	}
// 	return mux
// }
