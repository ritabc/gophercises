package ahref

import (
	"testing"
)

var testCases = []struct {
	exFile string
	output Href
}{}

func TestHref(t *testing.T) {
	for _, test := range testCases {
		// if ok := Valid(test.input); ok != test.ok {
		// 	t.Fatalf("Valid(%s): %s\n\t Expected: %t\n\t Got: %t", test.input, test.description, test.ok, ok)
		// }
	}
}
