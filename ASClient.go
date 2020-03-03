package go_aspace

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
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

type Configuration struct {
	url      string
	username string
	password string
}

var conf Configuration

func initConfig() error {
	viper.SetConfigFile("conf.json")
	viper.AddConfigPath("github.com/nyudlts/go-aspace")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return err
	}

	conf.username = viper.GetString("username")
	conf.password = viper.GetString("password")
	conf.url = viper.GetString("url")

	return nil
}

func GetRootURL() (string, error) {
	return conf.url, nil
}

func getSessionKey() (string, error) {

	sessionKey := ""

	err := initConfig()
	if err != nil {
		return sessionKey, err
	}

	url := fmt.Sprintf("%s/users/%s/login?password=%s", conf.url, conf.username, conf.password)

	request, err := http.Post(url, "text/json", nil)
	if err != nil {
		return sessionKey, err
	}

	if request.StatusCode != 200 {
		return sessionKey, fmt.Errorf("Did not return a 200 while authenticating, recieved a %d", request.StatusCode)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return sessionKey, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return sessionKey, err
	}
	sessionKey = fmt.Sprintf("%v", result["session"])

	if sessionKey != "" {
		return sessionKey, nil
	} else {
		return sessionKey, fmt.Errorf("Session field was empty")
	}
}
