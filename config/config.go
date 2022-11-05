package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type rabbit struct {
	Url string `yaml:"url"`
}
type mgo struct {
	Url string `yaml:"url"`
}
type Config struct {
	Rabbit rabbit `yaml:"rabbit"`
	Mongo  mgo    `yaml:"mgo"`
}

func Load(filePath string) (*Config, error) {
	var (
		err error
		cfg Config
	)
	if filePath == "" {
		err = errors.New("filePath is empty")
		return nil, err
	}
	config, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(config, &cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println("1", cfg.Rabbit.Url)
	return &cfg, nil
}
