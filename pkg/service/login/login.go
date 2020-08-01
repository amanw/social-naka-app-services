package login

import (
	"fmt"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	"github.com/amanw/social-naka-app-services/pkg/domain"
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

//LoginUser -> Checks if the user is valid
func (svc *Service) LoginUser(req *models.LoginUser) (bool, error) {

	response, err := svc.svcdb.LoginUser(req)

	if err != nil && !response {
		return false, fmt.Errorf("Failed to login the user %w", err)
	}

	return response, nil

}
