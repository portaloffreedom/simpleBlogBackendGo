#!/bin/bash

export GOPATH=~/.go

#export OPENSHIFT_MONGODB_DB_USERNAME="admin"
#export OPENSHIFT_MONGODB_DB_PASSWORD="admin"
export OPENSHIFT_MONGODB_DB_HOST="localhost"
export OPENSHIFT_MONGODB_DB_PORT="27017"
export HOST="localhost"
export PORT="4000"

go run main.go


