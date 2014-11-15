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
	PRIVATE_API_ENDPOINT = "https://btc-e.com/tapi"
)

type clientTLS struct {
	public string
	secret string
	hmac bool
}

func NewClient(public string, secret string, hmac bool) *clientTLS {
	return &clientTLS{public, secret, hmac};
}


func (this *clientTLS) HMAC_SHA512(secret string, message string) string {
	hash := hmac.New(sha512.New, []byte(secret));
	hash.Write([]byte(message));
	return strings.ToLower(hex.EncodeToString(hash.Sum(nil)));
}

func (this *clientTLS) Request(data url.Values, location string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true },
	};
	httpClient := &http.Client{Transport: tr};
	params := data.Encode();
	if this.hmac == true {
		request, err := http.NewRequest("POST", PRIVATE_API_ENDPOINT, bytes.NewBufferString(params));
		if err != nil {
			return nil, err;
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded");
		request.Header.Set("Key", this.public);
		request.Header.Set("Sign", this.HMAC_SHA512(this.secret, params));
		return httpClient.Do(request);
	} else {
		request, err := http.NewRequest("GET", location, bytes.NewBufferString(params));
		if err != nil {
			return nil, err;
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded");
		return httpClient.Do(request);
	}
}
