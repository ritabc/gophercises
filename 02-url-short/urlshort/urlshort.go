package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		if long, ok := pathsToUrls[req.URL.Path]; ok {
			http.Redirect(resWriter, req, long, http.StatusFound)
		} else {
			fallback.ServeHTTP(resWriter, req)
		}
	}
}

type pathURL struct {
	Path string
	URL  string
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsedYML []pathURL
	err := yaml.UnmarshalStrict(yml, &parsedYML)
	if err != nil {
		return nil, err
	}
	pathMap := make(map[string]string)
	for _, link := range parsedYML {
		pathMap[link.Path] = link.URL
	}
	return MapHandler(pathMap, fallback), nil
}
