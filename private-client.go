package btce

import (
	"bytes"
	"strings"
	"net/url"
	"net/http"
	"crypto/tls"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

const (
	TRADE_API_ENDPOINT = "https://btc-e.com/tapi"
)

type client struct {
	public string
	secret string
}

func NewClient(public string, secret string) (c *client) {
	return &client{public, secret};
}


func (this *client) HMAC_SHA512(secret string, message string) string {
	hash := hmac.New(sha512.New, []byte(secret));
	hash.Write([]byte(message));
	return strings.ToLower(hex.EncodeToString(hash.Sum(nil)));
}

func (this *client) Request(data url.Values) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true },
	};
	httpClient := &http.Client{Transport: tr};
	params := data.Encode();
	request, err := http.NewRequest("POST", TRADE_API_ENDPOINT, bytes.NewBufferString(params));
	if err != nil {
		return nil, err;
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded");
	request.Header.Set("Key", this.public);
	request.Header.Set("Sign", this.HMAC_SHA512(this.secret, params));
	return httpClient.Do(request);
}
