package database

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	dbSession *mgo.Session
	db        *mgo.Database
)

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
	log.Print("Trying connection")

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

	db = dbSession.DB("blog")
}

// Disconnect from the database
func Disconnect() {
	dbSession.Close()
	log.Print("Connection closed")
}
