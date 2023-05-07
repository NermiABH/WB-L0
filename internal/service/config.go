package service

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

const configPath = "configs/config.yml"

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Username string `yaml:"username"`
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		DBName   string `yaml:"dbname"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"db"`
	Stan struct {
		ClusterID string `yaml:"clusterID"`
		ClientID  string `yaml:"clientID"`
		Chanel    string `yaml:"chanel"`
	} `yaml:"nats-streaming"`
}

func NewConfig() *Config {
	config := &Config{}
	log.Println("read application configuration")
	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Fatalln(err)
	}
	return config
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("user=%s host=%s dbname=%s password=%s sslmode=%s ",
		c.DB.Username, c.DB.Host, c.DB.DBName, c.DB.Password, c.DB.SSLMode)
}

func (c *Config) ServerAddr() string {
	return c.Server.Host + c.Server.Port
}
