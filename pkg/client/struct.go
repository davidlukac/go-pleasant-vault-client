package client

// Vault represents Pleasant Vault instance.
type Vault struct {
	URL      string
	Username string
	Password string
}

// Secret object as returned from Pleasant Vault.
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
