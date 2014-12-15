#!/bin/bash

export GOPATH=~/.go

#export OPENSHIFT_MONGODB_DB_USERNAME="admin"
#export OPENSHIFT_MONGODB_DB_PASSWORD="admin"
export MONGO_PORT_27017_TCP_ADDR="172.17.0.2"
export MONGO_PORT_27017_TCP_PORT="27017"
export HOST="localhost"
export PORT="4000"

go run main.go
