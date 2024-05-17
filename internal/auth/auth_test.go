package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case: Valid Authorization header
	t.Run("Valid Authorization header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey someapikey")
		expectedAPIKey := "someapikey"
		expectedError := error(nil)

		apiKey, err := GetAPIKey(headers)

		if apiKey != expectedAPIKey {
			t.Errorf("expected API key %q, got %q", expectedAPIKey, apiKey)
		}
		if !errors.Is(err, expectedError) {
			t.Errorf("expected error %v, got %v", expectedError, err)
		}
	})
}