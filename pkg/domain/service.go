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

//LoginSvc -> Creates a Login Service
type LoginSvc interface {
	LoginUser(*models.LoginUser) (bool, error)
}

//EventSvc -> Creates a service for Event Information
type EventSvc interface {
	CreateEvent(*models.Event) (*string, error)
	GetEvents() ([]*models.Event, error)
	GetEventByID(string) (*models.Event, error)
	UpdateEventByID(*models.Event) (*string, error)
	DeleteEventByID(string) (bool, error)
}

//PostSvc -> Creates a service for Post Information
type PostSvc interface {
	CreatePost(*models.Post) (*string, error)
	GetPosts() ([]*models.Post, error)
	GetPostByID(string) (*models.Post, error)
	UpdatePostByID(*models.Post) (*string, error)
	DeletePostByID(string) (bool, error)
}
