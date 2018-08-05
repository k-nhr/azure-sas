package sas

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

func Generate(uri, key string, expiry int64) (string, error) {
	encURI := url.QueryEscape(uri)
	ttl := time.Now().Unix() + expiry
	signKey := fmt.Sprintf("%s\n%d", encURI, ttl)

	decKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	hash := hmac.New(sha256.New, decKey)
	if _, err := hash.Write([]byte(signKey)); err != nil {
		return "", err
	}
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(hash.Sum(nil)))

	token := fmt.Sprintf("sr=%s&sig=%s&se=%d", encURI, signature, ttl)

	return "SharedAccessSignature " + token, nil
}
