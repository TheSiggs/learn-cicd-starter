package auth

import (
	"net/http"
	"testing"
)

func TestAuthOnly(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "boobs")
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fail()
	}
}

func TestCorrectKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey boobs")
	_, err := GetAPIKey(headers)
	if err != nil {
		t.Fail()
	}
}
