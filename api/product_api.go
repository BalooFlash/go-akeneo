package api

import (
	"net/http"
)

type ProductApi struct {
	resourceClient resourceClient
}

const ProductUri = "api/rest/v1/products/%s"

func NewProductApi(resourceClient resourceClient) ProductApi {
	return ProductApi{resourceClient: resourceClient}
}

func (a ProductApi) Get(code string) (*http.Response, error) {
	return a.resourceClient.GetResource(ProductUri, []string{code})
}
