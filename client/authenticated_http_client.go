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
	authentication    security.Authentication
	authenticationApi authenticationApi
}

func NewAuthenticatedHttpClient(httpClient HttpClient, auth security.Authentication, authenticationApi authenticationApi) AuthenticatedHttpClient {
	return AuthenticatedHttpClient{
		httpClient:        httpClient,
		authentication:    auth,
		authenticationApi: authenticationApi,
	}
}

func (a AuthenticatedHttpClient) SendRequest(method, uri string, headers map[string]string, body io.Reader) (*http.Response, error) {
	if a.authentication.AccessToken == "" {
		ar, err := a.authenticationApi.AuthenticateByPassword(a.authentication.ClientId, a.authentication.Secret, a.authentication.Username, a.authentication.Password)
		if err != nil {
			return nil, err
		}
		a.authentication.AccessToken = ar.AccessToken
		a.authentication.RefreshToken = ar.RefreshToken
	}

	headers["Authorization"] = fmt.Sprintf("Bearer %s", a.authentication.AccessToken)
	res, err := a.httpClient.SendRequest(method, uri, headers, body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusUnauthorized {
		ar, err := a.renewTokens()
		if err != nil {
			return nil, err
		}
		a.authentication.AccessToken = ar.AccessToken
		a.authentication.RefreshToken = ar.RefreshToken

		headers["Authorization"] = fmt.Sprintf("Bearer %s", a.authentication.AccessToken)
		res, err = a.httpClient.SendRequest(method, uri, headers, body)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (a AuthenticatedHttpClient) renewTokens() (api.AuthenticationResponse, error) {
	return a.authenticationApi.AuthenticateByRefreshToken(a.authentication.ClientId, a.authentication.Secret, a.authentication.RefreshToken)
}
