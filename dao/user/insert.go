package user

import (
	"picture_community/entity/db"
	"picture_community/global"
	"time"
)

func InsertUser(newUser db.User) error {
	//todo:中途出错要回滚
	err := global.MysqlDB.Create(&newUser).Error
	if err != nil {
		return err
	}

	newUserDetail := db.UserDetail{
		//todo: UID生成
		UID:           newUser.UID,
		Nickname:      "",
		Sex:           false,
		Birthday:      time.Now(),
		Address:       "",
		Motto:         "",
		Profile:       "",
		OriginProfile: "",
	}

	err = global.MysqlDB.Create(&newUserDetail).Error
	if err != nil {
		return err
	}

	newUserData := db.UserData{
		UID:              newUser.UID,
		FollowersNumber:  0,
		FansNumber:       0,
		PostsNumber:      0,
		CollectionNumber: 0,
		ForwardNumber:    0,
	}

	err = global.MysqlDB.Create(&newUserData).Error
	if err != nil {
		return err
	}
	return nil
}
