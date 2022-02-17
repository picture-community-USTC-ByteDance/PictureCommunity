package user

import (
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

// 在用户主页，查自己的用户信息
func QueryUserDataByUserID(uid uint) (error, _response.UserData) {
	var userdata _response.UserData
	// 在db.UserData{}数据库中查询，查询结果保存到结构体变量userdata中
	err := global.MysqlDB.Model(db.UserData{}).
		Select("nickname,sex,profile,motto,followers_number,fans_number,posts_number,collection_number,forward_number").
		Joins("inner join user_detail on user_data.uid = user_detail.uid").
		Where("user_detail.uid=?", uid).Scan(&userdata).Error

	return err, userdata
}

// 在用户主页，查自己的帖子数组信息
func QueryUserPostsByUserID(uid uint, page int, pageSize int) (int64, []_response.UserPosts) {
	var userposts []_response.UserPosts
	var count int64
	// 在db.Post{}数据库中查询，查询结果保存到结构体变量userposts中
	global.MysqlDB.Model(&db.Post{}).
		Select("p_id,title_photo_url,like_number,comment_number,forward_number").
		Where("uid=?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&userposts)

	return count, userposts
}
