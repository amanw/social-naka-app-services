package handlers

import (
	"context"
	"fmt"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	api "github.com/amanw/social-naka-app-services/pkg/api/restapi/operations/users"
	"github.com/amanw/social-naka-app-services/pkg/api/utils"
	"github.com/amanw/social-naka-app-services/pkg/domain"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// UserHandler : Object
type UserHandler struct {
	userService domain.UserSvc
}

//NewUserHandler : will be used to initialize the User Handler Object
func NewUserHandler(userService domain.UserSvc) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

//RegisterUser will enable the client to post the user to the DB
func (u UserHandler) RegisterUser(ctx context.Context, params api.RegisterUserParams) middleware.Responder {
	// fmt.Println(params.HTTPRequest.Header.Get("X-API-Key"))
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), "")
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewRegisterUserUnauthorized().WithPayload(&err)
	// }
	postStatus, err := u.userService.CreateUser(params.UserBody)

	fmt.Println("postStatus:", *postStatus)
	if err != nil {
		postErr := models.Error{
			Status:  400,
			Message: err.Error(),
		}
		return api.NewRegisterUserBadRequest().WithPayload(&postErr)
		logrus.Warnf(postErr.Message)
	}
	postErr := utils.ParseStatus(*postStatus)
	if postErr == nil {
		logrus.Infoln("The user has been created in the DB:", *postStatus)
		params.UserBody.Password = nil
		params.UserBody.ID = postStatus
		return api.NewRegisterUserOK().WithPayload(params.UserBody)
	}

	return api.NewRegisterUserConflict().WithPayload(postErr)
}

//GetUsers will enable the client to get all the users from the DB
func (u UserHandler) GetUsers(ctx context.Context, params api.GetUsersParams) middleware.Responder {
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), "")
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewGetUsersUnauthorized().WithPayload(&err)
	// }
	userList, err := u.userService.GetUsers()
	if err != nil {
		postErr := fmt.Errorf("failed to get the users: %w", err)
		logrus.Warnf(postErr.Error())
	}
	logrus.Infoln("The user has been created in the DB:", userList)

	return api.NewGetUsersOK().WithPayload(userList)
}

//GetUserbyID pulls the user information By its id
func (u UserHandler) GetUserbyID(ctx context.Context, params api.GetUserbyIDParams) middleware.Responder {
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), params.ID)
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewGetUserbyIDUnauthorized().WithPayload(&err)
	// }
	user, err := u.userService.GetUserByID(params.ID)
	if err != nil {
		postErr := fmt.Errorf("failed to get the users: %w", err)
		logrus.Warnf(postErr.Error())
	}
	logrus.Infoln("The user has been created in the DB:", user)

	return api.NewGetUserbyIDOK().WithPayload(user)
}

//UpdateUserbyID -> updates the user information By its id
func (u UserHandler) UpdateUserbyID(ctx context.Context, params api.UpdateUserbyIDParams) middleware.Responder {
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), params.ID)
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewUpdateUserbyIDUnauthorized().WithPayload(&err)
	// }
	editStatus, err := u.userService.UpdateUserByID(params.UserBody)

	fmt.Println("editStatus:", *editStatus)
	if err != nil {
		editErr := models.Error{
			Status:  400,
			Message: err.Error(),
		}
		return api.NewUpdateUserbyIDBadRequest().WithPayload(&editErr)
		logrus.Warnf(editErr.Message)
	}
	editErr := utils.ParseStatus(*editStatus)
	if editErr == nil {
		logrus.Infoln("The user has been updated in the DB:", *editStatus)
		params.UserBody.Password = nil
		params.UserBody.ID = editStatus
		return api.NewUpdateUserbyIDOK().WithPayload(params.UserBody)
	}

	return api.NewUpdateUserbyIDConflict().WithPayload(editErr)
}

//DeleteUserbyID -> deletes the user information By its id
func (u UserHandler) DeleteUserbyID(ctx context.Context, params api.DeleteUserbyIDParams) middleware.Responder {
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), params.ID)
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewRegisterUserUnauthorized().WithPayload(&err)
	// }
	status, err := u.userService.DeleteUserByID(params.ID)
	if err != nil {
		postErr := fmt.Errorf("failed to get the users: %w", err)
		logrus.Warnf(postErr.Error())
		Errstatus := models.Error{}
		Errstatus.Status = 500
		Errstatus.Message = postErr.Error()
		return api.NewDeleteUserbyIDInternalServerError().WithPayload(&Errstatus)
	}
	modelStatus := models.Error{}
	if status {
		logrus.Infoln("The user has been deleted in the DB:", status)
		modelStatus.Status = 200
		modelStatus.Message = "The user has been deleted from the DB"
	}

	return api.NewDeleteUserbyIDOK().WithPayload(&modelStatus)
}
