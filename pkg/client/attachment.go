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

// GetAttachment returns specific attachment object for given attachment and entry IDs.
func (c Vault) GetAttachment(entryID string, attachmentID string) *Attachment {
	token := c.getToken()
	attachment := c.getAttachment(token, entryID, attachmentID)

	return attachment
}

// GetAttachments provides list of attachment objects for given Entry ID.
func (c Vault) GetAttachments(entryID string) []Attachment {
	token := c.getToken()
	attachments := c.getAttachments(token, entryID)

	return attachments
}

// CreateAttachment - upload of a new Attachment to given Entry identified by ID.
func (c Vault) CreateAttachment(entryID string, attachment Attachment) string {
	token := c.getToken()
	attachmentId := c.createAttachment(token, entryID, attachment)

	return attachmentId
}

// GET        <url>/api/v5/rest/entries/{entryId:guid}/attachments/{attachmentId:guid}
func (c Vault) getAttachment(token string, entryID string, attachmentID string) *Attachment {
	var err error

	attachmentURL := fmt.Sprintf("%s/api/v5/rest/entries/%s/attachments/%s", c.URL, entryID, attachmentID)

	req, err := http.NewRequest(http.MethodGet, attachmentURL, nil)
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

	var attachment Attachment
	err = json.Unmarshal(body, &attachment)
	if err != nil {
		panic(err)
	}

	return &attachment
}

// GET        <url>/api/v5/rest/entries/{entryId:guid}/attachments
func (c Vault) getAttachments(token string, id string) []Attachment {
	var err error

	attachmentsURL := fmt.Sprintf("%s/api/v5/rest/entries/%s/attachments", c.URL, id)

	req, err := http.NewRequest(http.MethodGet, attachmentsURL, nil)
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

	var attachments []Attachment
	err = json.Unmarshal(body, &attachments)
	if err != nil {
		panic(err)
	}

	return attachments
}

// POST        <url>/api/v5/rest/entries/{entryId:guid}/attachments
func (c Vault) createAttachment(token string, entryID string, attachment Attachment) string {
	var err error

	attachmentURL := fmt.Sprintf("%s/api/v5/rest/entries/%s/attachments", c.URL, entryID)

	attachmentJSON, err := json.Marshal(&attachment)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, attachmentURL, bytes.NewBuffer(attachmentJSON))
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

	attachmentID := strings.Trim(string(body), "\"")

	return attachmentID
}
