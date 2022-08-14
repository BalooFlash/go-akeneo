package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type httpClient interface {
	SendRequest(method, uri string, headers map[string]string, body io.Reader) (*http.Response, error)
}

type resourceClient interface {
	GetResource(uri string, uriParams []string) (*http.Response, error)
	CreateResource(uri string, uriParams []string, body json.RawMessage) (*http.Response, error)
	DeleteResource(uri string, uriParams []string) error
}
