package main

import (
	"github.com/amanw/social-naka-app-services/pkg/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Applicattion instantiation
	app, err := app.NewApp()

	if err != nil {
		log.WithError(err).Fatal("The initialization of the application failed.")
	}

	// Starting the app
	if err = app.Start(); err != nil {
		log.WithError(err).Fatal("Failed during running the application.")
	}

}
