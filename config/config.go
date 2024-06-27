package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Configurator mendefinisikan metode untuk mengambil konfigurasi
type Configurator interface {
	LoadConfig() Config
}

// devConfigurator adalah implementasi untuk konfigurasi development
type devConfigurator struct{}

// prodConfigurator adalah implementasi untuk konfigurasi production
type prodConfigurator struct{}

// Config adalah struktur untuk menyimpan konfigurasi aplikasi
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig adalah konfigurasi untuk server
type ServerConfig struct {
	Port int `yaml:"port"`
}

// DatabaseConfig adalah konfigurasi untuk database
type DatabaseConfig struct {
	ConnectionString string `yaml:"connection_string"`
}

// LoadDevConfig mengambil konfigurasi untuk lingkungan development
func LoadDevConfig() Config {
	devConfig := devConfigurator{}
	return devConfig.LoadConfig()
}

// LoadProdConfig mengambil konfigurasi untuk lingkungan production
func LoadProdConfig() Config {
	prodConfig := prodConfigurator{}
	return prodConfig.LoadConfig()
}

func (dc devConfigurator) LoadConfig() Config {
	configPath := filepath.Join(getProjectRoot(), "config", "dev.yaml")
	return loadConfig(configPath)
}

func (pc prodConfigurator) LoadConfig() Config {
	configPath := filepath.Join(getProjectRoot(), "config", "prod.yaml")
	return loadConfig(configPath)
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

func getProjectRoot() string {
	// Ganti dengan fungsi yang mengembalikan path absolut root dari proyek
	return "/Users/zaydanmubaraq/Documents/Project/Golang/golang-rest-api/"
}
