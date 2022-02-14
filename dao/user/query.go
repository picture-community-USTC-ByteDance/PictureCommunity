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
		Select("user_data.uid,nickname,user.username,sex,profile,motto,followers_number,fans_number,posts_number,collection_number,forward_number").
		Joins("left join user_detail on user_data.uid = user_detail.uid").
		Joins("left join user on user.uid = user_data.uid").
		Where("user_data.uid=?", uid).Scan(&userdata).Error

	return err, userdata
}

// 根据username查其他用户的信息
func QueryUserDataByUsername(username string) (error, _response.UserData) {
	var userdata _response.UserData

	err := global.MysqlDB.Model(db.User{}).
		Select("user.uid,nickname,user.username,sex,profile,motto,followers_number,fans_number,posts_number,collection_number,forward_number").
		Joins("left join user_detail on user.uid = user_detail.uid").
		Joins("left join user_data on user.uid = user_data.uid").
		Where("user.username=?", username).Scan(&userdata).Error

	return err, userdata
}

// 查当前用户与指定用户的关注和被关注情况
func QueryIsFollowAndFan(username string, uid uint) (int64, int64) {
	// 若查成功，则countFollow大于0，即当前用户(uid) 的关注者包括 指定用户(username)
	var countFollow int64
	global.MysqlDB.Model(db.Follow{}).
		Select("follow.id").
		Joins("inner join user on user.uid = follow.followed_id").
		Where("user.username = ? AND follow.uid = ?", username, uid).Count(&countFollow)

	// 若查成功，则countFan大于0，即当前用户(uid) 的粉丝包括 指定用户(username)
	var countFan int64
	global.MysqlDB.Model(db.Fans{}).
		Select("fans.id").
		Joins("inner join user on user.uid = fans.fans_id").
		Where("user.username = ? AND fans.uid = ?", username, uid).Count(&countFan)

	return countFollow, countFan
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
