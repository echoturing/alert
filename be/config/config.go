package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host  string       `yaml:"host"`
	Port  int          `yaml:"port"`
	Mysql *mysqlConfig `yaml:"mysql"`
}

type mysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}

func LoadConfig(configFile string) *Config {
	cfg := &Config{}
	fd, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
