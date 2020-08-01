package handlers

import (
	"context"
	"fmt"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	api "github.com/amanw/social-naka-app-services/pkg/api/restapi/operations/events"
	"github.com/amanw/social-naka-app-services/pkg/domain"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// EventHandler : Object
type EventHandler struct {
	eventService domain.EventSvc
}

//NewEventHandler : will be used to initialize the Event Handler Object
func NewEventHandler(eventService domain.EventSvc) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

//PostEvent will enable the client to post the user to the DB
func (e EventHandler) PostEvent(ctx context.Context, params api.PostEventParams) middleware.Responder {
	// fmt.Println(params.HTTPRequest.Header.Get("X-API-Key"))
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), "")
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewRegisterUserUnauthorized().WithPayload(&err)
	// }
	eventRes := &models.EventResponse{}
	postStatus, err := e.eventService.CreateEvent(params.EventBody)

	fmt.Println("postStatus:", *postStatus)
	if err != nil {
		postErr := models.Response{
			Code:    400,
			Status:  "Failed",
			Message: err.Error(),
		}
		logrus.Warnf(postErr.Message)
		return api.NewPostEventBadRequest().WithPayload(&postErr)
	}
	logrus.Infoln("The event has been created in the DB:", *postStatus)

	eventRes.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "Event Posted Successfully",
	}
	params.EventBody.ID = *postStatus
	eventRes.Event = params.EventBody
	return api.NewPostEventOK().WithPayload(eventRes)
}

//GetEvents will enable the client to get all the users from the DB
func (e EventHandler) GetEvents(ctx context.Context, params api.GetEventsParams) middleware.Responder {

	eventResponse := &models.EventResponse{}
	eventList, err := e.eventService.GetEvents()
	if err != nil {
		postErr := fmt.Errorf("failed to get the events: %w", err)
		logrus.Warnf(postErr.Error())
		responseErr := &models.Response{
			Status:  "Failed",
			Code:    400,
			Message: postErr.Error(),
		}
		return api.NewGetEventsBadRequest().WithPayload(responseErr)
	}

	logrus.Infoln("We get the Event List:", eventList)
	eventResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "Events has been fetched Successfully",
	}
	eventResponse.Events = eventList

	return api.NewGetEventsOK().WithPayload(eventResponse)
}

//GetEventbyID pulls the user information By its id
func (e EventHandler) GetEventbyID(ctx context.Context, params api.GetEventbyIDParams) middleware.Responder {

	eventResponse := &models.EventResponse{}
	event, err := e.eventService.GetEventByID(params.ID)
	if err != nil {
		postErr := fmt.Errorf("failed to get the event: %w", err)
		logrus.Warnf(postErr.Error())
		getEventErr := &models.Response{
			Status:  "Failed",
			Code:    400,
			Message: postErr.Error(),
		}
		return api.NewDeleteEventbyIDBadRequest().WithPayload(getEventErr)

	}
	eventResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "Event has been fetched Successfully",
	}
	eventResponse.Event = event

	logrus.Infoln("The event has been fetched:", event)

	return api.NewGetEventbyIDOK().WithPayload(eventResponse)
}

//UpdateEventbyID -> updates the event information By its id
func (e EventHandler) UpdateEventbyID(ctx context.Context, params api.UpdateEventbyIDParams) middleware.Responder {
	editResponse := &models.EventResponse{}
	params.EventBody.ID = params.ID
	editStatus, err := e.eventService.UpdateEventByID(params.EventBody)

	fmt.Println("editStatus:", editStatus)
	if err != nil {
		editErr := models.Response{
			Code:    400,
			Status:  "Failed",
			Message: err.Error(),
		}
		logrus.Warnf(editErr.Message)
		return api.NewUpdateEventbyIDBadRequest().WithPayload(&editErr)
	}

	logrus.Infoln("The event has been updated in the DB:", *editStatus)
	params.EventBody.ID = *editStatus
	editResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "The event has been edited successfully",
	}
	editResponse.Event = params.EventBody
	return api.NewUpdateEventbyIDOK().WithPayload(editResponse)
}

//DeleteEventbyID -> deletes the event information By its id
func (e EventHandler) DeleteEventbyID(ctx context.Context, params api.DeleteEventbyIDParams) middleware.Responder {
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), params.ID)
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewRegisterUserUnauthorized().WithPayload(&err)
	// }
	status, err := e.eventService.DeleteEventByID(params.ID)
	if err != nil {
		postErr := fmt.Errorf("Deletion of the event failed: %w", err)
		logrus.Warnf(postErr.Error())
		Errstatus := models.Response{}
		Errstatus.Code = 500
		Errstatus.Status = "Failed"
		Errstatus.Message = postErr.Error()
		return api.NewDeleteEventbyIDInternalServerError().WithPayload(&Errstatus)
	}
	modelStatus := models.Response{}
	if status {
		logrus.Infoln("The event has been deleted in the DB:", status)
		modelStatus.Code = 200
		modelStatus.Status = "Success"
		modelStatus.Message = "The event has been deleted from the DB"
	}

	return api.NewDeleteEventbyIDOK().WithPayload(&modelStatus)
}
