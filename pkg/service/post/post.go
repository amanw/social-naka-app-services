package post

import (
	"fmt"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	"github.com/amanw/social-naka-app-services/pkg/domain"
)

//Service -> Post Service
type Service struct {
	svcdb domain.PostRepo
}

//NewService -> Creates a new post service
func NewService(postRepo domain.PostRepo) (*Service, error) {
	if postRepo == nil {
		return &Service{}, fmt.Errorf("db connection is nil")
	}

	return &Service{
		svcdb: postRepo,
	}, nil
}

//CreatePost -> This service helps Post to create the Post in the db
func (svc *Service) CreatePost(req *models.Post) (*string, error) {

	response, respErr := svc.svcdb.CreateNewPost(req)
	if respErr != nil {
		return nil, fmt.Errorf("failed to insert new Post %w", respErr)
	}

	return response, nil

}

// GetPosts -> Get the information of all the Posts
func (svc *Service) GetPosts() ([]*models.Post, error) {

	response, err := svc.svcdb.GetAllPosts()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch the Posts %w", err)
	}

	return response, nil

}

//GetPostByID -> This service gets the Post information By Id
func (svc *Service) GetPostByID(id string) (*models.Post, error) {

	response, err := svc.svcdb.GetSinglePost(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get the Post %w", err)
	}

	return response, nil

}

//UpdatePostByID -> This service updates the Post information By Id
func (svc *Service) UpdatePostByID(req *models.Post) (*string, error) {
	var err error

	response, err := svc.svcdb.EditPost(req)

	if err != nil {
		return nil, fmt.Errorf("failed to update the Post %w", err)
	}

	return response, nil

}

//DeletePostByID -> This service deletes the Post information By Id
func (svc *Service) DeletePostByID(id string) (bool, error) {

	response, err := svc.svcdb.DeletePost(id)

	if err != nil && !response {
		return false, fmt.Errorf("failed to delete the Post %w", err)
	}

	return response, nil

}
