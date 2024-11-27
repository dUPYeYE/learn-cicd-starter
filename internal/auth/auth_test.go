package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "No authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer token"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Correct authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey correctapikey"},
			},
			expectedKey:   "correctapikey",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectedKey {
				t.Errorf("expected key %v, got %v", tt.expectedKey, key)
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}
