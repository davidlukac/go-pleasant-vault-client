package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetSecret - Fetch and return Secret from the Vault for given UUID.
func (c Vault) GetSecret(uuid string) Secret {
	log.Printf("Connecting to Pleasant Vault with '%s:%s@%s'.\n", c.Username, ObfuscatePassword(c.Password), c.URL)

	token := c.getToken()
	secret := c.getSecret(token, uuid)

	return secret
}

func (c Vault) getSecret(token string, uuid string) Secret {
	var err error

	secretURL := fmt.Sprintf("%s/api/v5/rest/credential/%s", c.URL, uuid)

	req, err := http.NewRequest(http.MethodGet, secretURL, nil)
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

	var secret Secret
	err = json.Unmarshal(body, &secret)
	if err != nil {
		panic(err)
	}

	return secret
}

func (c Vault) getToken() string {
	var err error

	tokenURL := fmt.Sprintf("%s/OAuth2/Token", c.URL)

	postData := url.Values{}
	postData.Set("grant_type", "password")
	postData.Set("username", c.Username)
	postData.Set("password", c.Password)

	req, err := http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(postData.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("%d - %s", resp.StatusCode, string(body)))
	}

	var tokenResponse tokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		panic(err)
	}

	return tokenResponse.AccessToken
}
