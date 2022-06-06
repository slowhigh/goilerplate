#!/usr/bin/env bash
set -xe

# install package and dependencies
go get github.com/gin-gonic/gin

go get github.com/go-playground/validator/v10

# build command
go build -o bin/server server.go