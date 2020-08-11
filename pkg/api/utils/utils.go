package utils

import (
	"fmt"
	"log"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	repo "github.com/amanw/social-naka-app-services/pkg/repo/models"
	"golang.org/x/crypto/bcrypt"
)

//ParseStatus -> Parsing the status and converting it into a message
func ParseStatus(status string) *models.Response {
	err := models.Response{}
	err.Code = 409
	err.Status = "Failed"
	if status == "S,UE,EE" {
		err.Message = "Username and email already exist. Please try again."
	} else if status == "S,UE,ME" {
		err.Message = "Username and mobile no already exist. Please try again."
	} else if status == "S,EE,ME" {
		err.Message = "Email address and mobile no already exist. Please try again."
	} else if status == "S,UE" {
		err.Message = "The username already exist. Please try with different username."
	} else if status == "S,EE" {
		err.Message = "The email address already exist. Please try with different email address."
	} else if status == "S,ME" {
		err.Message = "The mobile no already exist. Please try with different mobile no."
	} else if status == "S,UE,EE,ME" {
		err.Message = "Username, Email address and mobile no already exist. Please try again."
	} else {
		return nil
	}

	return &err
}

//ParseUser -> Repo User model to api User Model
func ParseUser(user repo.User) *models.User {
	newUser := &models.User{}
	newUser.ID = user.ID.Hex()
	newUser.Username = user.Username
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.EmailAddress = user.EmailAddress
	newUser.MobileNo = user.MobileNo
	newUser.Password = user.Password
	newUser.IsStayingHere = user.IsStayingHere
	newUser.Description = user.Description
	newUser.Dob = &user.Dob
	newUser.CreatedAt = user.CreatedAt
	newUser.IsActive = user.IsActive
	newUser.Location = user.Location
	newUser.UpdatedAt = user.UpdatedAt
	newUser.RoleName = user.RoleName
	newUser.Sex = user.Sex
	newUser.HouseName = user.HouseName
	newUser.Location = user.Location

	return newUser

}

//ParseEvent -> Repo event model to api event Model
func ParseEvent(event repo.Event) *models.Event {
	newevent := &models.Event{}
	newevent.ID = event.ID.Hex()
	newevent.EventDate = &event.EventDate
	newevent.EventDescription = event.EventDescription
	newevent.EventLocation = event.EventLocation
	newevent.EventName = event.EventName
	newevent.IsActive = event.IsActive
	newevent.ApproverID = event.ApproverID
	newevent.CreatedAt = event.CreatedAt
	newevent.UpdatedAt = event.UpdatedAt
	newevent.IsVerified = event.IsVerified
	newevent.UserID = event.UserID

	return newevent
}

//ParsePost -> Repo post model to api post Model
func ParsePost(post repo.Post) *models.Post {
	newPost := &models.Post{}
	newPost.ID = post.ID.Hex()
	newPost.PostDate = post.PostDate
	newPost.PostDescription = post.PostDescription
	newPost.PostImgURL = post.PostImgURL
	newPost.PostVideoURL = post.PostVideoURL
	newPost.IsActive = post.IsActive
	newPost.CreatedAt = post.CreatedAt
	newPost.UpdatedAt = post.UpdatedAt
	newPost.IsVerified = post.IsVerified
	newPost.UserID = post.UserID

	return newPost
}

// SetEditUser -> Setter for the Edit stuff
func SetEditUser(req, dbUser *models.User) *models.User {

	if req.Username != "" {
		dbUser.Username = req.Username
	}
	if req.FirstName != "" {
		dbUser.FirstName = req.FirstName
	}
	if req.LastName != "" {
		dbUser.LastName = req.LastName
	}
	if req.MobileNo != "" {
		dbUser.MobileNo = req.MobileNo
	}
	if req.EmailAddress != "" {
		dbUser.EmailAddress = req.EmailAddress
	}
	if req.Password != "" {
		dbUser.Password = req.Password
	}
	if req.Description != "" {
		dbUser.Description = req.Description
	}
	if req.Dob != nil {
		dbUser.Dob = req.Dob
	}
	if req.RoleName != "" {
		dbUser.RoleName = req.RoleName
	}
	if req.Sex != "" {
		dbUser.Sex = req.Sex
	}
	if req.HouseName != "" {
		dbUser.HouseName = req.HouseName
	}
	if req.Location != "" {
		dbUser.Location = req.Location
	}
	if req.ID != "" {
		dbUser.ID = req.ID
	}
	dbUser.IsActive = req.IsActive
	dbUser.IsStayingHere = req.IsStayingHere

	return dbUser
}

// SetEditEvent -> Setter for the Edit stuff
func SetEditEvent(req, dbEvent *models.Event) *models.Event {

	if req.EventName != "" {
		dbEvent.EventName = req.EventName
	}
	if req.EventDescription != "" {
		dbEvent.EventDescription = req.EventDescription
	}
	if req.EventLocation != "" {
		dbEvent.EventLocation = req.EventLocation
	}
	if req.UserID != "" {
		dbEvent.UserID = req.UserID
	}
	if req.ApproverID != "" {
		dbEvent.ApproverID = req.ApproverID
	}
	if req.EventDate != nil {
		dbEvent.EventDate = req.EventDate
	}
	if req.ID != "" {
		dbEvent.ID = req.ID
	}

	dbEvent.IsActive = req.IsActive
	dbEvent.IsVerified = req.IsVerified

	return dbEvent
}

// SetEditPost -> Setter for the Post stuff
func SetEditPost(req, dbPost *models.Post) *models.Post {

	if req.PostDescription != "" {
		dbPost.PostDescription = req.PostDescription
	}
	if req.PostImgURL != "" {
		dbPost.PostImgURL = req.PostImgURL
	}
	if req.UserID != "" {
		dbPost.UserID = req.UserID
	}
	if req.PostDate != nil {
		dbPost.PostDate = req.PostDate
	}
	if req.ID != "" {
		dbPost.ID = req.ID
	}

	dbPost.IsActive = req.IsActive
	dbPost.IsVerified = req.IsVerified

	return dbPost
}

//EncryptPass -> hashes the password
func EncryptPass(pass []byte) (string, error) {
	var err error

	encryptPass, passErr := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if passErr != nil {
		err = fmt.Errorf("Failed Encrypting the password %w", passErr)
	}

	return string(encryptPass), err
}

//GetPwd -> gets the password in byte
func GetPwd(pwd string) []byte {
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)
}

//HashAndSalt -> hashes the password
func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

//ComparePasswords -> Check if the password is correct
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
