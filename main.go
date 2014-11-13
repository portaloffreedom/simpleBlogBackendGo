package main

import (
	"fmt"
	"log"
	"os"

	"github.com/portaloffreedom/simpleBlogBackendGo/config"
	"github.com/portaloffreedom/simpleBlogBackendGo/database"
	"github.com/portaloffreedom/simpleBlogBackendGo/network"
)

func main() {
	log.Print("Hello, 世界")
	conf, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error reading config file: " + err.Error())
	}

	if conf.Mongo.MongoHost == "" {
		conf.Mongo.MongoHost = os.Getenv("OPENSHIFT_MONGODB_DB_HOST")
	}
	if conf.Mongo.MongoPort == "" {
		conf.Mongo.MongoPort = os.Getenv("OPENSHIFT_MONGODB_DB_PORT")
	}
	if conf.HTTP.BindAddress == "" {
		conf.HTTP.BindAddress = os.Getenv("HOST")
	}
	if conf.HTTP.BindPort == "" {
		conf.HTTP.BindPort = os.Getenv("PORT")
	}

	var dbAuth string
	if conf.Mongo.MongoUser != "" && conf.Mongo.MongoPasswd != "" {
		dbAuth = fmt.Sprintf("%s:%s@", conf.Mongo.MongoUser, conf.Mongo.MongoPasswd)
	}

	dbAddress := fmt.Sprintf("mongodb://%s%s:%s", dbAuth, conf.Mongo.MongoHost, conf.Mongo.MongoPort)
	log.Print("trying to connect to " + dbAddress)
	database.Connect(dbAddress)

	bind := fmt.Sprintf("%s:%s", conf.HTTP.BindAddress, conf.HTTP.BindPort)
	network.StartServer(bind)

	database.Disconnect()
}
