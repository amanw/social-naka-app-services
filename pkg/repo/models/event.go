package repo

import (
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event -> Repo Object
type Event struct {

	// approver id
	ApproverID string `json:"approver_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// event date
	// Format: date
	EventDate strfmt.Date `json:"event_date,omitempty"`

	// event description
	EventDescription string `json:"event_description,omitempty"`

	// event location
	EventLocation string `json:"event_location,omitempty"`

	// event name
	EventName string `json:"event_name,omitempty"`

	// id
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	// is active
	IsActive bool `json:"is_active,omitempty"`

	// is verified
	IsVerified bool `json:"is_verified,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}
