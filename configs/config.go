package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DriverNameDB       string `mapstructure:"DB_DRIVER"`
	ConnectionStringDB string `mapstructure:"CONNECTION_STRING_DB"`
	ServerPort         string `mapstructure:"SERVER_PORT"`
	ServerTimeout      int    `mapstructure:"SERVER_TIMEOUT_SECONDS"`
	EconomyBaseURL     string `mapstructure:"ECONOMY_BASE_URL"`
	DevelopmentMode    bool   `mapstructure:"DEVELOPMENT_MODE"`
	HttpClientTimeout  int    `mapstructure:"HTTP_CLIENT_TIMEOUT_MS"`
	DBTimeout          int    `mapstructure:"DB_TIMEOUT_MS"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
