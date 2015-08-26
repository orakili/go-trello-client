package trello

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
)

type WebHook struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	IdModel     string `json:"idModel"`
	CallbackURL string `json:"callbackURL"`
	Active      bool   `json:"active"`
}

// Convert to "binary" to be compatible with the way node.js handles crypto.
// @See https://trello.com/c/IMVYmVG1/180-webhooks (Doug Patti, Mar 7 at 1:59 AM).
func toBinaryString(b []byte) []byte {
	l := utf8.RuneCount(b)
	buf := make([]byte, l)
	for i := 0; i < l; i++ {
		r, size := utf8.DecodeRune(b)
		b = b[size:]
		buf[i] = uint8(r & 0xff)
	}
	return buf
}

// Compute the base64 digest of the message.
func base64Digest(message []byte, key []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(message)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Validate the call the webhook.
func CheckWebhook(callbackURL string, secret string, r *http.Request) (bool, error) {
	if r.ContentLength <= 0 || r.Body == nil {
		return false, errors.New("Webhook: Invalid request body")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, err
	} else if len(body) == 0 {
		return false, errors.New("Webhook: Empty body")
	}

	header := r.Header.Get("X-Trello-Webhook")

	key := []byte(secret)
	message := append(body, []byte(callbackURL)...)

	// Double-HMAC to blind any timing channel attacks
	// https://www.isecpartners.com/blog/2011/february/double-hmac-verification.asp
	singleHash := base64Digest(toBinaryString(message), key)
	doubleHash := base64Digest([]byte(singleHash), key)
	headerHash := base64Digest([]byte(header), key)

	if doubleHash == headerHash {
		return true, nil
	} else {
		return false, nil
	}
}
