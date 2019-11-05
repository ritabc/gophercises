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
				href: "/other-page",
				text: "A link to another page",
			},
		},
	},
	{
		exFile: "./examples/two.html",
		output: []Ahref{
			{
				href: "https://www.twitter.com/joncalhoun",
				text: "Check me out on twitter",
			},
			{
				href: "https://github.com/gophercises",
				text: "Gophercises is on Github!",
			},
		},
	},
	{
		exFile: "./examples/three.html",
		output: []Ahref{
			{
				href: "#",
				text: "Login",
			},
			{
				href: "/lost",
				text: "Lost? Need help?",
			},
			{
				href: "https://twitter.com/marcusolsson",
				text: "@marcusolsson",
			},
		},
	},
	{
		exFile: "./examples/four.html",
		output: []Ahref{
			{
				href: "/dog-cat",
				text: "dog cat",
			},
		},
	},
	{
		exFile: "./examples/five.html",
		output: []Ahref{
			{
				href: "/dog",
				text: "Something in a span IN A SPAN in a span Text not in a span Bold text!",
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
			t.Errorf("Error parsing links: %s", err.Error())
		}

		if !reflect.DeepEqual(test.output, links) {
			t.Errorf("Expected output for file: %s was: %+v\n\nLinks found were: %+v", test.exFile, test.output, links)
		}
	}
}
