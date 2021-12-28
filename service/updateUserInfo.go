package service

import (
	userupdate "picture_community/dao/user"
	"picture_community/entity"
)

func UpdateUserInfo(parm entity.UpdatePost) (status int, message string) {

	newUserDetail := entity.UserDetail{
		ID:       parm.ID,
		NickName: parm.Content.Nickname,
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
