package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	userupdate "picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/db"
	"strconv"
)

func UpdateUserDetailService(param _request.UpdateUserDetailInfo, uid uint) (isOK bool, message string) {

	newUserDetail := db.UserDetail{
		Nickname:      param.Nickname,
		Sex:           param.Sex,
		Birthday:      param.Birthday,
		Address:       param.Address,
		Motto:         param.Motto,
		Profile:       param.Profile,
		OriginProfile: param.OriginProfile,
	}

	err := userupdate.UpdateUserDetailByID(uid, newUserDetail)
	if err != nil {
		fmt.Println(err)
		return false, "更新用户信息失败"
	} else {
		return true, "更新用户信息成功"
	}
}

func EmailIsUniqueService(param _request.EmailIsUniqueInfo) (isOK bool, message string) {
	_, _, err := userupdate.QueryIDAndPasswordByEmail(param.Email)

	if err == nil {
		return false, "email重复"
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, "email可用"
	} else {
		fmt.Println(err)
		return false, "数据库错误"
	}
}

func TelephoneIsUniqueService(param _request.TelephoneIsUniqueInfo) (isOK bool, message string) {
	_, _, err := userupdate.QueryIDAndPasswordByTelephone(strconv.Itoa(int(param.Telephone)))

	if err == nil {
		return false, "电话号码重复"
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, "电话号码可用"
	} else {
		fmt.Println(err)
		return false, "数据库错误"
	}
}

func UpdateUserEmailService(param _request.UpdateUserEmailInfo, uid uint) (isOK bool, message string) {
	_, _, err := userupdate.QueryIDAndPasswordByEmail(param.Email)
	if err == nil {
		return false, "email重复"
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
		return false, "数据库错误"
	}

	err = userupdate.UpdateEmailByID(uid, param.Email)
	if err != nil {
		return false, "email更新失败"
	} else {
		return true, "email更新成功"
	}
}

func UpdateUserTelephoneService(param _request.UpdateUserTelephoneInfo, uid uint) (isOK bool, message string) {
	_, _, err := userupdate.QueryIDAndPasswordByTelephone(strconv.Itoa(int(param.Telephone)))
	if err == nil {
		return false, "电话号码重复"
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
		return false, "数据库错误"
	}

	err = userupdate.UpdateTelephoneByID(uid, param.Telephone)
	if err != nil {
		return false, "电话号码更新失败"
	} else {
		return true, "电话号码更新成功"
	}
}
