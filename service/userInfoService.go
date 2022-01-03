package service

import (
	"fmt"
	userupdate "picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/db"
	"strconv"
)

func UpdateUserInfo(parm _request.UpdatePost) (status int, message string) {
	fmt.Println("contents::", parm.ID, parm.Content)

	newUserDetail := db.UserDetail{
		//ID :666
		Nickname: parm.Content.Nickname,
		Sex:      parm.Content.Sex,
		Birthday: parm.Content.Birthday,
		Address:  parm.Content.Address,
		Motto:    parm.Content.Motto,
		Profile:  parm.Content.Profile,
	}

	//_, err := daopost.InsertPostByUserID(newPost)
	err := userupdate.UpdateUserDetailByID(newUserDetail, parm.ID)
	if err != nil {
		return -1, err.Error()
	}
	return 0, "update user information success"
}

func EmailIsUnique(param _request.EmailIsUniqueInfo) (isOK bool, message string) {
	ID, _ := userupdate.QueryIDAndPasswordByEmail(param.Email)

	if ID != 0 {
		return false, "email重复"
	} else {
		return true, "email可用"
	}
}

func TelephoneIsUnique(param _request.TelephoneIsUniqueInfo) (isOK bool, message string) {
	ID, _ := userupdate.QueryIDAndPasswordByTelephone(strconv.Itoa(int(param.Telephone)))

	if ID != 0 {
		return false, "电话号重复"
	} else {
		return true, "电话号可用"
	}
}
