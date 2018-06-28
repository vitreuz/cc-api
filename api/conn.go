package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type conn struct {
	url   string
	token string

	client *http.Client
}

func newConn(url, token string) *conn {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return &conn{url: url, token: token, client: client}
}

func (c conn) get(endpoint string, guid ...string) (*http.Response, error) {
	url := c.url + endpoint

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "bearer "+c.token)

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("invalid status: %s", resp.Status)
	}

	return resp, nil
}
