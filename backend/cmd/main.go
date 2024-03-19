package main

import (
	"backend/adapters/delivery/rest"
	db "backend/adapters/persistence"
	"backend/config"
)

func main() {
	// Start Config
	_ = config.GetConfig()
	// Start DB
	db.ConnectX()

	// Start Server
	rest.Start()

	defer db.Close()
}
