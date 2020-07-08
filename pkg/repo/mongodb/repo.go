package repo

import (
	"context"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB defines the Client
type MongoDB struct {
	DB *mongo.Database
}

type User struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// dob
	// Required: true
	// Format: date
	Dob *strfmt.Date `json:"dob"`

	// email address
	// Required: true
	EmailAddress *string `json:"email_address"`

	// first name
	// Required: true
	FirstName *string `json:"first_name"`

	// id
	// Required: true
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	// is active
	IsActive bool `json:"is_active,omitempty"`

	// is staying here
	IsStayingHere bool `json:"is_staying_here,omitempty"`

	// last name
	// Required: true
	LastName *string `json:"last_name"`

	// location
	// Required: true
	Location *string `json:"location"`

	// mobile no
	// Required: true
	MobileNo *string `json:"mobile_no"`

	// password
	// Required: true
	Password *string `json:"password"`

	// role name
	RoleName string `json:"role_name,omitempty"`

	// sex
	// Required: true
	Sex *string `json:"sex"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// DbConnect is use to connect the DB
func DbConnect(url string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err == nil {
		log.Println("Connected!")
	} else {
		return nil, err
	}

	db := client.Database("social-naka-db")

	return &MongoDB{
		DB: db,
	}, nil
}
