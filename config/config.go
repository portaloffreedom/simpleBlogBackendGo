// Package config managing config file loading
package config

import "os"

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

var Conf *Config

// LoadConfig read the config of the file
func LoadConfig() *Config {
	Conf = &Config{}

	Conf.Mongo.MongoUser = os.Getenv("OPENSHIFT_MONGODB_DB_USERNAME")
	Conf.Mongo.MongoPasswd = os.Getenv("OPENSHIFT_MONGODB_DB_PASSWORD")
	Conf.Mongo.DBName = "blog"
	Conf.Mongo.MongoHost = os.Getenv("OPENSHIFT_MONGODB_DB_HOST")
	Conf.Mongo.MongoPort = os.Getenv("OPENSHIFT_MONGODB_DB_PORT")
	Conf.HTTP.BindAddress = os.Getenv("HOST")
	Conf.HTTP.BindPort = os.Getenv("PORT")

	return Conf
}
