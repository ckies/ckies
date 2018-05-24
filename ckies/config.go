package ckies

import (
	"errors"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Config stores data from the config file
type Config struct {
	ContainsCookies

	Name    string
	Info    string
	Website string

	Options struct {
		ForceTables bool
	}

	Links struct {
		Policy   string
		Settings string
	}

	Services []string
}

// Load tries to load a config file
func Load(file string) (*Config, error) {
	var data Config

	content, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, errors.New("Unable to find configuration: " + file)
	}

	if _, err := toml.Decode(string(content), &data); err != nil {
		return nil, errors.New("Cannot parse TOML configuration: " + file)
	}

	return &data, nil
}
