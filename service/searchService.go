package service

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"picture_community/entity"
	"picture_community/global"
	"picture_community/response"
)

func SearchService(keywords string, pageSize int, page int) response.ResStruct {
	var searchUsers []entity.ResponseSearchUsers

	global.MYSQL_DB.Offset((page-1)*pageSize).Limit(pageSize).Model(entity.UserDetail{}).
		Select("profile,user_name,nick_name,motto").
		Joins("inner join users on user_details.u_id = users.id").
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
