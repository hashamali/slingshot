package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

// Config contains config for this project.
type Config struct {
	SMTPConfig      `mapstructure:",squash"`
	EmailsConfig    `mapstructure:",squash"`
	RecaptchaConfig `mapstructure:",squash"`
}

// GetConfigFromEnvVars get config from environment variables.
func GetConfigFromEnvVars() (*Config, error) {
	c := Config{}
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// GetConfigFromFile will get config from YAML file.
func GetConfigFromFile() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
