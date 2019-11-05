// Package ahref assists in parsing an HTML file and extracting links
package ahref

// Ahref is an <a> tag composed of link and text displayed
type Ahref struct {
	Href string
	Text string
}

// ParseAhref should ultimately return []Ahref
func ParseAhref(html string) string {
	return html
}
