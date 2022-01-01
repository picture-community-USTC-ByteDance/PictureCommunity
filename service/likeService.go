package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picture_community/dao/like"
	"picture_community/entity/db"
	"picture_community/response"
	"time"
)

func CreateLike(u_id uint, post_id uint) response.ResStruct {
	newLike := db.Liked{
		ToLikePostID: u_id,
		FromUserID:   post_id,
		State:        true,
		UpdateTime:   time.Time{},
		CreateTime:   time.Time{},
	}
	likeID, err := like.InsertLikeByUserID(newLike)
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusGatewayTimeout,
			Code:       response.FailCode,
			Message:    err.Error(),
			Data:       gin.H{"like_id": likeID},
		}
	}
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"like_id": likeID},
	}
}

func QueryLike(u_id uint, post_id uint) response.ResStruct {
	res, _ := like.QueryLikeByUserID(u_id, post_id)
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"State": res},
	}
}
func CancelLike(u_id uint, post_id uint) response.ResStruct {
	state, _ := like.QueryLikeByUserID(u_id, post_id)
	if state == true {
		_ = like.CancelLikeByUserID(u_id, post_id)
	}
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       nil,
	}

}
