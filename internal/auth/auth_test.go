package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name        string
		value       string
		shouldError bool
	}{
		{name: "Authorization", value: "ApiKey test123", shouldError: true},
		{name: "Test", value: "ApiKey test123", shouldError: true},
		{name: "Authorization", value: " test123", shouldError: true},
	}

	for _, c := range cases {
		header := http.Header{}
		header.Add(c.name, c.value)
		key, err := GetAPIKey(header)
		if (err == nil) == c.shouldError {
			t.Errorf("\nIsErr: %v\nShouldErr: %v\nc.value: %v\nkey: %v", err == nil, c.shouldError, c.value, key)
		}
	}
}
