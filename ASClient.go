package aspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"time"
)

type ASClient struct {
	sessionKey string
	rootURL    string
	nclient    *http.Client
}

type Creds struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewClient(configFile string, environment string, timeout int) (*ASClient, error) {

	var client *ASClient
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return client, fmt.Errorf("Configuration file %s does not exist", configFile)
	}

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return client, err
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    time.Duration(timeout) * time.Second,
		DisableCompression: true,
	}

	nclient := &http.Client{
		Transport: tr,
	}

	creds, err := getCreds(environment, bytes)
	if err != nil {
		return client, err
	}

	token, err := getSessionKey(creds)
	if err != nil {
		return client, err
	}

	client = &ASClient{
		sessionKey: token,
		rootURL:    creds.URL,
		nclient:    nclient,
	}

	return client, err
}

func getCreds(environment string, configBytes []byte) (Creds, error) {
	credsMap := map[string]Creds{}

	err := yaml.Unmarshal(configBytes, &credsMap)
	if err != nil {
		return Creds{}, err
	}

	for k, v := range credsMap {
		if environment == k {
			return v, nil
		}
	}

	return Creds{}, fmt.Errorf("Credentials file did not contain %s\n", environment)
}

func getSessionKey(creds Creds) (string, error) {

	url := fmt.Sprintf("%s/users/%s/login?password=%s", creds.URL, creds.Username, creds.Password)

	response, err := http.Post(url, "text/json", nil)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("Did not return a 200 while authenticating, recieved a %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		return "", err
	}
	sessionKey := fmt.Sprintf("%v", responseMap["session"])

	// --> TODO the session key should be validated against a regex
	if sessionKey != "" {
		return sessionKey, nil
	} else {
		return sessionKey, fmt.Errorf("Session field was empty")
	}
}

func (a *ASClient) GetSessionKey() string {
	return a.sessionKey
}

func (a *ASClient) do(request *http.Request, authenticated bool) (*http.Response, error) {
	var response *http.Response

	if authenticated {
		request.Header.Add("X-ArchivesSpace-Session", a.sessionKey)
	}

	response, err := a.nclient.Do(request)
	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		return response, fmt.Errorf("ArchivesSpace responded with a non-200:\nstatus-code: %d\n%s", response.StatusCode, string(body))
	}
	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *ASClient) get(endpoint string, authenticated bool) (*http.Response, error) {
	var response *http.Response
	url := a.rootURL + endpoint

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}

	response, err = a.do(request, authenticated)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *ASClient) post(endpoint string, authenticated bool, body string) (*http.Response, error) {
	var response *http.Response
	url := a.rootURL + endpoint
	request, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		return response, err
	}

	response, err = a.do(request, authenticated)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *ASClient) delete(endpoint string) (*http.Response, error) {
	var response *http.Response
	url := a.rootURL + endpoint
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return response, err
	}
	response, err = a.do(request, true)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *ASClient) JsonRequest(endpoint string, method string, body string) (*http.Response, error) {
	var response *http.Response
	url := a.rootURL + endpoint
	request, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		return response, err
	}

	response, err = a.do(request, true)
	if err != nil {
		return response, err
	}

	return response, nil
}
