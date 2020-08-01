package main

import (
	"os"

	"github.com/amanw/social-naka-app-services/pkg/app"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	// this is just to build arguments for dev
	args := os.Args
	logrus.Print(args)
	var arg string
	if len(args) > 1 {
		arg = args[1]
	}
	// Applicattion instantiation
	app, err := app.NewApp()

	if err != nil {
		log.WithError(err).Fatal("The initialization of the application failed.")
	}

	// Starting the app
	if err = app.Start(arg); err != nil {
		log.WithError(err).Fatal("Failed during running the application.")
	}

}
