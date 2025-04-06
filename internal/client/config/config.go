package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

const (
	defaultConfigFilepath = "./config/client/config.yaml"
)

var (
	_once sync.Once
)

type Config struct {
	UsersServiceHost   string `env:"USERS_SERVICE_HOST" yaml:"users_service_host"`
	SecretsServiceHost string `env:"SECRETS_SERVICE_HOST" yaml:"secrets_service_host"`
}

func MustGetOnce() Config {
	var err error
	var config Config

	_once.Do(func() {
		config, err = parseConfig()
		if err != nil {
			panic(err)
		}
	})

	return config
}

func parseConfig() (Config, error) {
	var config Config

	f := parseFlags()

	configFilePath := defaultConfigFilepath
	if f.ConfigFilePath != "" {
		configFilePath = f.ConfigFilePath
	}

	err := config.parseEnv()
	if err != nil {
		return config, fmt.Errorf("parsing env: %w", err)
	}

	err = config.parseYaml(configFilePath)
	if err != nil {
		return config, fmt.Errorf("parsing config file: %w", err)
	}

	return config, nil
}

func (c *Config) parseEnv() error {
	return env.Parse(c)
}

func (c *Config) parseYaml(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open config json file: %w", err)
	}
	defer file.Close()

	return yaml.NewDecoder(file).Decode(c)
}

type flags struct {
	ConfigFilePath string
}

func parseFlags() flags {
	f := flags{}
	flag.StringVar(&f.ConfigFilePath, "c", "", "Path to config file")

	flag.Parse()

	return f
}
