package user

import (
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryPasswordByID(id uint) (string, error) {
	var password string
	err := global.MysqlDB.Model(db.User{}).Select("password").Where("uid=?", id).Find(&password).Error
	return password, err
}

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
		Select("user.uid,user.username,user_detail.profile,user_detail.nickname,user_detail.motto").
		Joins("inner join user on user.uid= user_detail.uid").
		Where("nickname like ?", keywords).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询他的关注用户列表 每个用户返回头像、昵称、个性签名、uid
func QueryFollowListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
	global.MysqlDB.Model(db.Follow{}).
		Select("user_detail.uid,username,profile,nickname,motto").
		Joins("inner join user_detail on follow.followed_id = user_detail.uid").
		Joins("inner join user on user.uid= user_detail.uid").
		Where("follow.uid = ? AND follow.state= ?", uid, true).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询他的粉丝用户列表 每个用户返回头像、昵称、个性签名、uid
func QueryFansListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
  
	global.MysqlDB.Model(db.Follow{}).
		Select("user_detail.uid,username,profile,nickname,motto").
		Joins("inner join user_detail on follow.uid = user_detail.uid").
		Joins("inner join user on user.uid= user_detail.uid").
		Where("follow.followed_id = ? AND follow.state= ?", uid, true).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询给他帖子点赞的用户列表 每个用户返回头像、昵称、个性签名、uid
func QueryLikeList1ByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
	global.MysqlDB.Model(db.Post{}).
		Select("user_detail.uid,username,profile,nickname,motto").
		Joins("inner join liked on post.p_id = liked.to_like_post_id").
		Joins("inner join user_detail on liked.from_user_id = user_detail.uid").
		Joins("inner join user on user.uid= user_detail.uid").
		Where("post.uid= ? AND liked.state=? ", uid, true).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

func QueryUserDetailByUID(uid uint) (_response.UserDetail, error) {
	var userDetail _response.UserDetail
	err := global.MysqlDB.Model(db.UserDetail{}).
		Select("user_detail.uid,nickname,sex,birthday,address,motto,profile,origin_profile,user.email,user.telephone").
		Joins("inner join user on user_detail.uid = user.uid").
		Where("user_detail.uid=?", uid).First(&userDetail).Error
	return userDetail, err
}

func QueryChatListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseChatUsers) {
	var searchUsers []_response.ResponseChatUsers
	var count int64

	sql := "select uid,profile,nickname,unread from user_detail a join " +
		"(select from_id,count((is_read=0) OR null) as unread from chat_message where to_id=?  group by from_id union " +
		"select to_id,count(null) as unread from chat_message where  from_id=?  AND to_id not in (select from_id from chat_message where to_id=? group by from_id)  group by to_id) " +
		"as b on a.uid=b.from_id  order by unread desc  limit ? offset ?"
	sql2 := "select count(*) from " + "(" + "select uid,profile,nickname from user_detail where  uid in " +
		"(select distinct to_id from chat_message where from_id=? union  select distinct from_id from chat_message where to_id=?)" + ") as t"

	global.MysqlDB.Raw(sql2, uid, uid).Scan(&count).Raw(sql, uid, uid, uid, pageSize, (page-1)*pageSize).Scan(&searchUsers)
	return count, searchUsers
}

func QueryHistoryMsg(uid uint, toId uint) (int64, db.UserDetail, []_response.ResponseHistoryMsg) {
	var HistoryMsgList []_response.ResponseHistoryMsg
	var count int64
	global.MysqlDB.Model(db.ChatMessage{}).
		Select("from_id,content,created_at").
		Where("(from_id= ? AND to_id=?) OR (from_id=? AND to_id= ?)", uid, toId, toId, uid).
		Count(&count).Order("created_at").Scan(&HistoryMsgList)
	var userInfo db.UserDetail
	global.MysqlDB.First(&userInfo, toId)

	global.MysqlDB.Model(&db.ChatMessage{}).
		Where("to_id= ? AND from_id= ?", uid, toId).
		Update("is_read", 1)
	return count, userInfo, HistoryMsgList
}

func QueryUnreadMsg(uid uint) int64 {
	var count int64
	global.MysqlDB.Model(db.ChatMessage{}).
		Where(" to_id= ? AND is_read=?", uid, 0).
		Count(&count)

	return count
}
