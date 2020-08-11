package handlers

import (
	"context"
	"fmt"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
	api "github.com/amanw/social-naka-app-services/pkg/api/restapi/operations/posts"
	"github.com/amanw/social-naka-app-services/pkg/domain"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// PostHandler : Object
type PostHandler struct {
	postService domain.PostSvc
}

//NewPostHandler : will be used to initialize the Post Handler Object
func NewPostHandler(postService domain.PostSvc) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

//CreatePost will enable the client to post the user to the DB
func (p PostHandler) CreatePost(ctx context.Context, params api.CreatePostParams) middleware.Responder {
	// fmt.Println(params.HTTPRequest.Header.Get("X-API-Key"))
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), "")
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewRegisterUserUnauthorized().WithPayload(&err)
	// }
	postRes := &models.PostResponse{}
	postStatus, err := p.postService.CreatePost(params.PostBody)

	fmt.Println("postStatus:", *postStatus)
	if err != nil {
		postErr := models.Response{
			Code:    400,
			Status:  "Failed",
			Message: err.Error(),
		}
		logrus.Warnf(postErr.Message)
		return api.NewCreatePostBadRequest().WithPayload(&postErr)
	}
	logrus.Infoln("The post has been created in the DB:", *postStatus)

	postRes.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "Post Created Successfully",
	}
	params.PostBody.ID = *postStatus
	postRes.Post = params.PostBody
	return api.NewCreatePostOK().WithPayload(postRes)
}

//GetPosts will enable the client to get all the posts from the DB
func (p PostHandler) GetPosts(ctx context.Context, params api.GetPostsParams) middleware.Responder {

	postResponse := &models.PostResponse{}
	postList, err := p.postService.GetPosts()
	if err != nil {
		postErr := fmt.Errorf("failed to get the posts: %w", err)
		logrus.Warnf(postErr.Error())
		responseErr := &models.Response{
			Status:  "Failed",
			Code:    400,
			Message: postErr.Error(),
		}
		return api.NewGetPostsBadRequest().WithPayload(responseErr)
	}

	logrus.Infoln("We get the Post List:", postList)
	postResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "Posts has been fetched Successfully",
	}
	postResponse.Posts = postList

	return api.NewGetPostsOK().WithPayload(postResponse)
}

//GetPostbyID pulls the post information By its id
func (p PostHandler) GetPostbyID(ctx context.Context, params api.GetPostbyIDParams) middleware.Responder {

	postResponse := &models.PostResponse{}
	post, err := p.postService.GetPostByID(params.ID)
	if err != nil {
		postErr := fmt.Errorf("failed to get the post: %w", err)
		logrus.Warnf(postErr.Error())
		getPostErr := &models.Response{
			Status:  "Failed",
			Code:    400,
			Message: postErr.Error(),
		}
		return api.NewGetPostbyIDBadRequest().WithPayload(getPostErr)

	}
	postResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "Post has been fetched Successfully",
	}
	postResponse.Post = post

	logrus.Infoln("The post has been fetched:", post)

	return api.NewGetPostbyIDOK().WithPayload(postResponse)
}

//UpdatePostbyID -> updates the post information By its id
func (p PostHandler) UpdatePostbyID(ctx context.Context, params api.UpdatePostbyIDParams) middleware.Responder {
	editResponse := &models.PostResponse{}
	params.PostBody.ID = params.ID
	editStatus, err := p.postService.UpdatePostByID(params.PostBody)

	fmt.Println("editStatus:", editStatus)
	if err != nil {
		editErr := models.Response{
			Code:    400,
			Status:  "Failed",
			Message: err.Error(),
		}
		logrus.Warnf(editErr.Message)
		return api.NewUpdatePostbyIDBadRequest().WithPayload(&editErr)
	}

	logrus.Infoln("The post has been updated in the DB:", *editStatus)
	params.PostBody.ID = *editStatus
	editResponse.Response = &models.Response{
		Code:    200,
		Status:  "Success",
		Message: "The post has been edited successfully",
	}
	editResponse.Post = params.PostBody
	return api.NewUpdatePostbyIDOK().WithPayload(editResponse)
}

//DeletePostbyID -> deletes the post information By its id
func (p PostHandler) DeletePostbyID(ctx context.Context, params api.DeletePostbyIDParams) middleware.Responder {
	// validateErr := u.userService.Validate(params.HTTPRequest.Header.Get("X-API-Key"), params.ID)
	// if validateErr != nil {
	// 	err := models.Error{
	// 		Message: validateErr.Error(),
	// 	}
	// 	return api.NewRegisterUserUnauthorized().WithPayload(&err)
	// }
	status, err := p.postService.DeletePostByID(params.ID)
	if err != nil {
		postErr := fmt.Errorf("Deletion of the post failed: %w", err)
		logrus.Warnf(postErr.Error())
		Errstatus := models.Response{}
		Errstatus.Code = 500
		Errstatus.Status = "Failed"
		Errstatus.Message = postErr.Error()
		return api.NewDeletePostbyIDInternalServerError().WithPayload(&Errstatus)
	}
	modelStatus := models.Response{}
	if status {
		logrus.Infoln("The post has been deleted in the DB:", status)
		modelStatus.Code = 200
		modelStatus.Status = "Success"
		modelStatus.Message = "The post has been deleted from the DB"
	}

	return api.NewDeletePostbyIDOK().WithPayload(&modelStatus)
}
