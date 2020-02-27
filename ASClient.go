package go_aspace

import (
	"net/http"
	"time"
)

type ASClient struct {
	sessionKey string
	rootURL    string
	nclient    *http.Client
}

func NewClient(timeout int) (*ASClient, error) {

	var client *ASClient

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

	client = &ASClient{
		sessionKey: token,
		rootURL:    aspaceRootURL,
		nclient:    nclient,
	}

	return client, nil
}
