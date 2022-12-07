package http_client

import (
	"crypto/tls"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func NewClient(insecureSkipVerify bool) HTTPClient {
	client := &http.Client{Timeout: 2 * time.Second}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}

	return client
}
