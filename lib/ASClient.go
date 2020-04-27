package lib

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
)

var Client ASClient

type ASClient struct {
	sessionKey string
	rootURL    string
	nclient    *http.Client
}

func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}

	if err := newClient(100); err != nil { // set the timeout in the configuration
		panic(err)
	}
}

func newClient(timeout int) error {

	var client *ASClient

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    time.Duration(timeout) * time.Second,
		DisableCompression: true,
	}

	nclient := &http.Client{
		Transport: tr,
	}

	token, err := getSessionKey()
	if err != nil {
		return err
	}

	aspaceRootURL, err := GetRootURL()
	if err != nil {
		return err
	}

	client = &ASClient{
		sessionKey: token,
		rootURL:    aspaceRootURL,
		nclient:    nclient,
	}

	Client = *client

	return nil
}

type Configuration struct {
	url      string
	username string
	password string
}

var conf Configuration

func initConfig() error {

	viper.SetConfigName("go-aspace")
	viper.AddConfigPath("/etc/sysconfig")
	viper.SetConfigType("json")

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
