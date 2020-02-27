package go_aspace

import (
	"net/http"
	"time"
)

type Client struct {
	session   string
	netclient *http.Client
}

func NewClient(timeout int) (*Client, error) {

	var client *Client

	nclient := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	sessionKey, err := getSessionKey()
	if err != nil {
		return client, err
	}

	client = &Client{
		session:   sessionKey,
		netclient: nclient,
	}

	return client, nil
}
