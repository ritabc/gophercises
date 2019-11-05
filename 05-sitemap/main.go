package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"gophercises/04-parse-html-links/ahref"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Sitemap struct {
	XMLName  xml.Name `xml:"urlset"`
	Protocol string   `xml:"xmlns,attr"`
	Sites    []URL
}
type URL struct {
	XMLName  xml.Name `xml:"url"`
	Location string   `xml:"loc"`
}

func main() {
	// Setup flags with a default value
	baseURL := flag.String("url", "https://gophercises.com/demos/cyoa/", "URL To Sitemap")
	var depth int
	flag.IntVar(&depth, "depth", 3, "Maximumn depth to traverse")
	flag.Parse()

	// use net/url to get scheme & domain information
	u, err := url.Parse(*baseURL)
	if err != nil {
		fmt.Println("Error parsing input URL: ", err.Error())
	}
	baseURLHost := u.Host
	prefix := fmt.Sprintf("%s://%s", u.Scheme, baseURLHost)

	fmt.Println("Welcome to URL Sitemap")

	// urlGraph is slice of slices: outer has length of depth + 1.
	// BaseURL doesn't count as part of depth
	var urlGraph = make([][]string, depth+1)
	var allURLs = make(map[string]bool) // TODO: Whenever link is added to urlGraph, add it to allURLS also

	// This will probably happen outside all loops
	// We have our base URL, add it to data structures
	urlGraph[0] = []string{*baseURL}
	allURLs[*baseURL] = true

	// range over all levels
	for i := 0; i < depth+1; i++ {

		// urls are all the links found at a certain depth
		var levelURLs []string

		// within each level, follow all links
		for _, linkToFollow := range urlGraph[i] {

			// Get response ƒrom URL GET Request
			resp, err := http.Get(linkToFollow)
			if err != nil {
				fmt.Printf("Error performing GET request to: %v: %s", linkToFollow, err.Error())
			}
			// Search response's body for all links
			pageLinks, err := ahref.ParseAhref(resp.Body)
			if err != nil {
				fmt.Printf("Error parsing body HTML for ahrefs for link: %v: %s", linkToFollow, err.Error())
			}
			// links exist at level i of urlGraph
			for _, ahref := range pageLinks {
				url := ahref.Href

				// clean and validate url
				// So long as url != "/", discard trailing slash for consistency (so hm.org/ and hm.org aren't both added)
				if url != "/" && string(url[len(url)-1]) == "/" {
					url = url[:len(url)-1]
				}

				// discard any "www."
				url = strings.Replace(url, "www.", "", 1)

				// if starts with slash, it's a local url, so prepend `prefix`
				if string(url[0]) == "/" {
					url = fmt.Sprintf("%s%s", prefix, url)
				}

				// continue if to another domain
				if isExternal := isDifferentDomain(baseURLHost, url); isExternal {
					continue
				}

				// check if in allURLs already
				_, found := allURLs[url]
				if found {
					continue
				}

				// If we get here, url needs to be added to our DSs
				levelURLs = append(levelURLs, url)
				if i == depth {
					continue
				}
				allURLs[url] = true
			}
		}
		// if last iteration, stop - don't add level 4 URLs
		if i == depth {
			break
		}
		urlGraph[i+1] = levelURLs
	}

	// loop over allURLs map, compose Sitemap literal
	sm := Sitemap{
		Protocol: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	for k, _ := range allURLs {
		sm.Sites = append(sm.Sites, URL{
			Location: k,
		})
	}
	out, err := xml.MarshalIndent(sm, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling sitemap into xml: ", err.Error())
	}

	// Print XML
	fmt.Println(string(out))
	os.Exit(0)
}

func isDifferentDomain(ourHost, link string) bool {
	u, err := url.Parse(link)
	if err != nil {
		fmt.Printf("Error parsing URL in check for external domain func: %v: %s", link, err.Error())
	}
	return ourHost != u.Host
}
