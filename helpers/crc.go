package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/twitter-bot-golang/config"
)

//CreateCRCToken for Twitter API
func CreateCRCToken(data string) string {
	consumerSecret := config.GetString("token.consumer_secret_key")
	h := hmac.New(sha256.New, []byte(consumerSecret))
	h.Write([]byte(data))

	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return sha
}
