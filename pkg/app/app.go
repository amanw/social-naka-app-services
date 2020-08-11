package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/amanw/social-naka-app-services/pkg/api/handlers"
	"github.com/amanw/social-naka-app-services/pkg/api/restapi"
	"github.com/amanw/social-naka-app-services/pkg/service/event"
	"github.com/amanw/social-naka-app-services/pkg/service/login"
	"github.com/amanw/social-naka-app-services/pkg/service/post"
	"github.com/amanw/social-naka-app-services/pkg/service/user"
	interpose "github.com/carbocation/interpose/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	repo "github.com/amanw/social-naka-app-services/pkg/repo/mongodb"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// App is use to define the social-naka-app-services
type App struct {
}

// NewApp will create a new application
func NewApp() (*App, error) {
	return &App{}, nil
}

// Init will Initialize the application with devType
func (a *App) Init(devType string) (http.Handler, error) {
	if devType == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	logrus.Print("Creating a DB Connection")
	// Create a mongo DB connection
	db := a.InitDB()
	logrus.Printf("The Db is connected: %v", db)

	userService, userErr := user.NewService(db)
	if userErr != nil {
		return nil, fmt.Errorf("User Service cannot be created %+v", userErr)
	}

	eventService, eventErr := event.NewService(db)
	if eventErr != nil {
		return nil, fmt.Errorf("Event Service cannot be created %+v", eventErr)
	}

	loginService, loginErr := login.NewService(db)
	if eventErr != nil {
		return nil, fmt.Errorf("Login Service cannot be created %+v", loginErr)
	}

	postService, postErr := post.NewService(db)
	if postErr != nil {
		return nil, fmt.Errorf("Post Service cannot be created %+v", postErr)
	}

	socialNakaAPIConfig := restapi.Config{
		Logger:          log.Debugf,
		UsersAPI:        handlers.NewUserHandler(userService),
		EventsAPI:       handlers.NewEventHandler(eventService),
		LoginAPI:        handlers.NewLoginHandler(loginService),
		PostsAPI:        handlers.NewPostHandler(postService),
		InnerMiddleware: interpose.NegroniLogrus(),
	}
	handler, err := restapi.Handler(socialNakaAPIConfig)
	if err != nil {
		log.Fatalf("Failed to init the rest API")
	}
	return setupGlobalMiddleware(handler), nil
}

// Run will begin serving the passed handler
func (a *App) Run(socialNakaHandler http.Handler) error {
	log.Info("Starting Social-Naka-App Server on port 8080")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), socialNakaHandler); err != nil {
		log.Fatalf("Service Failed to Start with Run Error: %s\n", err)
	}
	return nil
}

// Start will initialize and run the application
func (a *App) Start(arg string) error {
	socialNakaHandler, err := a.Init(arg)
	if err != nil {
		log.Fatalf("Failed to create the handler object from the configuration: %s", err)
	}

	return a.Run(socialNakaHandler)
}

// InitDB is use to connect the DB
func (a *App) InitDB() *repo.MongoDB {
	database, err := repo.DbConnect(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	return &repo.MongoDB{
		DB: database.DB,
	}
}

// The middleware configuration happens before anything,
// this middleware also applies to serving the swagger.json document.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "DELETE", "POST", "PUT"},
		MaxAge:         1000,
	})
	return corsHandler.Handler(handler)
}
