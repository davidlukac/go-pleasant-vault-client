package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c Vault) CreateEntry(entry *Secret) *Secret {
	token := c.getToken()
	newEntryId := c.createEntry(token, entry)
	newEntry := c.getSecret(token, newEntryId)

	return &newEntry
}

func (c Vault) PatchEntry(id string, jsonPatch string) {
	token := c.getToken()
	c.patchEntry(token, id, jsonPatch)
}

// GetPassword return password string for given Entry ID.
func (c Vault) GetPassword(id string) string {
	token := c.getToken()
	password := c.getPassword(token, id)

	return password
}

// GetEntryWithPassword returns an Entry object for given ID, enriched by the password (which is by
// default returned empty).
func (c Vault) GetEntryWithPassword(uuid string) *Secret {
	entry := c.GetSecret(uuid)
	entry.Password = c.GetPassword(uuid)

	return &entry
}

// POST    /api/v5/rest/entries
func (c Vault) createEntry(token string, entry *Secret) string {
	var err error

	rootFolderURL := fmt.Sprintf("%s/api/v5/rest/entries", c.URL)

	entryJson, err := json.Marshal(&entry)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, rootFolderURL, bytes.NewBuffer(entryJson))
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

// PATCH    /api/v5/rest/entries/{entryId:guid}
func (c Vault) patchEntry(token string, id string, jsonPatch string) {
	var err error

	entryUrl := fmt.Sprintf("%s/api/v5/rest/entries/%s", c.URL, id)

	req, err := http.NewRequest(http.MethodPatch, entryUrl, bytes.NewBuffer([]byte(jsonPatch)))
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

	if resp.StatusCode != http.StatusNoContent {
		panic(fmt.Sprintf("%d - %s", resp.StatusCode, string(body)))
	}
}

// getPassword returns password string for given Entry ID.
// GET    /api/v5/rest/entries/{entryId:guid}/password
func (c Vault) getPassword(token string, id string) string {
	var err error

	passwordUrl := fmt.Sprintf("%s/api/v5/rest/entries/%s/password", c.URL, id)

	req, err := http.NewRequest(http.MethodGet, passwordUrl, nil)
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

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}()

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
