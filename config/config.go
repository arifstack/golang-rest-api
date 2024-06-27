package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	ConnectionString string `yaml:"connection_string"`
}

func LoadDevConfig() Config {
	return loadConfig("config/dev.yaml")
}

func LoadProdConfig() Config {
	return loadConfig("config/prod.yaml")
}

func loadConfig(filename string) Config {
	var cfg Config

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode config: %v", err)
	}

	return cfg
}
