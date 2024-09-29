package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type config struct {
	Server server `yaml:"server"`
	Db     db     `yaml:"db"`
}

type server struct {
	Port  string `yaml:"port"`
	Model string `yaml:"model"`
}
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	Maxopen  int    `yaml:"maxopen"`
}

var Config *config

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
