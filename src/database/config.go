package database

import "github.com/spf13/viper"

type Config struct {
	Host                           string `mapstructure:"HOST"`
	DB_USER                        string `mapstructure:"DB_USER"`
	DB_PASS                        string `mapstructure:"DB_PASS"`
	DB_PORT                        string `mapstructure:"DB_PORT"`
	DB_NAME                        string `mapstructure:"DB_NAME"`
	CONN_NAME                      string `mapstructure:"CONN_NAME"`
	GOOGLE_APPLICATION_CREDENTIALS string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
	INSTANCE_CONNECTION_NAME       string `mapstructure:"INSTANCE_CONNECTION_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("gofinance")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
