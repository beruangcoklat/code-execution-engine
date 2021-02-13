package configlib

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var cfg Config

func Init(configPath string) error {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return err
	}

	return nil
}

func Get() Config {
	return cfg
}
