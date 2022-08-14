package goakeneo

import (
	"github.com/go-akeneo/go-akeneo/api"
	"github.com/go-akeneo/go-akeneo/security"
)

type PimClient struct {
	Auth       security.Auth
	ProductApi api.ProductApi
}

func NewPimClient(auth security.Auth, productApi api.ProductApi) PimClient {
	return PimClient{Auth: auth, ProductApi: productApi}
}

func (p PimClient) GetProductApi() api.ProductApi {
	return p.ProductApi
}
