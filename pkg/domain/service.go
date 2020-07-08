package domain

import "github.com/amanw/social-naka-app-services/pkg/api/models"

//UserSvc -> Creates a service for User Information
type UserSvc interface {
	CreateUser(*models.User) (*string, error)
	GetUsers() ([]*models.User, error)
	GetUserByID(string) (*models.User, error)
	UpdateUserByID(*models.User) (*string, error)
	DeleteUserByID(string) (bool, error)
	Validate(token, id string) error
}
