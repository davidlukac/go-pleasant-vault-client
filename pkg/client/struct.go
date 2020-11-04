package client

type Vault struct {
	Url      string
	Username string
	Password string
}

type Secret struct {
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	CustomUserFields map[string]string `json:"CustomUserFields"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
