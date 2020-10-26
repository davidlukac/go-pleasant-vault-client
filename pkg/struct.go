package pkg

type Secret struct {
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	CustomUserFields map[string]string `json:"CustomUserFields"`
}

type Pleasant struct {
	Url      string
	Username string
	Password string
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
