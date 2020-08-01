package handlers

import (
	"context"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	api "github.com/amanw/social-naka-app-services/pkg/api/restapi/operations/login"
	"github.com/amanw/social-naka-app-services/pkg/domain"
	"github.com/go-openapi/runtime/middleware"
)

// LoginHandler : Object
type LoginHandler struct {
	loginService domain.LoginSvc
}

//NewLoginHandler : will be used to initialize the Login Handler Object
func NewLoginHandler(loginService domain.LoginSvc) *LoginHandler {
	return &LoginHandler{
		loginService: loginService,
	}
}

//LoginUser will check if the username and password exists in the db
func (l *LoginHandler) LoginUser(ctx context.Context, params api.LoginUserParams) middleware.Responder {

	response, loginErr := l.loginService.LoginUser(params.LoginBody)
	status := &models.Response{}
	if loginErr != nil {
		status.Code = 400
		status.Status = "Failed"
		status.Message = loginErr.Error()
		return api.NewLoginUserBadRequest().WithPayload(status)
	}
	if response {
		status.Code = 200
		status.Status = "Success"
		status.Message = "The user has been logged in Successfully"
		return api.NewLoginUserOK().WithPayload(&status)
	}
	status.Code = 404
	status.Status = "Not Found"
	status.Message = "The user does not exist in the DB"
	return api.NewLoginUserNotFound().WithPayload(status)

}
