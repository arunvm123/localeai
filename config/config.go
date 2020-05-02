package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var configuration config

type config struct {
	Port     string         `yaml:"port"`
	Database databaseConfig `yaml:"database"`
}

type databaseConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database_name"`
	Host         string `yaml:"host"`
	Port         string `json:"port"`
}

// Initialise reads from the yaml file
// into the config struct
func Initialise(filepath string) (*config, error) {
	err := cleanenv.ReadConfig(filepath, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func GetConfig() (*config, error) {
	return &configuration, nil
}