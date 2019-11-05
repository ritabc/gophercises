// Package ahref assists in parsing an HTML file and extracting links
package ahref

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

// Ahref is an <a> tag composed of link and text displayed
type Ahref struct {
	Href string
	Text string
}

// ParseAhref reads from io.Reader and returns list of our link structs, along with error
func ParseAhref(file io.Reader) ([]Ahref, error) {
	var links []Ahref
	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}
	// Declare getLinks() function before defining so that definition doesn't see getLinks as undeclared
	var getLinks func(*html.Node)
	getLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, att := range n.Attr {
				if att.Key == "href" {

					var out strings.Builder
					var getVisibleText func(*html.Node)
					getVisibleText = func(n *html.Node) {
						for child := n.FirstChild; child != nil; child = child.NextSibling {
							if child.Type == html.TextNode {
								out.WriteString(child.Data)
							}
							getVisibleText(child)
						}
					}
					// Actually make function call(s) and write to out
					getVisibleText(n)

					links = append(links, Ahref{
						Href: att.Val,
						Text: strings.TrimSpace(out.String()),
					})
				}
			}
		}
		for next := n.FirstChild; next != nil; next = next.NextSibling {
			getLinks(next)
		}
	}
	getLinks(doc)
	return links, nil
}
