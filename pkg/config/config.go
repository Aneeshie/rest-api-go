package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env:"HTTP_SERVER_ADDRESS"`
}

type Config struct {
	ENV         string `yaml:"env" env:"ENV"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to configuration file.")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is not set.")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config files does not exist %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Cannot read config file %s", err.Error())
	}

	return &cfg
}
