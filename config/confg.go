package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	MarvelPrivateKey string `mapstructure:"marvel_private_key"`
	MarvelPublicKey  string `mapstructure:"marvel_public_key"`
	MarvelAPIBaseURL string `mapstructure:"marvel_api_base_url"`
	Port             string `mapstructure:"port"`
}

var envVars = []string{
	"MARVEL_PUBLIC_KEY",
	"MARVEL_PRIVATE_KEY",
	"MARVEL_API_BASE_URL",
	"PORT",
}

func Config() (*Configuration, error) {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	for _, v := range envVars {
		os.Setenv(v, viper.GetString(v))
	}

	return &config, nil
}
