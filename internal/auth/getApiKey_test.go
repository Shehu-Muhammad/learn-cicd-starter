package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey abc123")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if apiKey != "abc123" {
		t.Errorf("Expected apiKey to be abc123, but got: %s", apiKey)
	}

}
