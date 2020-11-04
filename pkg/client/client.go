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

func (c Vault) GetSecret(uuid string) Secret {
	log.Printf("Connecting to Pleasant Vault with '%s:%s@%s'.\n", c.Username, ObfuscatePassword(c.Password), c.Url)

	token := c.getToken()
	secret := c.getSecret(token, uuid)

	return secret
}

func (c Vault) getSecret(token string, uuid string) Secret {
	var err error

	secretUrl := fmt.Sprintf("%s/api/v4/rest/credential/%s", c.Url, uuid)

	req, err := http.NewRequest(http.MethodGet, secretUrl, nil)
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

	tokenUrl := fmt.Sprintf("%s/OAuth2/Token", c.Url)

	postData := url.Values{}
	postData.Set("grant_type", "password")
	postData.Set("username", c.Username)
	postData.Set("password", c.Password)

	req, err := http.NewRequest(http.MethodPost, tokenUrl, strings.NewReader(postData.Encode()))
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
		log.Println(err)
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
