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
}
