package client

// Vault represents Pleasant Vault instance.
type Vault struct {
	URL      string
	Username string
	Password string
}

// Secret object as returned from Pleasant Vault.
type Secret struct {
	Id               string            `json:"Id"`
	Name             string            `json:"Name"`
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	CustomUserFields map[string]string `json:"CustomUserFields"`
	GroupId          string            `json:"GroupId"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Folder struct {
	Id          string   `json:"Id"`
	Name        string   `json:"Name"`
	ParentId    string   `json:"ParentId"`
	Children    []Folder `json:"Children"`
	Credentials []Secret `json:"Credentials"`
}

type folderResponse struct {
	Id          string   `json:"Id"`
	Name        string   `json:"Name"`
	ParentId    string   `json:"ParentId"`
	Children    []Folder `json:"Children"`
	Credentials []Secret `json:"Credentials"`
}
