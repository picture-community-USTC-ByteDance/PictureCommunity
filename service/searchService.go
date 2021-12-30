package service

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"
)

func SearchService(keywords string, pageSize int, page int) response.ResStruct {
	var searchUsers []_response.ResponseSearchUsers

	global.MysqlDB.Offset((page-1)*pageSize).Limit(pageSize).Model(db.UserDetail{}).
		Select("profile,user_name,nick_name,motto").
		Joins("inner join users on user_details.uid = users.uid").
		Where("nick_name like ?", keywords).Scan(&searchUsers)

	count := len(searchUsers)
	if count == 0 {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    "搜索用户不存在",
			Data:       nil,
		}
	}

	totalPage := math.Ceil(float64(count) / float64(pageSize))

	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       gin.H{"User": searchUsers, "totalPage": totalPage},
	}
}
