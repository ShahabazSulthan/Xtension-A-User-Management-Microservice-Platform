package config

import "github.com/spf13/viper"

type PortManager struct {
	PortNo         string `mapstructure:"PORTNO"`
	PostNrelSvcUrl string `mapstructure:"POSTNREL_SVC_URL"`
}

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBHost     string `mapstructure:"DBHOST"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBPort     string `mapstructure:"DBPORT"`
}

type RedisConfigs struct {
	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
	RedisDB   int    `mapstructure:"REDIS_DB"`
}

type Config struct {
	PortMngr PortManager
	DB       DataBase
	Redis    RedisConfigs
}

func LoadConfig() (*Config, error) {
	var PortManager PortManager
	var db DataBase
	var redis RedisConfigs

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&PortManager)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&redis)
	if err != nil {
		return nil, err
	}

	cofig := Config{PortMngr: PortManager, DB: db}
	return &cofig, nil
}
