package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts an api key from
//headers of an http request
//Example:
//Authorization : ApiKey{insert apikey here}

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malinformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malinformed first Part of auth header")
	}
	return vals[1], nil
}
