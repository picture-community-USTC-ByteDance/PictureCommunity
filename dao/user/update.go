package user

import (
	"fmt"
	"picture_community/entity"
	"picture_community/global"
	"time"
)

func UpdateUserDetailByID(UserDetail entity.UserDetail, Id int64) error {
	res, err := global.MYSQL_DB.DB().Exec("update user_details set nickname=?,sex=?,updatedate=?,birthday=?,address=?,motto=?,profile=? where u_id=?", UserDetail.Nickname, UserDetail.Sex, time.Now(), UserDetail.Birthday, UserDetail.Address, UserDetail.Motto, UserDetail.Profile, Id)
	if err != nil {
		fmt.Println("Update Failed", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Rows Failed", err)
		return err
	}
	fmt.Println("Update succ: ", row)
	return err
}
