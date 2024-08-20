package main

import (
	injections "github.com/dock-tech/notes-api/internal/config/injections/server"
	"github.com/dock-tech/notes-api/pkg/config"
)

func main() {
	config.LoadEnv("config")
	server, err := injections.InitializeServer()
	if err != nil {
		panic(err)
	}

	server.Serve()
}
