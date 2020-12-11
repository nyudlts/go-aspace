package aspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"net/http"
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

func NewClient(environment string, timeout int) (*ASClient, error) {

	var client *ASClient

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    time.Duration(timeout) * time.Second,
		DisableCompression: true,
	}

	nclient := &http.Client{
		Transport: tr,
	}

	creds, err := getCreds(environment)
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

func getCreds(environment string) (Creds, error) {
	credsMap := map[string]Creds{}
	source, err := ioutil.ReadFile("/etc/go-aspace.yml")
	if err != nil {
		return Creds{}, err
	}

	err = yaml.Unmarshal(source, &credsMap)
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

	request, err := http.Post(url, "text/json", nil)
	if err != nil {
		return "", err
	}

	if request.StatusCode != 200 {
		return "", fmt.Errorf("Did not return a 200 while authenticating, recieved a %d", request.StatusCode)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	sessionKey := fmt.Sprintf("%v", result["session"])

	if sessionKey != "" {
		return sessionKey, nil
	} else {
		return sessionKey, fmt.Errorf("Session field was empty")
	}
}

func (a *ASClient) do(request *http.Request, authenticated bool) (*http.Response, error) {
	var response *http.Response

	if authenticated {
		request.Header.Add("X-ArchivesSpace-Session", a.sessionKey)
	}

	response, err := a.nclient.Do(request)
	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
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
