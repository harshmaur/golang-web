package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// Tampered checks if the data is tampered with or not.
func Tampered(data string) bool {
	xs := DecodeSplit(data)
	uPics := xs[1]
	uCode := xs[2]
	return uCode != GetCode(uPics)
}

// GetCode returns an HMAC string on the basis of data supplied.
func GetCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	return fmt.Sprintf("%x", h.Sum(nil))
}
