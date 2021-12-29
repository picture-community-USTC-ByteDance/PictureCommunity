package service

import (
	"fmt"
	userupdate "picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/db"
	"time"
)

func UpdateUserInfo(parm _request.UpdatePost) (status int, message string) {
	fmt.Println("contents::", parm.ID, parm.Content)

	newUserDetail := db.UserDetail{
		UID:        parm.ID,
		UpdateDate: time.Now(),
		Nickname:   parm.Content.Nickname,
		Sex:        parm.Content.Sex,
		Birthday:   parm.Content.Birthday,
		Address:    parm.Content.Address,
		Motto:      parm.Content.Motto,
		Profile:    parm.Content.Profile,
	}

	//_, err := daopost.InsertPostByUserID(newPost)
	err := userupdate.UpdateUserDetailByID(newUserDetail, parm.ID)
	if err != nil {
		return -1, err.Error()
	}
	return 0, "update user information success"
}
