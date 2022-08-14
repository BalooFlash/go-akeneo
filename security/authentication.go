package security

type Authentication struct {
	ClientId     string
	Secret       string
	Username     string
	Password     string
	AccessToken  string
	RefreshToken string
}

func NewFromPassword(clientId, secret, username, password string) Authentication {
	return Authentication{
		ClientId: clientId,
		Secret:   secret,
		Username: username,
		Password: password,
	}
}
