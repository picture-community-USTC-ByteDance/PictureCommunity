package service

import (
	"net/http"
	"picture_community/dao/post"
	"picture_community/entity"
	"picture_community/response"

	"github.com/gin-gonic/gin"
)

func CreatePost(parm entity.CreatePost) response.ResStruct {
	newPost := entity.Post{
		UID:              parm.ID,
		TitlePhotoUrl:    parm.ImageUrl,
		Content:          parm.Content,
		PhotoNumber:      1,
		CommentNumber:    0,
		LikeNumber:       0,
		ForwardNumber:    0,
		CollectionNumber: 0,
	}
	postID, err := post.InsertPostByUserID(newPost)
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusGatewayTimeout,
			Code:       response.FailCode,
			Message:    err.Error(),
			Data:       gin.H{"post_id": postID},
		}
	}
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"post_id": postID},
	}
}