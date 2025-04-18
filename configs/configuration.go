package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.BindEnv("WEB_SERVER_PORT")
	viper.BindEnv("WEATHER_API_KEY")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.WebServerPort == "" {
		return nil, fmt.Errorf("WEB_SERVER_PORT is required")
	}

	if cfg.WeatherApiKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY is required")
	}

	return cfg, nil
}
