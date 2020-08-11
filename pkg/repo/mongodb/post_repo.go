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

//PostCollection -> const for the post doc
const PostCollection = "Posts"

// CreateNewPost -> creates the post in the DB
func (db *MongoDB) CreateNewPost(post *models.Post) (*string, error) {
	post.CreatedAt = time.Now().String()
	result, err := db.DB.Collection(PostCollection).InsertOne(context.Background(), post)
	if err != nil {
		log.Printf("Error while inserting new post into db, Reason: %v\n", err)
		return nil, err
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	strID := objectID.Hex()
	decodedPost, postErr := db.GetSinglePost(strID)
	if postErr != nil {
		log.Printf("Error while getting a single post, Reason: %v\n", postErr)
	}
	decodedPost.ID = strID
	str, err := db.EditPost(decodedPost)
	if err != nil {
		log.Printf("Edit Error, Reason: %v\n", err)
	}
	fmt.Println(str)
	return &strID, nil
}

//GetAllPosts -> Get all the posts from the db
func (db *MongoDB) GetAllPosts() ([]*models.Post, error) {

	var post *models.Post
	posts := []*models.Post{}

	cursor, err := db.DB.Collection(PostCollection).Find(context.Background(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all posts, Reason: %v\n", err)

		return nil, err
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.Background()) {
		var repoPost *repo.Post
		cursor.Decode(&repoPost)
		post = utils.ParsePost(*repoPost)
		posts = append(posts, post)
	}

	return posts, nil
}

// GetSinglePost -> Get the Single post from the db.
func (db *MongoDB) GetSinglePost(id string) (*models.Post, error) {
	post := &models.Post{}
	repoPost := &repo.Post{}
	docID, docErr := primitive.ObjectIDFromHex(id)
	if docErr != nil {
		log.Printf("Error while getting the DOCID, Reason: %v\n", docErr)
	}
	result := db.DB.Collection(PostCollection).FindOne(context.Background(), bson.M{"_id": docID})
	if result.Err() == nil {
		result.Decode(&repoPost)
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	if repoPost != nil {
		post = utils.ParsePost(*repoPost)
	}

	return post, nil
}

//EditPost -> Edits the post info the db
func (db *MongoDB) EditPost(req *models.Post) (*string, error) {
	postID := req.ID
	docID, docErr := primitive.ObjectIDFromHex(postID)
	if docErr != nil {
		fmt.Printf("Error while getting the DOCID, Reason: %v\n", docErr)
	}

	if req != nil {
		postObj, err := db.GetSinglePost(req.ID)
		if err != nil {
			logrus.Printf("Error while getting the DOCID, Reason: %v\n", err)
		}
		req = utils.SetEditPost(req, postObj)
		logrus.Printf("req, %+v", req)
		req.CreatedAt = *&postObj.CreatedAt
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

	result := db.DB.Collection(PostCollection).FindOneAndUpdate(context.Background(), filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var v interface{}
	logrus.Printf("RESULT POST: %+v", result.Decode(v))
	doc := repo.Event{}
	decodeErr := result.Decode(&doc)
	if decodeErr != nil {
		logrus.Errorf("Error while decoding the Event object %+v", decodeErr)
	}

	fmt.Printf("%+v", doc)
	ID := doc.ID.Hex()

	return &ID, nil
}

// DeletePost -> Deletes the post from the db
func (db *MongoDB) DeletePost(id string) (bool, error) {

	result, err := db.DB.Collection(PostCollection).DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Printf("Error while deleting a single post, Reason: %v\n", err)
		return false, err
	}
	if result.DeletedCount == 1 {
		return true, nil
	}
	return false, nil
}
