package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type App struct {
	Port string `yaml:"app_port"`
}

type Auth struct {
	HashSalt   string `yaml:"hash_salt"`
	SingingKey string `yaml:"singing_key"`
	Tll        uint8  `yaml:"tll"`
}

type Database struct {
	Host     string `yaml:"db_host"`
	Port     string `yaml:"db_port"`
	User     string `yaml:"db_user"`
	Password string `yaml:"db_password"`
	Name     string `yaml:"db_name"`
	SslMode  string `yaml:"db_ssl_mode"`
}

type Tls struct {
	Mode     bool   `yaml:"ssl_mode"`
	CertFile string `yaml:"ssl_cert_file"`
	KeyFile  string `yaml:"ssl_key_file"`
}

type Queues struct {
	CreateUser string `yaml:"create_user"`
}

type Rabbit struct {
	AmqpServerUrl   string `yaml:"amqp_server_url"`
	DelayWorkersRun uint8  `yaml:"delay_workers_run"`
	Queues
}

type Config struct {
	App
	Auth
	Database
	Tls
	Rabbit
}

func (c *Config) Init() (*Config, error) {
	file, err := os.Open("config/config.yml")
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err = yaml.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
