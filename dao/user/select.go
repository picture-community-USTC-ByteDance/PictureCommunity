package user

import (
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryIDAndPasswordByUsername(username string) (int64, string, error) {
	var user db.User
	err := global.MysqlDB.Select("uid", "password").Where("username=?", username).First(&user).Error
	return int64(user.UID), user.Password, err
}

func QueryIDAndPasswordByEmail(email string) (int64, string, error) {
	var user db.User
	err := global.MysqlDB.Select("uid", "password").Where("email=?", email).First(&user).Error
	return int64(user.UID), user.Password, err
}

func QueryIDAndPasswordByTelephone(telephone string) (int64, string, error) {
	var user db.User
	err := global.MysqlDB.Select("uid", "password").Where("telephone=?", telephone).First(&user).Error
	return int64(user.UID), user.Password, err
}

//搜索 根据关键字查询搜索到的用户列表 每个用户返回头像、昵称、个性签名、uid
func QueryUserListByNickname(keywords string, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
	global.MysqlDB.Model(db.UserDetail{}).
		Select("profile,uid,nickname,motto").
		Where("nickname like ?", keywords).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询他的关注用户列表 每个用户返回头像、昵称、个性签名、uid
func QueryFollowListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
	global.MysqlDB.Model(db.Follow{}).
		Select("profile,user_detail.uid,nickname,motto").
		Joins("inner join user_detail on follow.followed_id = user_detail.uid").
		Where("follow.uid = ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询他的粉丝用户列表 每个用户返回头像、昵称、个性签名、uid
func QueryFansListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
	global.MysqlDB.Model(db.Fans{}).
		Select("profile,user_detail.uid,nickname,motto").
		Joins("inner join user_detail on fans.fans_id = user_detail.uid").
		Where("fans.uid = ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询给他帖子点赞的用户列表 每个用户返回头像、昵称、个性签名、uid
func QueryLikeList1ByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
	global.MysqlDB.Model(db.Post{}).
		Select("profile,user_detail.uid,nickname,motto").
		Joins("inner join liked on post.p_id = liked.to_like_post_id").
		Joins("inner join user_detail on liked.from_user_id = user_detail.uid").
		Where("post.uid= ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}
