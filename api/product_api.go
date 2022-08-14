package api

import (
	"encoding/json"
	"github.com/go-akeneo/go-akeneo/api/model"
	"net/http"
)

type ProductApi struct {
	resourceClient resourceClient
}

const ProductsUri = "api/rest/v1/products"
const ProductUri = "api/rest/v1/products/%s"

func NewProductApi(resourceClient resourceClient) ProductApi {
	return ProductApi{resourceClient: resourceClient}
}

func (a ProductApi) Get(code string) (*http.Response, error) {
	return a.resourceClient.GetResource(ProductUri, []string{code})
}

func (a ProductApi) Create(code string, product model.Product) (*http.Response, error) {
	product.Identifier = code

	j, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	return a.resourceClient.CreateResource(ProductsUri, []string{}, j)
}

func (a ProductApi) Delete(code string) error {
	return a.resourceClient.DeleteResource(ProductUri, []string{code});
}
