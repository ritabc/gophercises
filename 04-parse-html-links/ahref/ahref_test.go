package ahref

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var testCases = []struct {
	exFile string
	output Ahref
}{
	{
		exFile: "./examples/one.html",
		output: Ahref{},
	},
	{
		exFile: "./examples/two.html",
		output: Ahref{},
	},
	{
		exFile: "./examples/three.html",
		output: Ahref{},
	},
	{
		exFile: "./examples/four.html",
		output: Ahref{},
	},
}

func TestParseAhref(t *testing.T) {
	for _, test := range testCases {
		file, err := os.Open(test.exFile)
		if err != nil {
			t.Errorf("Could not open file: %s\n", test.exFile)
		}
		defer file.Close()

		contents, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("Could not read file: %s\n", test.exFile)
		}

		fmt.Printf("%s\n\n\n\n\n\n", string(contents))
		if string(contents) != ParseAhref(string(contents)) {
			t.Errorf("Not a match")
		}
	}
}
