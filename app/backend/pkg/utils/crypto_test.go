package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecryptString(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		content string
	}{
		{
			name:    "Normal string",
			key:     "0123456789abcdef",
			content: "hello world",
		},
		{
			name:    "Empty string",
			key:     "0123456789abcdef",
			content: "",
		},
		{
			name:    "Special characters",
			key:     "0123456789abcdef",
			content: "!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted := EncryptString(tt.key, tt.content)
			decrypted := DecryptString(tt.key, encrypted)
			assert.Equal(t, tt.content, decrypted)
		})
	}
}

func TestBase64StringOperations(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "Normal string",
			content: "hello world",
		},
		{
			name:    "Empty string",
			content: "",
		},
		{
			name:    "Special characters",
			content: "!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := Base64String(tt.content)
			decoded := DecodeBase64String(encoded)
			assert.Equal(t, tt.content, decoded)
		})
	}
}

func TestPasswordHashing(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "Normal password",
			password: "myPassword123",
		},
		{
			name:     "Empty password",
			password: "",
		},
		{
			name:     "Complex password",
			password: "P@ssw0rd!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			assert.NoError(t, err)
			assert.NotEmpty(t, hash)

			// Verify password comparison
			isValid := ComparePassword(tt.password, hash)
			assert.True(t, isValid)

			// Verify wrong password fails
			isValid = ComparePassword("wrongpassword", hash)
			assert.False(t, isValid)
		})
	}
}
