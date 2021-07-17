package client

// Vault represents Pleasant Vault instance.
type Vault struct {
	URL      string
	Username string
	Password string
}

// Secret object as returned from Pleasant Vault.
type Secret struct {
	ID               string            `json:"Id"`
	Name             string            `json:"Name"`
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	CustomUserFields map[string]string `json:"CustomUserFields"`
	GroupID          string            `json:"GroupId"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Folder object.
type Folder struct {
	ID          string   `json:"Id"`
	Name        string   `json:"Name"`
	ParentID    string   `json:"ParentId"`
	Children    []Folder `json:"Children"`
	Credentials []Secret `json:"Credentials"`
}

type folderResponse struct {
	ID          string   `json:"Id"`
	Name        string   `json:"Name"`
	ParentID    string   `json:"ParentId"`
	Children    []Folder `json:"Children"`
	Credentials []Secret `json:"Credentials"`
}
