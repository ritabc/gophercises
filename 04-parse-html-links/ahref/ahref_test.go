package ahref

import (
	"os"
	"reflect"
	"testing"
)

var testCases = []struct {
	exFile string
	output []Ahref
}{
	{
		exFile: "./examples/one.html",
		output: []Ahref{
			{
				Href: "/other-page",
				Text: "A link to another page",
			},
		},
	},
	{
		exFile: "./examples/two.html",
		output: []Ahref{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			},
		},
	},
	{
		exFile: "./examples/three.html",
		output: []Ahref{
			{
				Href: "#",
				Text: "Login",
			},
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
		},
	},
	{
		exFile: "./examples/four.html",
		output: []Ahref{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		},
	},
	{
		exFile: "./examples/five.html",
		output: []Ahref{
			{
				Href: "/dog",
				Text: "Something in a span IN A SPAN in a span Text not in a span Bold text!",
			},
		},
	},
}

func TestParseAhref(t *testing.T) {
	for _, test := range testCases {
		file, err := os.Open(test.exFile)
		if err != nil {
			t.Errorf("Could not open file: %s\n", test.exFile)
		}
		defer file.Close()

		links, err := ParseAhref(file)
		if err != nil {
			t.Errorf("Error parsing file for links: %s", err.Error())
		}

		if !reflect.DeepEqual(test.output, links) {
			t.Errorf("Expected output for file: %s was: %+v\n\nLinks found were: %+v", test.exFile, test.output, links)
		}
	}
}
