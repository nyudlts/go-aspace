package go_aspace

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	sessionKey string
	rootURL    string
	nclient    *http.Client
}

func NewClient(timeout int) (*Client, error) {

	var client *Client

	nclient := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	token, err := getSessionKey()
	if err != nil {
		return client, err
	}

	aspaceRootURL, err := GetRootURL()
	if err != nil {
		return client, err
	}

	client = &Client{
		sessionKey: token,
		rootURL:    aspaceRootURL,
		nclient:    nclient,
	}

	return client, nil
}

func ASGet(client *Client, endpoint string, authenticated bool) (*http.Response, error) {

	var response *http.Response
	if authenticated {
		return response, nil
	} else {
		url := client.rootURL + endpoint

		response, err := client.nclient.Get(url)

		if err != nil {
			return response, err
		}

		if response.StatusCode != 200 {
			return response, fmt.Errorf("Did not return a 200 while authenticating, recieved a %d", response.StatusCode)
		}

		return response, nil
	}
}
