package service

import (
	"errors"
	"net/http"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"

	"gorm.io/gorm"
)

func FollowUser(id uint, followedUID uint) response.ResStruct {
	var user db.User

	tx := global.MysqlDB.Begin()
	err := tx.Where("uid = ?", followedUID).First(&user).Error
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	var _follow db.Follow
	err = tx.Where("uid = ? and followed_id = ?", id, followedUID).First(&_follow).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    "Already followed!",
			Data:       nil,
		}
	}

	follow := db.Follow{
		UID:        id,
		FollowedID: followedUID,
		State:      true,
	}
	err = tx.Create(&follow).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	err = tx.Model(&db.UserData{}).Where("uid = ?", followedUID).Update("fans_number", gorm.Expr("fans_number + 1")).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	err = tx.Model(&db.UserData{}).Where("uid = ?", id).Update("followers_number", gorm.Expr("followers_number + 1")).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	tx.Commit()
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Message:    "follow success",
		Data:       nil,
	}
}

func UnfollowUser(id uint, followedUID uint) response.ResStruct {
	tx := global.MysqlDB.Begin()

	var follow db.Follow
	err := tx.Where("uid = ? and followed_id = ?", id, followedUID).First(&follow).Error
	if err != nil {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	err = tx.Where("uid = ? and followed_id = ?", id, followedUID).Delete(db.Follow{}).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	err = tx.Model(&db.UserData{}).Where("uid = ?", followedUID).Update("fans_number", gorm.Expr("fans_number - 1")).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	err = tx.Model(&db.UserData{}).Where("uid = ?", id).Update("followers_number", gorm.Expr("followers_number - 1")).Error
	if err != nil {
		tx.Rollback()
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       300,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	tx.Commit()
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Message:    "Unfollow success",
		Data:       nil,
	}
}
