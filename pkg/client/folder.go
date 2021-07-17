package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetRootFolder returns ID of the root folder.
func (c Vault) GetRootFolder() string {
	token := c.getToken()
	return c.getRootFolder(token)
}

// GetFolder returns Folder object for given ID.
func (c Vault) GetFolder(folderID string) *Folder {
	token := c.getToken()
	response := c.getFolder(token, folderID)

	folder := Folder{}
	folder.Id = response.Id
	folder.Name = response.Name
	folder.ParentId = response.ParentId
	folder.Children = response.Children
	folder.Credentials = response.Credentials

	return &folder
}

// CreateFolder creates new Folder from provided object and returns updated self.
// POST    /api/v5/rest/folders
func (c Vault) CreateFolder(folder *Folder) *Folder {
	token := c.getToken()
	newFolderID := c.createFolder(token, folder)
	folder = c.GetFolder(newFolderID)

	return folder
}

// POST    /api/v5/rest/folders
func (c Vault) createFolder(token string, folder *Folder) string {
	var err error

	rootFolderURL := fmt.Sprintf("%s/api/v5/rest/folders", c.URL)

	folderJSON, err := json.Marshal(&folder)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, rootFolderURL, bytes.NewBuffer(folderJSON))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("%d - %s", resp.StatusCode, string(body)))
	}

	bodyString := strings.Trim(string(body), "\"")

	return bodyString
}

// GET    /api/v5/rest/folders/root
func (c Vault) getRootFolder(token string) string {
	var err error

	rootFolderURL := fmt.Sprintf("%s/api/v5/rest/folders/root", c.URL)

	req, err := http.NewRequest(http.MethodGet, rootFolderURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("%d - %s", resp.StatusCode, string(body)))
	}

	bodyString := strings.Trim(string(body), "\"")

	return bodyString
}

// GET    /api/v5/rest/folders/{folderId:guid}
func (c Vault) getFolder(token string, folderID string) *folderResponse {
	var err error

	rootFolderURL := fmt.Sprintf("%s/api/v5/rest/folders/%s", c.URL, folderID)

	req, err := http.NewRequest(http.MethodGet, rootFolderURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("%d - %s", resp.StatusCode, string(body)))
	}

	var folderResponse folderResponse
	err = json.Unmarshal(body, &folderResponse)
	if err != nil {
		panic(err)
	}

	return &folderResponse
}
