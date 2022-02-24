package service

import (
	"net/http"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"
)

func FollowUser(id uint, followedUID uint) response.ResStruct {
	var user db.User

	err := global.MysqlDB.Where("uid = ?", followedUID).First(&user).Error
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	follow := db.Follow{
		UID:        id,
		FollowedID: followedUID,
		State:      true,
	}
	err = global.MysqlDB.Create(&follow).Error
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Message:    "follow success",
		Data:       nil,
	}
}

func UnfollowUser(id uint, followedUID uint) response.ResStruct {
	err := global.MysqlDB.Where("uid = ? and followed_id = ?", id, followedUID).Delete(db.Follow{}).Error

	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Message:    "Unfollow success",
		Data:       nil,
	}
}
