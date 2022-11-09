package main

import (
	"github.com/joho/godotenv"
	"github.com/oxyrinchus/goilerplate/bootstrap"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
