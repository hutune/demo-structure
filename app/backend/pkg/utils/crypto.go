package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var _IV = []byte("!!SolumVina@2025")

func EncryptString(key string, content string) string {

	k1 := []byte(key)
	data := []byte(content)
	block, _ := aes.NewCipher(k1)
	stream := cipher.NewCFBEncrypter(block, _IV)
	stream.XORKeyStream(data, data)
	return fmt.Sprintf("%x", data)
}

func DecryptString(key string, content string) string {

	bytes, _ := hex.DecodeString(content)
	block, _ := aes.NewCipher([]byte(key))
	stream := cipher.NewCFBDecrypter(block, _IV)

	stream.XORKeyStream(bytes, bytes)

	return string(bytes)
}

func Base64String(content string) string {
	rawDecodedText := base64.StdEncoding.EncodeToString([]byte(content))
	return rawDecodedText
}

func DecodeBase64String(content string) string {
	rawDecodedText, _ := base64.StdEncoding.DecodeString(content)
	return string(rawDecodedText)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
