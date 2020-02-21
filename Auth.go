package go_aspace

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	url      string
	username string
	password string
}

func Init() error {
	viper.SetConfigFile("conf.json")
	viper.AddConfigPath("github.com/nyudlts/go-aspace")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err := viper.Unmarshal(&configuration)

	return err

}
