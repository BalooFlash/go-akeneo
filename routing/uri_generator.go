package routing

import (
	"fmt"
	"net/url"
	"strings"
)

type UriGenerator struct {
	baseUri string
}

func NewUriGenerator(baseUri string) UriGenerator {
	return UriGenerator{baseUri: strings.TrimSuffix(baseUri, "/")}
}

func (u UriGenerator) Generate(path string, uriParams []string, queryParams map[string]string) string {
	uri := u.baseUri + "/" + fmt.Sprintf(strings.TrimPrefix(path, "/"), u.convertSlice(uriParams)...)

	if queryParams != nil {
		uri += "?" + u.buildQueryString(queryParams)
	}

	return uri
}

func (u UriGenerator) convertSlice(s []string) []interface{} {
	var i []interface{}
	for _, e := range s {
		i = append(i, e)
	}
	return i
}

func (u UriGenerator) buildQueryString(queryParams map[string]string) string {
	val := url.Values{}
	for k, v := range queryParams {
		val.Add(k, v)
	}
	return val.Encode()
}
