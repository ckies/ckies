package ckies

import (
	"errors"
	"io/ioutil"
	"path"

	"github.com/BurntSushi/toml"
)

// Service contains configuration with cookies
type Service struct {
	ContainsCookies

	Name    string
	Info    string
	Website string
}

// GetServiceFromPrefix returns a service configuration with cookies
func GetServiceFromPrefix(key string, prefix string) (*Service, error) {
	var data Service

	file := path.Join(prefix, key+".toml")
	content, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, errors.New("Unable to find service configuration: " + file)
	}

	if _, err := toml.Decode(string(content), &data); err != nil {
		return nil, errors.New("Cannot parse TOML configuration: " + file)
	}

	return &data, nil
}
