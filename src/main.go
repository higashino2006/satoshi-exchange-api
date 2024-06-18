package main

import (
	"se-api/src/internal/config"
	"se-api/src/internal/db"
	"se-api/src/internal/lib/firebase_client"
	"se-api/src/internal/routes"
)

func main() {
	// set configuration
	if config.Init() != nil {
		return
	}
	// initialize firebase
	if firebase_client.Init() != nil {
		return
	}
	// initialize database
	if db.Init() != nil {
		return
	}
	// start server
	routes.Start()
}
