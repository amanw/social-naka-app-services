package repo

import (
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Object for DB access
type User struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// dob
	// Format: date
	Dob strfmt.Date `json:"dob,omitempty"`

	// email address
	// Required: true
	EmailAddress string `json:"email_address"`

	// first name
	// Required: true
	FirstName string `json:"first_name"`

	// id
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	// is active
	IsActive bool `json:"is_active,omitempty"`

	// is staying here
	IsStayingHere bool `json:"is_staying_here,omitempty"`

	// last name
	// Required: true
	LastName string `json:"last_name"`

	// location
	Location string `json:"location,omitempty"`

	// mobile no
	// Required: true
	MobileNo string `json:"mobile_no"`

	// password
	// Required: true
	Password string `json:"password"`

	// role name
	RoleName string `json:"role_name,omitempty"`

	// sex
	Sex string `json:"sex,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// username
	// Required: true
	Username string `json:"username"`

	// house name
	HouseName string `json:"house_name,omitempty"`
}
