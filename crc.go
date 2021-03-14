package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

const consumerSecret = "5o7oBg7411Y5TWsIidbarD1LDgMjOuKNZxbh0GdNYhuKKWz9kI"

//CreateCRCToken for Twitter API
func CreateCRCToken(data string) string {
	h := hmac.New(sha256.New, []byte(consumerSecret))
	h.Write([]byte(data))

	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return sha
}
