package main

import (
	"fmt"
	"log"

	"github.com/portaloffreedom/simpleBlogBackendGo/config"
	"github.com/portaloffreedom/simpleBlogBackendGo/database"
	"github.com/portaloffreedom/simpleBlogBackendGo/network"
)

func main() {
	log.Print("Hello, 世界")
	conf := config.LoadConfig()

	var dbAuth string
	if conf.Mongo.MongoUser != "" && conf.Mongo.MongoPasswd != "" {
		dbAuth = fmt.Sprintf("%s:%s@", conf.Mongo.MongoUser, conf.Mongo.MongoPasswd)
	}

	dbAddress := fmt.Sprintf("mongodb://%s%s:%s", dbAuth, conf.Mongo.MongoHost, conf.Mongo.MongoPort)
	database.Connect(dbAddress)

	bind := fmt.Sprintf("%s:%s", conf.HTTP.BindAddress, conf.HTTP.BindPort)
	network.StartServer(bind)

	database.Disconnect()
}
