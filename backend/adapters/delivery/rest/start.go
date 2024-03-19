package rest

import (
	"backend/config"
	"log"
)

func Start() {
	cfg := config.GetConfig()
	port := cfg.HttpPort
	if port == "" {
		port = "8084"
	}
	log.Printf("Server starting in port %s", port)

	server := BuildServer()
	server.Addr = ":" + port

	// Start the server
	log.Fatal(server.ListenAndServe())
}
