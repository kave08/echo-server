package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var Cfg Config
type Config struct {
	Server `yml:"server"`
}

type Server struct {
	Port string `yml:"port"`
}

func LoadConfig() error{
	file, err := os.Open("config.yml")
	if err != nil{
		return err
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&Cfg)
	if err != nil{
		return err
	}

	return nil
}