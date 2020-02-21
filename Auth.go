package go_aspace

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type Configuration struct {
	url      string
	username string
	password string
}

var conf Configuration
var SessionKey string

func Init() error {
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

func getSessionKey() error {

	url := fmt.Sprintf("%s/users/%s/login?password=%s", conf.url, conf.username, conf.password)
	request, err := http.Post(url, "text/json", nil)
	if err != nil {
		return err
	}
	if request.StatusCode != 200 {
		return fmt.Errorf("Did not return a 200, recieved a %d", request.StatusCode)
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	var result map[string]string
	json.Unmarshal(body, &result)
	SessionKey = result["session"]
	return nil
}
