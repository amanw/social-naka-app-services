package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	"github.com/amanw/social-naka-app-services/pkg/api/utils"
	repo "github.com/amanw/social-naka-app-services/pkg/repo/models"
	"github.com/sirupsen/logrus"
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
		userStatus, checkErr = db.CheckUser(user.Username, user.EmailAddress, user.MobileNo)
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
		decodedUser.ID = strID
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
	users := []*models.User{}

	cursor, err := db.DB.Collection(UserCollection).Find(context.Background(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)

		return nil, err
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.Background()) {
		var repoUser *repo.User
		cursor.Decode(&repoUser)
		user = utils.ParseUser(*repoUser)
		users = append(users, user)
	}

	return users, nil
}

// GetSingleUser -> Get the Single user from the db.
func (db *MongoDB) GetSingleUser(id string) (*models.User, error) {
	user := &models.User{}
	repoUser := &repo.User{}
	docID, docErr := primitive.ObjectIDFromHex(id)
	if docErr != nil {
		log.Printf("Error while getting the DOCID, Reason: %v\n", docErr)
	}
	result := db.DB.Collection(UserCollection).FindOne(context.Background(), bson.M{"_id": docID})
	if result.Err() == nil {
		result.Decode(&repoUser)
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	if repoUser != nil {
		user = utils.ParseUser(*repoUser)
	}

	return user, nil
}

//EditUser -> Edits the user info the db
func (db *MongoDB) EditUser(req *models.User) (*string, error) {
	var userStatus string
	var checkErr error
	if req != nil {
		userObj, err := db.GetSingleUser(req.ID)
		req = utils.SetEditUser(req, userObj)
		logrus.Printf("req, %+v", req)
		req.CreatedAt = *&userObj.CreatedAt
		if err != nil {
			log.Printf("Error while getting the user: %v", err)
		}
		if req.Username != userObj.Username {
			userStatus, checkErr = db.CheckUser(req.Username, "", "")
		} else if req.EmailAddress != userObj.EmailAddress {
			userStatus, checkErr = db.CheckUser("", req.EmailAddress, "")

		} else if req.MobileNo != userObj.MobileNo {
			userStatus, checkErr = db.CheckUser("", "", req.MobileNo)
		} else {
			userStatus = "NE"
		}
		if checkErr != nil {
			log.Printf("Error while checking if the user exists, Reason: %v\n", checkErr)
		}
	}

	if userStatus == "NE" {
		userID := req.ID
		docID, docErr := primitive.ObjectIDFromHex(userID)
		// user, err := db.GetSingleUser(req.ID)
		// if err != nil {
		// 	logrus.Errorf("ERROR, %+v", err)
		// }
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
		doc := repo.User{}
		decodeErr := result.Decode(&doc)
		if decodeErr != nil {
			logrus.Errorf("Error while decoding the object %+v", decodeErr)
		}

		fmt.Printf("%+v", doc)
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

//LoginUser -> Checks if username, email or mobileno with password exists
func (db *MongoDB) LoginUser(loginModel *models.LoginUser) (bool, error) {
	passwordFilter := bson.M{
		"$or": []bson.M{
			{"username": *loginModel.Username},
			{"emailaddress": *loginModel.Username},
			{"mobileno": *loginModel.Username},
		},
	}
	result := db.DB.Collection(UserCollection).FindOne(context.Background(), passwordFilter)
	if result.Err() != nil {
		return false, result.Err()
	}
	doc := repo.User{}
	decodeErr := result.Decode(&doc)
	if decodeErr != nil {
		logrus.Errorf("Error while decoding the object %+v", decodeErr)
	}
	pass := utils.GetPwd(*loginModel.Password)
	checkPassword := utils.ComparePasswords(doc.Password, pass)
	if checkPassword {
		logrus.Println("The password has been matched")
	}

	filter := bson.M{
		"$or": []bson.M{
			{"username": *loginModel.Username},
			{"emailaddress": *loginModel.Username},
			{"mobileno": *loginModel.Username},
		},
		"$and": []bson.M{
			{"isactive": true},
		},
	}

	loginCount, err := db.DB.Collection(UserCollection).CountDocuments(context.Background(), filter)
	logrus.Println("Login", loginCount)
	if err != nil {
		return false, err
	}
	if loginCount == 1 {
		return true, nil
	}

	return false, nil
}
