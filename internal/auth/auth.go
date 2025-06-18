package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIkey extract the API key from the headers of the HTTP request.
// example usage:
//authorization: apikey{insert your API key here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing Authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header format")
	}
	if vals[0] != "apikey" {
		return "", errors.New("invalid Authorization header prefix, expected 'apikey'")
	}
	return vals[1], nil
}
