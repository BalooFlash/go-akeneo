package goakeneo

import (
	"github.com/go-akeneo/go-akeneo/api"
	"github.com/go-akeneo/go-akeneo/client"
	"github.com/go-akeneo/go-akeneo/routing"
	"github.com/go-akeneo/go-akeneo/security"
)

type ClientBuilder struct {
	baseUrl string
}

func NewClientBuilder(baseUrl string) ClientBuilder {
	return ClientBuilder{baseUrl: baseUrl}
}

func (c ClientBuilder) BuildAuthenticatedByPassword(clientId, secret, username, password string) PimClient {
	authentication := security.NewFromPassword(clientId, secret, username, password)

	return c.BuildAuthenticatedClient(authentication)
}

func (c ClientBuilder) BuildAuthenticatedClient(authentication security.Auth) PimClient {
	uriGenerator := routing.NewUriGenerator(c.baseUrl)

	httpClient := client.NewHttpClient()
	authApi := api.NewAuthenticationApi(httpClient, uriGenerator)
	authHttpClient := client.NewAuthenticatedHttpClient(httpClient, authentication, authApi)
	resourceClient := client.NewResourceClient(authHttpClient, uriGenerator)

	return NewPimClient(
		authentication,
		api.NewProductApi(resourceClient),
	)
}
