package goakeneo

import (
	"github.com/go-akeneo/go-akeneo/api"
	"github.com/go-akeneo/go-akeneo/security"
)

type PimClient struct {
	Auth       security.Authentication
	ProductApi api.ProductApi
}

func NewPimClient(auth security.Authentication, productApi api.ProductApi) PimClient {
	return PimClient{Auth: auth, ProductApi: productApi}
}

func (p PimClient) GetProductApi() api.ProductApi {
	return p.ProductApi
}
