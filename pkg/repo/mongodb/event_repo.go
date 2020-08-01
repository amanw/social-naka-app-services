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

//EventCollection -> const for the event doc
const EventCollection = "Events"

// CreateNewEvent -> creates the event in the DB
func (db *MongoDB) CreateNewEvent(user *models.Event) (*string, error) {
	user.CreatedAt = time.Now().String()
	result, err := db.DB.Collection(EventCollection).InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error while inserting new event into db, Reason: %v\n", err)
		return nil, err
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	strID := objectID.Hex()
	decodedEvent, eventErr := db.GetSingleEvent(strID)
	if eventErr != nil {
		log.Printf("Error while getting a single event, Reason: %v\n", eventErr)
	}
	decodedEvent.ID = strID
	str, err := db.EditEvent(decodedEvent)
	if err != nil {
		log.Printf("Edit Error, Reason: %v\n", err)
	}
	fmt.Println(str)
	return &strID, nil
}

//GetAllEvents -> Get all the events from the db
func (db *MongoDB) GetAllEvents() ([]*models.Event, error) {

	var event *models.Event
	events := []*models.Event{}

	cursor, err := db.DB.Collection(EventCollection).Find(context.Background(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)

		return nil, err
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.Background()) {
		var repoEvent *repo.Event
		cursor.Decode(&repoEvent)
		event = utils.ParseEvent(*repoEvent)
		events = append(events, event)
	}

	return events, nil
}

// GetSingleEvent -> Get the Single event from the db.
func (db *MongoDB) GetSingleEvent(id string) (*models.Event, error) {
	event := &models.Event{}
	repoEvent := &repo.Event{}
	docID, docErr := primitive.ObjectIDFromHex(id)
	if docErr != nil {
		log.Printf("Error while getting the DOCID, Reason: %v\n", docErr)
	}
	result := db.DB.Collection(EventCollection).FindOne(context.Background(), bson.M{"_id": docID})
	if result.Err() == nil {
		result.Decode(&repoEvent)
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	if repoEvent != nil {
		event = utils.ParseEvent(*repoEvent)
	}

	return event, nil
}

//EditEvent -> Edits the user info the db
func (db *MongoDB) EditEvent(req *models.Event) (*string, error) {
	eventID := req.ID
	docID, docErr := primitive.ObjectIDFromHex(eventID)
	if docErr != nil {
		fmt.Printf("Error while getting the DOCID, Reason: %v\n", docErr)
	}

	if req != nil {
		eventObj, err := db.GetSingleEvent(req.ID)
		if err != nil {
			logrus.Printf("Error while getting the DOCID, Reason: %v\n", err)
		}
		req = utils.SetEditEvent(req, eventObj)
		logrus.Printf("req, %+v", req)
		req.CreatedAt = *&eventObj.CreatedAt
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

	result := db.DB.Collection(EventCollection).FindOneAndUpdate(context.Background(), filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var v interface{}
	logrus.Printf("RESULT EVENT: %+v", result.Decode(v))
	doc := repo.Event{}
	decodeErr := result.Decode(&doc)
	if decodeErr != nil {
		logrus.Errorf("Error while decoding the Event object %+v", decodeErr)
	}

	fmt.Printf("%+v", doc)
	ID := doc.ID.Hex()

	return &ID, nil
}

// DeleteEvent -> Deletes the event from the db
func (db *MongoDB) DeleteEvent(id string) (bool, error) {

	result, err := db.DB.Collection(EventCollection).DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Printf("Error while deleting a single Event, Reason: %v\n", err)
		return false, err
	}
	if result.DeletedCount == 1 {
		return true, nil
	}
	return false, nil
}
