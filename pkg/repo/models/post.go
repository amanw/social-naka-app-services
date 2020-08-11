package repo

import (
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post Obj
type Post struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	// is active
	IsActive bool `json:"is_active,omitempty"`

	// is verified
	IsVerified bool `json:"is_verified,omitempty"`

	// post date
	// Format: date
	PostDate *strfmt.Date `json:"post_date,omitempty"`

	// post description
	PostDescription string `json:"post_description,omitempty"`

	// post img url
	PostImgURL string `json:"post_img_url,omitempty"`

	// post video url
	PostVideoURL string `json:"post_video_url,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}
