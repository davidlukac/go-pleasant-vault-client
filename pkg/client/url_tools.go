package client

import (
	"fmt"
	"net/url"
)

func ParseUUID(s string, queryKey string) string {
	if queryKey == "" {
		queryKey = "itemId"
	}

	// Try parsing UUID string as URL.
	u, err := url.ParseRequestURI(s)
	if err != nil {
		// Returning provided UUID as it came, as it's not an ULR.
		return s
	}

	u, err = url.Parse(s)
	if err != nil {
		panic(err)
	}

	query := u.Query()
	uuidQuery := query.Get(queryKey)

	if uuidQuery == "" {
		panic(fmt.Sprintf("'%s' is an URL but doesn't contain '%s'!", s, queryKey))
	}

	return uuidQuery
}
