package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API key from
// the headers of an HTTP requests
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of headers")
	}

	return vals[1], nil
}
