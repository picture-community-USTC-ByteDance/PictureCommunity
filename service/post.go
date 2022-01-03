package service

import (
	"mime/multipart"
	"net/http"
	"path"
	"picture_community/dao/post"
	"picture_community/entity/db"
	"picture_community/response"
	"picture_community/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	StorageLocation = "storage"
	ServerName      = "localhost"
)

func CreatePost(c *gin.Context, id uint, file *multipart.FileHeader, content string) response.ResStruct {
	utils.PathExists(StorageLocation)

	filesuffix := path.Ext(file.Filename)
	file.Filename = strconv.FormatUint(uint64(id), 10) + strconv.FormatInt(time.Now().Unix(), 10) + utils.RandStr(20) + filesuffix
	dst := path.Join(StorageLocation, file.Filename)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusGatewayTimeout,
			Code:       response.FailCode,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	dst = "http://" + ServerName + ":8080/upload/pictures/" + file.Filename
	newPost := db.Post{
		UID:              id,
		TitlePhotoUrl:    dst,
		Content:          content,
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
			Data:       nil,
		}
	}
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"post_id": postID},
	}
}
