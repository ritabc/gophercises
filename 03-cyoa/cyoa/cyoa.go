package cyoa

import (
// "encoding/json"
)

// Arc is the entire story
type Arc map[ShortTitle]Scene

// ShortTitle is the key or path, for eg 'intro' or 'sean-kelly'
type ShortTitle string // escape

// Scene consists of the title, scene setup, and options
type Scene struct {
	FullTitle string   `json:"title"` // escape
	Story     []string `json:"story"` // escape
	Options   []Option `json:"options"`
}

// Option is an option or choice the user can make. It includes a choice and an Outcome (story)
type Option struct {
	Choice  string     `json:"text"` // escape
	Outcome ShortTitle `json:"arc"`
}
