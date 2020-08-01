package domain

import (
	"github.com/amanw/social-naka-app-services/pkg/api/models"
)

//UserRepo -> Interface for User DB
type UserRepo interface {
	CreateNewUser(*models.User) (*string, error)
	GetAllUsers() ([]*models.User, error)
	GetSingleUser(string) (*models.User, error)
	EditUser(*models.User) (*string, error)
	DeleteUser(string) (bool, error)
	LoginUser(*models.LoginUser) (bool, error)
}

//EventRepo -> Interface for Event DB
type EventRepo interface {
	CreateNewEvent(*models.Event) (*string, error)
	GetAllEvents() ([]*models.Event, error)
	GetSingleEvent(string) (*models.Event, error)
	EditEvent(*models.Event) (*string, error)
	DeleteEvent(string) (bool, error)
}
