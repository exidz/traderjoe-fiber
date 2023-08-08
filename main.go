package main

import (
	"github.com/exidz/traderjoe-fiber/cmd/server"
	"github.com/exidz/traderjoe-fiber/config"
)

// @title TraderJoe Fiber Price pool feed API
// @version 1.0
// @description Traderjoe price pool feed Rest API, developed in Golang using Fiber. Support v2.1 Liquidity Book Contracts, v2 Liquidity Book Contracts, and v1 Trader Joe Contracts across three major chains: Avalanche, Arbitrum, and Binance Smart Chain.
// @contact.name API Support
// @contact.email exidz@protonmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	config.LoadAllConfigs(".env")

	server.Server()
}
