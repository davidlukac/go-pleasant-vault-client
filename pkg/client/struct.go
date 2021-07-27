package client

// Vault represents Pleasant Vault instance.
type Vault struct {
	URL      string
	Username string
	Password string
}

// Tag in an Entry (secret).
type Tag struct {
	Name string `json:"Name"`
}

// Secret object as returned from Pleasant Vault.
type Secret struct {
	ID               string            `json:"Id"`
	Name             string            `json:"Name"`
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	GroupID          string            `json:"GroupId"`
	Url              string            `json:"Url"`
	CustomUserFields map[string]string `json:"CustomUserFields"`
	Tags             []Tag             `json:"Tags"`
	Notes            string            `json:"Notes"`
	ExpirationDate   string            `json:"Expires"`
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

// Attachment - file attachment to an Entry.
type Attachment struct {
	ID       string `json:"AttachmentId"`
	EntryID  string `json:"CredentialObjectId"`
	FileName string `json:"FileName"`
	FileData string `json:"FileData"`
	FileSize int    `json:"FileSize"`
}
