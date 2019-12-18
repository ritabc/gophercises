package normalize

import "regexp"

// Normalize takes a phone number and converts it to a normal format (nothing but digits).
// Normalize also validates - returning an error if the number does not have 10 digits
func Normalize(input string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(input, "")
}
