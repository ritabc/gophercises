package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Arc is the entire story
type Arc map[ShortTitle]Scene

// ShortTitle is the key or path, for eg 'intro' or 'sean-kelly'
type ShortTitle string

// Scene consists of the title, scene setup, and options
type Scene struct {
	FullTitle string   `json:"title"`
	Story     []string `json:"story"`
	Options   []Option `json:"options"`
}

// Option is an option or choice the user can make. It includes a choice and an Outcome (story)
type Option struct {
	Choice  string     `json:"text"`
	Outcome ShortTitle `json:"arc"`
}

func JSONStory(r io.Reader) (Arc, error) {
	// Read JSON file
	d := json.NewDecoder(r)
	// Unmarshal JSON data
	var arc Arc = make(map[ShortTitle]Scene)
	if err := d.Decode(&arc); err != nil {
		return nil, err
	}

	return arc, nil
}

func NewHandler(a Arc) http.Handler {
	return handler{a}
}

type handler struct {
	a Arc
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get html template file
	prefix, err := filepath.Abs("../../html/")
	if err != nil {
		fmt.Println("Error getting absolute filepath: ", err.Error())
	}
	templatePath := prefix + "/template.html"
	tpl := template.Must(template.ParseFiles(templatePath))

	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	// "/intro" => "intro" => ShortTitle{path}
	path = path[1:]
	title := ShortTitle(path)

	if chapter, ok := h.a[title]; ok {
		err = tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err.Error())
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)

}
