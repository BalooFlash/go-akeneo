package routing

import (
	"fmt"
	"strings"
)

type UriGenerator struct {
	baseUri string
}

func NewUriGenerator(baseUri string) UriGenerator {
	return UriGenerator{baseUri: strings.TrimSuffix(baseUri, "/")}
}

func (u UriGenerator) Generate(path string, uriParams []string) string {
	uri := u.baseUri + "/" + fmt.Sprintf(strings.TrimPrefix(path, "/"), u.convertSlice(uriParams)...)
	return uri
}

func (u UriGenerator) convertSlice(s []string) []interface{} {
	var i []interface{}
	for _, e := range s {
		i = append(i, e)
	}
	return i
}
