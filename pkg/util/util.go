package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func CalculateSha256(reader io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
