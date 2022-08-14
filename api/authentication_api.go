package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/go-akeneo/go-akeneo/routing"
	"io"
	"io/ioutil"
)

type AuthenticationApi struct {
	client       httpClient
	uriGenerator routing.UriGenerator
}

type passwordRequestBody struct {
	Grant    string `json:"grant_type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type refreshRequestBody struct {
	Grant        string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type AuthenticationResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

const TokenUri = "api/oauth/v1/token"

func NewAuthenticationApi(client httpClient, uriGenerator routing.UriGenerator) AuthenticationApi {
	return AuthenticationApi{client: client, uriGenerator: uriGenerator}
}

func (a AuthenticationApi) AuthenticateByPassword(clientId, secret, username, password string) (AuthenticationResponse, error) {
	requestBody := passwordRequestBody{
		Grant:    "password",
		Username: username,
		Password: password,
	}

	b, _ := json.Marshal(requestBody)

	return a.Authenticate(clientId, secret, bytes.NewReader(b))
}

func (a AuthenticationApi) AuthenticateByRefreshToken(clientId, secret, refreshToken string) (AuthenticationResponse, error) {
	requestBody := refreshRequestBody{
		Grant:        "refresh_token",
		RefreshToken: refreshToken,
	}

	b, _ := json.Marshal(requestBody)

	return a.Authenticate(clientId, secret, bytes.NewReader(b))
}

func (a AuthenticationApi) Authenticate(clientId, secret string, requestBody io.Reader) (AuthenticationResponse, error) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(clientId+":"+secret)),
	}

	uri := a.uriGenerator.Generate(TokenUri, nil)
	ar := AuthenticationResponse{}

	res, err := a.client.SendRequest("POST", uri, headers, requestBody)
	if err != nil {
		return ar, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ar, err
	}

	if err = json.Unmarshal(resBody, &ar); err != nil {
		return ar, err
	}

	return ar, nil
}
