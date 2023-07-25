package main

import (
	"github.com/exidz/traderjoe-fiber/cmd/server"
	"github.com/exidz/traderjoe-fiber/config"
)

func main() {
	config.LoadAllConfig()

	server.Server()
}
