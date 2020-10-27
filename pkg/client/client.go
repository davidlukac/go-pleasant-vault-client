package client

import (
	"encoding/json"
	"fmt"
	. "github.com/davidlukac/go-pleasant-vault-client/internal"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetSecret(uuid string) Secret {
	pleasant := loadAndValidateEnv()
	log.Printf("Connecting to Pleasant with '%s:%s@%s'.\n", pleasant.Username, ObfuscatePassword(pleasant.Password), pleasant.Url)

	token := getToken(pleasant)
	secret := getSecret(pleasant, token, uuid)

	return secret
}

func getSecret(pleasant Pleasant, token string, uuid string) Secret {
	var err error

	secretUrl := fmt.Sprintf("%s/api/v4/rest/credential/%s", pleasant.Url, uuid)

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

func getToken(pleasant Pleasant) string {
	var err error

	tokenUrl := fmt.Sprintf("%s/OAuth2/Token", pleasant.Url)

	postData := url.Values{}
	postData.Set("grant_type", "password")
	postData.Set("username", pleasant.Username)
	postData.Set("password", pleasant.Password)

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

	defer resp.Body.Close()

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

func loadAndValidateEnv() Pleasant {
	err := godotenv.Overload(".env", ".local.env")

	if err != nil {
		panic(err)
	}

	var pleasant Pleasant

	pleasant.Url = strings.TrimSpace(os.Getenv(PleasantUrlVar))
	pleasant.Username = strings.TrimSpace(os.Getenv(PleasantUsernameVar))
	pleasant.Password = strings.TrimSpace(os.Getenv(PleasantPasswordVar))

	if strings.TrimSpace(pleasant.Url) == "" {
		panic(fmt.Sprintf("Environment variable '%s' must be set!", PleasantUrlVar))
	}

	if strings.TrimSpace(pleasant.Username) == "" {
		panic(fmt.Sprintf("Environment variable '%s' must be set!", PleasantUrlVar))
	}

	if strings.TrimSpace(pleasant.Password) == "" {
		panic(fmt.Sprintf("Environment variable '%s' must be set!", PleasantUrlVar))
	}

	return pleasant
}
