package configs

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Host    string        `envconfig:"HCMP_APP_HOST"`
	Port    uint32        `envconfig:"HCMP_APP_PORT"`
	Timeout time.Duration `envconfig:"HCMP_APP_TIMEOUT" default:"10s"`
}

type PageConfig struct {
	PageTitle    string `envconfig:"HCMP_PAGE_TITLE"`
	PageLanguage string `envconfig:"HCMP_PAGE_LANGUAGE"`
}

type DBConfig struct {
	Host string `envconfig:"HCMP_DB_HOST"`
	Port uint32 `envconfig:"HCMP_DB_PORT"`
	User string `envconfig:"HCMP_DB_USER"`
	Pass string `envconfig:"HCMP_DB_PASS"`
	Name string `envconfig:"HCMP_DB_NAME"`
}

type Config struct {
	App      AppConfig
	Page     PageConfig
	Database DBConfig
}

func NewConfig() (Config, error) {
	var conf Config
	err := envconfig.Process("hcmp", &conf)

	return conf, err
}
