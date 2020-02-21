package go_aspace

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	url      string
	username string
	password string
}

var conf Configuration

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
