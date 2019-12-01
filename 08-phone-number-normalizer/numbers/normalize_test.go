package normalize

import (
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "already normalized", input: "1234567890", want: "1234567890"},
		{name: "with spaces", input: "123 456 7890", want: "1234567890"},
		{name: "with dashes", input: "123-456-7890", want: "1234567890"},
		{name: "with dots", input: "123.456.7890", want: "1234567890"},
		{name: "with dots", input: "123.456.7890", want: "1234567890"},
		{name: "with parens", input: "(123)456-7890", want: "1234567890"},
	}

	for _, tt := range tests {
		got := Normalize(tt.input)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tt.name, tt.want, got)
		}
	}
}
