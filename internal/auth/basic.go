package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetBasicAuthAPIKey returns the API key from the Authorization header
// Example: Authorization: Basic <api_key>
// Returns: <api_key>
func GetBasicAuthAPIKey(headers http.Header) (string, error) {
	auth_header_val := headers.Get("Authorization")

	if auth_header_val == "" {
		return "", errors.New("no Authorization header found")
	}

	vals := strings.Split(auth_header_val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if vals[0] != "Basic" {
		return "", errors.New("invalid Authorization header")
	}

	return vals[1], nil
}
