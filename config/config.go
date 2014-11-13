// Package config managing config file loading
package config

import (
	"encoding/json"
	"io/ioutil"
)

type (
	// Mongo config for MongoDB connection
	Mongo struct {
		MongoUser,
		MongoPasswd,
		DBName,
		MongoHost,
		MongoPort string
	}
	// HTTPServer config for Http connection
	HTTPServer struct {
		BindAddress,
		BindPort string
	}
	// Config for MongoDB connection
	Config struct {
		Mongo Mongo
		HTTP  HTTPServer
	}
)

const configPath string = "config/config.json"

// ReadConfig read the config of the file
func ReadConfig() (*Config, error) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var conf Config

	err = json.Unmarshal(content, &conf)

	return &conf, nil
}
