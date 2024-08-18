package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PageTitle    string `envconfig:"HCMP_PAGE_TITLE"`
	PageLanguage string `envconfig:"HCMP_PAGE_LANGUAGE"`
}

func NewConfig() (Config, error) {
	var conf Config
	err := envconfig.Process("hcmp", &conf)

	return conf, err
}

func GetConfig() Config {
	return Config{
		PageTitle:    "GUACHIN",
		PageLanguage: "en",
	}

}
