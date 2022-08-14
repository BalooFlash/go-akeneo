package security

type Auth struct {
	ClientId     string
	Secret       string
	Username     string
	Password     string
	AccessToken  string
	RefreshToken string
}

func NewFromPassword(clientId, secret, username, password string) Auth {
	return Auth{
		ClientId: clientId,
		Secret:   secret,
		Username: username,
		Password: password,
	}
}
