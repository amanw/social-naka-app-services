package event

import (
	"fmt"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	"github.com/amanw/social-naka-app-services/pkg/domain"
)

//Service -> Event Service
type Service struct {
	svcdb domain.EventRepo
}

//NewService -> Creates a new event service
func NewService(eventRepo domain.EventRepo) (*Service, error) {
	if eventRepo == nil {
		return &Service{}, fmt.Errorf("db connection is nil")
	}

	return &Service{
		svcdb: eventRepo,
	}, nil
}

//CreateEvent -> This service helps Event to create the Event in the db
func (svc *Service) CreateEvent(req *models.Event) (*string, error) {

	response, respErr := svc.svcdb.CreateNewEvent(req)
	if respErr != nil {
		return nil, fmt.Errorf("failed to insert new Event %w", respErr)
	}

	return response, nil

}

// GetEvents -> Get the information of all the Events
func (svc *Service) GetEvents() ([]*models.Event, error) {

	response, err := svc.svcdb.GetAllEvents()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch the Events %w", err)
	}

	return response, nil

}

//GetEventByID -> This service gets the Event information By Id
func (svc *Service) GetEventByID(id string) (*models.Event, error) {

	response, err := svc.svcdb.GetSingleEvent(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get the Event %w", err)
	}

	return response, nil

}

//UpdateEventByID -> This service updates the Event information By Id
func (svc *Service) UpdateEventByID(req *models.Event) (*string, error) {
	var err error

	response, err := svc.svcdb.EditEvent(req)

	if err != nil {
		return nil, fmt.Errorf("failed to update the Event %w", err)
	}

	return response, nil

}

//DeleteEventByID -> This service deletes the Event information By Id
func (svc *Service) DeleteEventByID(id string) (bool, error) {

	response, err := svc.svcdb.DeleteEvent(id)

	if err != nil && !response {
		return false, fmt.Errorf("failed to delete the event %w", err)
	}

	return response, nil

}
