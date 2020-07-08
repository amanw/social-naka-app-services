package user

import (
	"fmt"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	"github.com/amanw/social-naka-app-services/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

//Service -> User Service
type Service struct {
	svcdb domain.UserRepo
}

//NewService -> Creates a new user service
func NewService(userRepo domain.UserRepo) (*Service, error) {
	if userRepo == nil {
		return &Service{}, fmt.Errorf("db connection is nil")
	}

	return &Service{
		svcdb: userRepo,
	}, nil
}

//CreateUser -> This service helps user to create the user in the db
func (svc *Service) CreateUser(req *models.User) (*string, error) {
	var err error
	var pass string

	if req.Password != nil {
		pass = *req.Password
	}

	*req.Password, err = EncryptPass([]byte(pass))

	if err != nil {
		return nil, err
	}

	response, respErr := svc.svcdb.CreateNewUser(req)
	if respErr != nil {
		return nil, fmt.Errorf("failed to insert new user %w", respErr)
	}

	return response, nil

}

// GetUsers -> Get the information of all the users
func (svc *Service) GetUsers() ([]*models.User, error) {

	response, err := svc.svcdb.GetAllUsers()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch the users %w", err)
	}

	return response, nil

}

//GetUserByID -> This service gets the user information By Id
func (svc *Service) GetUserByID(id string) (*models.User, error) {

	response, err := svc.svcdb.GetSingleUser(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get the user %w", err)
	}

	return response, nil

}

//UpdateUserByID -> This service updates the user information By Id
func (svc *Service) UpdateUserByID(req *models.User) (*string, error) {
	var err error
	var pass string

	if req.Password != nil {
		pass = *req.Password
	}

	*req.Password, err = EncryptPass([]byte(pass))

	response, err := svc.svcdb.EditUser(req)

	if err != nil {
		return nil, fmt.Errorf("failed to update the user %w", err)
	}

	return response, nil

}

//DeleteUserByID -> This service deletes the user information By Id
func (svc *Service) DeleteUserByID(id string) (bool, error) {

	response, err := svc.svcdb.DeleteUser(id)

	if err != nil && !response {
		return false, fmt.Errorf("failed to delete the user %w", err)
	}

	return response, nil

}

//EncryptPass -> hashes the password
func EncryptPass(pass []byte) (string, error) {
	var err error

	encryptPass, passErr := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if passErr != nil {
		err = fmt.Errorf("Failed Encrypting the password %w", passErr)
	}

	return string(encryptPass), err
}
