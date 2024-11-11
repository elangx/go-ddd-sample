package util

import (
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

func ScryptWithSalt(data string, salt string) (string, error) {
	r, err := scrypt.Key([]byte(data), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(r), nil
}
