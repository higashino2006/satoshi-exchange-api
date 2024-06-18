package firebase_client

import (
	"context"
	"log"
	"se-api/src/internal/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var authClient *auth.Client

func Init() error {
	if config.AppConfig.TEST_MODE {
		return nil
	}

	// Initialize default app
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Printf("Error initializing firebase: %v", err)
		return err
	}

	// Access auth service from the default app
	authClient, err = app.Auth(context.Background())
	if err != nil {
		log.Printf("Error initializing firebase auth: %v", err)
		return err
	}

	return nil
}

func Auth() *auth.Client {
	return authClient
}
