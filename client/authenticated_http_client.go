package client

import (
	"fmt"
	"github.com/go-akeneo/go-akeneo/api"
	"io"
	"net/http"

	"github.com/go-akeneo/go-akeneo/security"
)

type authenticationApi interface {
	Authenticate(clientId, secret string, requestBody io.Reader) (api.AuthenticationResponse, error)
	AuthenticateByPassword(clientId, secret, username, password string) (api.AuthenticationResponse, error)
	AuthenticateByRefreshToken(clientId, secret, refreshToken string) (api.AuthenticationResponse, error)
}

type AuthenticatedHttpClient struct {
	httpClient        HttpClient
	auth              security.Auth
	authenticationApi authenticationApi
}

func NewAuthenticatedHttpClient(httpClient HttpClient, auth security.Auth, authenticationApi authenticationApi) AuthenticatedHttpClient {
	return AuthenticatedHttpClient{
		httpClient:        httpClient,
		auth:              auth,
		authenticationApi: authenticationApi,
	}
}

func (a AuthenticatedHttpClient) SendRequest(method, uri string, headers map[string]string, body io.Reader) (*http.Response, error) {
	if a.auth.AccessToken == "" {
		ar, err := a.authenticationApi.AuthenticateByPassword(a.auth.ClientId, a.auth.Secret, a.auth.Username, a.auth.Password)
		if err != nil {
			return nil, err
		}
		a.auth.AccessToken = ar.AccessToken
		a.auth.RefreshToken = ar.RefreshToken
	}

	headers["Authorization"] = fmt.Sprintf("Bearer %s", a.auth.AccessToken)
	res, err := a.httpClient.SendRequest(method, uri, headers, body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusUnauthorized {
		ar, err := a.renewTokens()
		if err != nil {
			return nil, err
		}
		a.auth.AccessToken = ar.AccessToken
		a.auth.RefreshToken = ar.RefreshToken

		headers["Authorization"] = fmt.Sprintf("Bearer %s", a.auth.AccessToken)
		res, err = a.httpClient.SendRequest(method, uri, headers, body)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (a AuthenticatedHttpClient) renewTokens() (api.AuthenticationResponse, error) {
	return a.authenticationApi.AuthenticateByRefreshToken(a.auth.ClientId, a.auth.Secret, a.auth.RefreshToken)
}
