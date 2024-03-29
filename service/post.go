package service

import (
	"errors"
	"mime/multipart"
	"net/http"
	"path"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"
	"picture_community/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FileUpload(c *gin.Context, id uint, file *multipart.FileHeader) response.ResStruct {
	utils.PathExists(global.FileStorageLocation)

	filesuffix := path.Ext(file.Filename)
	file.Filename = strconv.FormatUint(uint64(id), 10) + strconv.FormatInt(time.Now().Unix(), 10) + utils.RandStr(20) + filesuffix
	dst := path.Join(global.FileStorageLocation, file.Filename)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusGatewayTimeout,
			Code:       response.FailCode,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	dst = "http://" + global.ServerName + ":8080/upload/pictures/" + file.Filename
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"url": dst},
	}
}

func CreatePost(c *gin.Context, id uint, urls []string, content string) response.ResStruct {
	len := len(urls)
	if len <= 0 || len > 10 {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    "number of photos cannot exceed 10 per post",
			Data:       nil,
		}
	}

	tx := global.MysqlDB.Begin()

	newPost := db.Post{
		PID:              global.PostIDGenerator.NewID(),
		UID:              id,
		TitlePhotoUrl:    urls[0],
		Content:          content,
		PhotoNumber:      len,
		CommentNumber:    0,
		LikeNumber:       0,
		ForwardNumber:    0,
		CollectionNumber: 0,
	}
	err := tx.Create(&newPost).Error
	postID := newPost.PID
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusGatewayTimeout,
			Code:       response.FailCode,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	postPhoto := db.PostPhoto{
		PID:          postID,
		UID:          id,
		Url:          urls[0],
		SerialNumber: 0,
	}
	for i, url := range urls {
		postPhoto.SerialNumber = uint(i)
		postPhoto.Url = url
		postPhoto.ID = global.PostIDGenerator.NewID()
		err = tx.Create(&postPhoto).Error
		if err != nil {
			tx.Rollback()
			return response.ResStruct{
				HttpStatus: http.StatusGatewayTimeout,
				Code:       response.FailCode,
				Message:    err.Error(),
				Data:       nil,
			}
		}
	}

	tx.Model(&db.UserData{}).Where("uid = ?", id).Update("posts_number", gorm.Expr("posts_number + 1"))

	tx.Commit()
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"post_id": postID},
	}
}

func DeletePost(uid uint, pid uint) response.ResStruct {
	tx := global.MysqlDB.Begin()

	post := db.Post{
		PID: pid,
	}

	err := tx.Where("uid = ?", uid).First(&post).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    err.Error(),
		}
	}

	err = tx.Where("uid = ?", uid).Delete(&post).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    err.Error(),
		}
	}

	err = tx.Where("p_id = ?", pid).Delete(&db.PostPhoto{}).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusInternalServerError,
			Code:       response.FailCode,
			Message:    err.Error(),
		}
	}

	tx.Model(&db.UserData{}).Where("uid = ?", uid).Update("posts_number", gorm.Expr("posts_number - 1"))

	tx.Commit()
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "delete success",
		Data:       nil,
	}
}

func NewForward(uid uint, pid uint, content string) response.ResStruct {
	var forward db.Forward
	var post db.Post
	err := global.MysqlDB.Where("p_id = ?", pid).First(&post).Error
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Message:    "Target post not exist",
			Data:       err.Error(),
		}
	}

	err = global.MysqlDB.Where("author_user_id = ? AND to_forward_post_id = ?", uid, pid).First(&forward).Error
	if err == nil {
		return response.ResStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Message:    "Already forwarded",
			Data: gin.H{
				"forward_id": forward.FID,
			},
		}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		forward = db.Forward{
			AuthorUserID:    uid,
			ToForwardPostID: pid,
			CommentNumber:   0,
			LikeNumber:      0,
			Content:         content,
		}
		err = global.MysqlDB.Create(&forward).Error
	}

	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusForbidden,
			Code:       response.FailCode,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	global.MysqlDB.Model(&db.UserData{}).Where("uid = ?", uid).Update("forward_number", gorm.Expr("forward_number + 1"))
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "Forward success",
		Data: gin.H{
			"forward_id": forward.FID,
		},
	}
}

func DeleteForward(uid uint, fid uint) response.ResStruct {
	forward := db.Forward{
		FID: fid,
	}

	err := global.MysqlDB.Where("author_user_id = ?", uid).First(&forward).Error
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    err.Error(),
		}
	}

	err = global.MysqlDB.Where("author_user_id = ?", uid).Delete(&forward).Error
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    err.Error(),
		}
	}

	global.MysqlDB.Model(&db.UserData{}).Where("uid = ?", uid).Update("forward_number", gorm.Expr("forward_number - 1"))

	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "delete success",
		Data:       nil,
	}
}
