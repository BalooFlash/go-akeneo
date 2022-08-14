package api

import (
	"io"
	"net/http"
)

type httpClient interface {
	SendRequest(method, uri string, headers map[string]string, body io.Reader) (*http.Response, error)
}

type resourceClient interface {
	GetResource(uri string, uriParams []string) (*http.Response, error)
}
