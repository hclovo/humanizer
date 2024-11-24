package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Auth struct {
		JwtSecret string `yaml:"jwt_secret"`
	} `yaml:"auth"`
	Database struct {
		Driver string `yaml:"driver"`
		DSN    string `yaml:"dsn"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		PoolSize int    `yaml:"pool_size"`
		PoolTimeout int `yaml:"pool_timeout"`
	} `yaml:"redis"`
	SMTP struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Sender string `yaml:"sender"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`
	Supabase struct {
		Url string `yaml:"url"`
		Key string `yaml:"key"`
		Project string `yaml:"project"`
	} `yaml:"supabase"`
}

var AppConfig Config

func LoadConfig() {
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
}
