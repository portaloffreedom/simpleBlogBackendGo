package database

import (
	"fmt"
	"log"
	"time"

	"github.com/portaloffreedom/simpleBlogBackendGo/config"
	"gopkg.in/mgo.v2"
)

var (
	dbSession *mgo.Session
	db        *mgo.Database
)

// DatabaseError data structure for database error
type DatabaseError struct {
	When time.Time
	What string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func checkError(err error) bool {
	if err != nil {
		log.Print("#ERROR! " + err.Error())
	}
	return err != nil
}

// Connect to the database
func Connect(url string) {
	log.Print("Trying to connect to the database: " + url)

	session, err := mgo.Dial(url)
	dbSession = session

	if checkError(err) {
		log.Fatal("error connecting")
	} else {
		log.Print("Connection Successfull")
	}

	info, err := dbSession.BuildInfo()
	if !checkError(err) {
		log.Print("database info" +
			"\nversion: " + info.Version +
			"\nGitVersion: " + info.GitVersion +
			"\nOpenSSLVersion: " + info.OpenSSLVersion +
			"\nSysInfo: " + info.SysInfo)
	}

	db = dbSession.DB(config.Conf.Mongo.DBName)
	log.Print("connection to database established")
}

// Disconnect from the database
func Disconnect() {
	dbSession.Close()
	dbSession = nil
	db = nil
	log.Print("Connection closed")
}
