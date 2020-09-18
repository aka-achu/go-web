package main

import (
	"github.com/aka-achu/go-web/cmd"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/repo"
	"github.com/subosito/gotenv"
	"log"
)

func init() {
	if gotenv.Load(".env") != nil {
		log.Fatal("Failed to load the env file")
	}
	logging.Initialize()
	repo.Initialize()
}

func main() {
	cmd.Execute()
}
