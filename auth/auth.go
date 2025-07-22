package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization key: ApiKey {insert apikey here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("no Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("no Authorization header")
	}

	return vals[1], nil
}
