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

// ParseAhref should ultimately return []Ahref
func ParseAhref(file io.Reader) ([]Ahref, error) {
	var links []Ahref
	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}
	var getLinks func(*html.Node)
	getLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, att := range n.Attr {
				if att.Key == "href" {

					// don't use closures FOR NOW
					var out strings.Builder
					getVisibleText(n, &out)
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

func getVisibleText(n *html.Node, out *strings.Builder) {
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			out.WriteString(child.Data)
		}
		getVisibleText(child, out)
	}
}
