package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	UserSvcUrl string `mapstructure:"USER_SVC_URL"`
}

func LoadConfig() (*Config, error) {
	var c Config

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in Config : ", err)
		return nil, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("Error  marshaling : ", err)

		return nil, err
	}

	return &c, nil
}
