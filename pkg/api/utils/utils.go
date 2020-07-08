package utils

import (
	"github.com/amanw/social-naka-app-services/pkg/api/models"
)

//ParseStatus -> Parsing the status and converting it into a message
func ParseStatus(status string) *models.Error {
	err := models.Error{}
	if status == "S,UE,EE" {
		err.Status = 409
		err.Message = "Username and email already exist. Please try again."
	} else if status == "S,UE,ME" {
		err.Status = 409
		err.Message = "Username and mobile no already exist. Please try again."
	} else if status == "S,EE,ME" {
		err.Status = 409
		err.Message = "Email address and mobile no already exist. Please try again."
	} else if status == "S,UE" {
		err.Status = 409
		err.Message = "The username already exist. Please try with different username."
	} else if status == "S,EE" {
		err.Status = 409
		err.Message = "The email address already exist. Please try with different email address."
	} else if status == "S,ME" {
		err.Status = 409
		err.Message = "The mobile no already exist. Please try with different mobile no."
	} else if status == "S,UE,EE,ME" {
		err.Status = 409
		err.Message = "Username, Email address and mobile no already exist. Please try again."
	} else {
		return nil
	}

	return &err
}
