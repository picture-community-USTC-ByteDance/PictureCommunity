package service

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"picture_community/entity/_request"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"
)

func SearchService(keywords string, pageSize int, page int) response.ResStruct {
	var searchUsers []_request.SearchUsers
	var count int64

	global.MysqlDB.Offset((page-1)*pageSize).Limit(pageSize).Model(db.User{}).
		Select("profile,username,nickname,motto").
		Joins("inner join csm_detail on user.id = user_detail.id").
		Where("nickname like ?", keywords).Count(&count).Scan(&searchUsers)
	if count == 0 {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    "搜索用户不存在",
			Data:       nil,
		}
	}

	totalPage := math.Ceil(float64(int(count) / pageSize))

	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"User": searchUsers, "totalPage": totalPage},
	}
}
