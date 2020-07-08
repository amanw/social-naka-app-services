package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//UserCollection -> const for the user doc
const UserCollection = "Users"

// CreateNewUser -> creates the user in the DB
func (db *MongoDB) CreateNewUser(user *models.User) (*string, error) {
	var userStatus string
	var checkErr error
	if user != nil {
		userStatus, checkErr = db.CheckUser(*user.Username, *user.EmailAddress, *user.MobileNo)
		if checkErr != nil {
			log.Printf("Error while checking if the user exists, Reason: %v\n", checkErr)
		}
	}
	if userStatus == "NE" {
		user.CreatedAt = time.Now().String()
		result, err := db.DB.Collection(UserCollection).InsertOne(context.Background(), user)
		if err != nil {
			log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
			return nil, err
		}

		objectID := result.InsertedID.(primitive.ObjectID)
		strID := objectID.Hex()
		decodedUser, userErr := db.GetSingleUser(strID)
		if userErr != nil {
			log.Printf("Error while getting a single todo, Reason: %v\n", userErr)
		}
		decodedUser.ID = &strID
		str, err := db.EditUser(decodedUser)
		if err != nil {
			log.Printf("Edit Error, Reason: %v\n", err)
		}
		fmt.Println(str)
		return &strID, nil
	}
	return &userStatus, nil
}

//GetAllUsers -> Get all the users from the db
func (db *MongoDB) GetAllUsers() ([]*models.User, error) {

	var user *models.User
	var repoUser *User
	users := []*models.User{}

	cursor, err := db.DB.Collection(UserCollection).Find(context.Background(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)

		return nil, err
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.Background()) {
		cursor.Decode(&repoUser)
		if repoUser != nil {
			cursor.Decode(&user)
			ID := repoUser.ID.Hex()
			user.ID = &ID
		}
		users = append(users, user)
	}

	return users, nil
}

// GetSingleUser -> Get the Single user from the db.
func (db *MongoDB) GetSingleUser(id string) (*models.User, error) {

	user := &models.User{}
	repoUser := &User{}
	docID, docErr := primitive.ObjectIDFromHex(id)
	if docErr != nil {
		log.Printf("Error while getting the DOCID, Reason: %v\n", docErr)
	}
	result := db.DB.Collection(UserCollection).FindOne(context.Background(), bson.M{"_id": docID})
	if result.Err() == nil {
		result.Decode(&user)
		result.Decode(&repoUser)
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	if user != nil {
		ID := repoUser.ID.Hex()
		user.ID = &ID
	}

	return user, nil
}

//EditUser -> Edits the user info the db
func (db *MongoDB) EditUser(req *models.User) (*string, error) {
	var userStatus string
	var checkErr error
	if req != nil {
		userObj, err := db.GetSingleUser(*req.ID)
		req.CreatedAt = *&userObj.CreatedAt
		if err != nil {
			log.Printf("Error while getting the user: %v", err)
		}
		if *req.Username != *userObj.Username {
			userStatus, checkErr = db.CheckUser(*req.Username, "", "")
		} else if *req.EmailAddress != *userObj.EmailAddress {
			userStatus, checkErr = db.CheckUser("", *req.EmailAddress, "")

		} else if *req.MobileNo != *userObj.MobileNo {
			userStatus, checkErr = db.CheckUser("", "", *req.MobileNo)
		} else {
			userStatus = "NE"
		}
		if checkErr != nil {
			log.Printf("Error while checking if the user exists, Reason: %v\n", checkErr)
		}
	}

	if userStatus == "NE" {
		userID := *req.ID
		docID, docErr := primitive.ObjectIDFromHex(userID)

		if docErr != nil {
			fmt.Printf("Error while getting the DOCID, Reason: %v\n", docErr)
		}

		req.UpdatedAt = time.Now().String()

		filter := bson.M{"_id": docID}
		update := bson.M{
			"$set": req,
		}

		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		result := db.DB.Collection(UserCollection).FindOneAndUpdate(context.Background(), filter, update, &opt)
		if result.Err() != nil {
			return nil, result.Err()
		}
		doc := User{}
		decodeErr := result.Decode(&doc)
		if decodeErr != nil {
			fmt.Errorf("Error while decoding the object %+v", decodeErr)
		}

		fmt.Println("%+v", doc)
		ID := doc.ID.Hex()

		return &ID, nil
	}
	return &userStatus, nil
}

// DeleteUser -> Deletes the user from the db
func (db *MongoDB) DeleteUser(id string) (bool, error) {

	result, err := db.DB.Collection(UserCollection).DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Printf("Error while deleting a single Background, Reason: %v\n", err)
		return false, err
	}
	if result.DeletedCount == 1 {
		return true, nil
	}
	return false, nil
}

// CheckUser ->  Checks if the username or email id exists
func (db *MongoDB) CheckUser(userName, userEmail, mobileNo string) (string, error) {
	// filter := bson.M{
	// 	"$or": []bson.M{
	// 		bson.M{"username": userName},
	// 		bson.M{"email_address": userEmail},
	// 		bson.M{"mobile_no": mobileNo},
	// 	},
	// }
	// TODO: Find a single query to do down below (refering the above filter query commented)
	var status string
	status = "S"
	userNameCount, err := db.DB.Collection(UserCollection).CountDocuments(context.Background(), bson.M{"username": userName})
	emailCount, err := db.DB.Collection(UserCollection).CountDocuments(context.Background(), bson.M{"emailaddress": userEmail})
	mobileNoCount, err := db.DB.Collection(UserCollection).CountDocuments(context.Background(), bson.M{"mobileno": mobileNo})
	if userNameCount > 0 {
		status = status + "," + "UE"
	}
	if emailCount > 0 {
		status = status + "," + "EE"
	}
	if mobileNoCount > 0 {
		status = status + "," + "ME"
	}
	if mobileNoCount == 0 && emailCount == 0 && userNameCount == 0 {
		return "NE", nil
	}

	if err != nil {
		return "", err
	}

	return status, nil
}
