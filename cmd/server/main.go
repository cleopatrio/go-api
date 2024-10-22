package main

import (
	injections "github.com/dock-tech/notes-api/internal/config/injections/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./env/.env.localstack")
	if err != nil {
		err := godotenv.Load("../../env/.env.localstack")
		if err != nil {
			panic(err)
		}
	}
	server, err := injections.Wire().InitializeServer()
	if err != nil {
		panic(err)
	}

	err = server.Serve()
	if err != nil {
		panic(err)
	}
}
