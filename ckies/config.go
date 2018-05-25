package ckies

import (
	"errors"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// ConfigBase stores basic information
type ConfigBase struct {
	Name    string
	Info    string
	Website string
}

// ConfigLinks stores information about customizable links in content
type ConfigLinks struct {
	Policy   string
	Settings string
}

// ConfigOptions enables/disables behaviour and custom blocks
type ConfigOptions struct {
	ForceTables              bool
	IncludeFacebookAudiences bool
}

// Config stores data from the config file
type Config struct {
	ContainsCookies
	ConfigBase

	Options ConfigOptions
	Links   ConfigLinks

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
