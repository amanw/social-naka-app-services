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
	registrationRes := &models.RegistrationResponse{}
	if params.UserBody.RoleName == "" {
		params.UserBody.RoleName = "USER"
	}
	postStatus, err := u.userService.CreateUser(params.UserBody)

	fmt.Println("postStatus:", *postStatus)
	if err != nil {
		postErr := models.Response{
			Code:    400,
			Status:  "Failed",
			Message: err.Error(),
		}
		logrus.Warnf(postErr.Message)
		return api.NewRegisterUserBadRequest().WithPayload(&postErr)
	}
	postErr := utils.ParseStatus(*postStatus)
	if postErr == nil {
		logrus.Infoln("The user has been created in the DB:", *postStatus)
		params.UserBody.Password = ""
		params.UserBody.ID = *postStatus
		// postErr.Code = 200
		// postErr.Status = "Success"
		// postErr.Message = "Registered Succesfully"
		logrus.Println(".........................................")
		registrationRes.Response = &models.Response{
			Code:    200,
			Status:  "Success",
			Message: "Registered Successfully",
		}
		registrationRes.User = params.UserBody
		logrus.Println(registrationRes.Response)
		return api.NewRegisterUserOK().WithPayload(registrationRes)
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
	userResponse := &models.UserResponse{}
	userList, err := u.userService.GetUsers()
	if err != nil {
		postErr := fmt.Errorf("failed to get the users: %w", err)
		logrus.Warnf(postErr.Error())
		responseErr := &models.Response{
			Status:  "Failed",
			Code:    400,
			Message: postErr.Error(),
		}
		return api.NewGetUsersBadRequest().WithPayload(responseErr)
	}

	logrus.Infoln("We get the User List:", userList)
	userResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "Users has been fetched Successfully",
	}
	userResponse.Users = userList

	return api.NewGetUsersOK().WithPayload(userResponse)
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
	userResponse := &models.RegistrationResponse{}
	user, err := u.userService.GetUserByID(params.ID)
	if err != nil {
		postErr := fmt.Errorf("failed to get the users: %w", err)
		logrus.Warnf(postErr.Error())
		getUserErr := &models.Response{
			Status:  "Failed",
			Code:    400,
			Message: postErr.Error(),
		}
		return api.NewGetUserbyIDBadRequest().WithPayload(getUserErr)

	}
	userResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "User has been fetched Successfully",
	}
	userResponse.User = user

	logrus.Infoln("The user has been fetched:", user)

	return api.NewGetUserbyIDOK().WithPayload(userResponse)
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
	editResponse := &models.RegistrationResponse{}
	params.UserBody.ID = params.ID
	editStatus, err := u.userService.UpdateUserByID(params.UserBody)

	fmt.Println("editStatus:", editStatus)
	if err != nil {
		editErr := models.Response{
			Code:    400,
			Status:  "Failed",
			Message: err.Error(),
		}
		logrus.Warnf(editErr.Message)
		return api.NewUpdateUserbyIDBadRequest().WithPayload(&editErr)
	}
	editErr := utils.ParseStatus(*editStatus)
	if editErr == nil {
		logrus.Infoln("The user has been updated in the DB:", *editStatus)
		params.UserBody.Password = ""
		params.UserBody.ID = *editStatus
		editResponse.Response = &models.Response{
			Code:    200,
			Status:  "Success",
			Message: "Edited Successfully",
		}
		editResponse.User = params.UserBody
		return api.NewUpdateUserbyIDOK().WithPayload(editResponse)
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
		postErr := fmt.Errorf("Deletion of the user failed: %w", err)
		logrus.Warnf(postErr.Error())
		Errstatus := models.Response{}
		Errstatus.Code = 500
		Errstatus.Status = "Failed"
		Errstatus.Message = postErr.Error()
		return api.NewDeleteUserbyIDInternalServerError().WithPayload(&Errstatus)
	}
	modelStatus := models.Response{}
	if status {
		logrus.Infoln("The user has been deleted in the DB:", status)
		modelStatus.Code = 200
		modelStatus.Status = "Success"
		modelStatus.Message = "The user has been deleted from the DB"
	}

	return api.NewDeleteUserbyIDOK().WithPayload(&modelStatus)
}
