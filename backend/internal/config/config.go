package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug *bool         `yaml:"is_debug" env-required:"true"`
	Listen  Listener      `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type Listener struct {
	Type   string `yaml:"type"`
	BindIp string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("../config.yml", instance); err != nil {
			_, _ = cleanenv.GetDescription(instance, nil)
			log.Fatal("Error reading config:", err)
		}
	})
	return instance
}
