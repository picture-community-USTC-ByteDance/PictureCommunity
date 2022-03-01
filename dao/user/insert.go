package user

import (
	"gorm.io/gorm"
	"picture_community/entity/db"
	"picture_community/global"
)

func InsertUser(newUser db.User) error {
	return global.MysqlDB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&newUser).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		defaultProfile := "http://" + global.ServerName + ":8080/upload/pictures/default.jpg"
		if err := tx.Create(&db.UserDetail{UID: newUser.UID, Nickname: newUser.Username, Profile: defaultProfile}).Error; err != nil {
			return err
		}

		if err := tx.Create(&db.UserData{UID: newUser.UID}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}
