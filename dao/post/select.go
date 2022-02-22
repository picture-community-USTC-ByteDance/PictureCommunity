package post

import (
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

func InsertPostByUserID(newPost db.Post) (int64, error) {
	err := global.MysqlDB.Create(&newPost).Error
	return int64(newPost.PID), err
}

//个人主页里面 根据用户UID查询他的收藏帖子列表 每个帖子返回封面、pid、点赞数、评论数
func QueryHomeCollectionListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseHomePost) {
	var responsePost []_response.ResponseHomePost
	var count int64
	global.MysqlDB.Model(db.Collection{}).
		Select("post.p_id,title_photo_url,comment_number,like_number").
		Joins("inner join post on post.p_id = collection.p_id").
		Where("collection.uid= ? AND collection.state= ?", uid, true).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Order("collection.create_time desc").Scan(&responsePost)
	return count, responsePost
}

//个人主页里面 根据用户UID查询自己点赞的帖子列表 每个帖子返回封面、pid、点赞数、评论数
func QueryLikePostByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseHomePost) {
	var searchUsers []_response.ResponseHomePost
	var count int64
	global.MysqlDB.Model(db.Liked{}).
		Select("post.p_id,title_photo_url,comment_number,like_number").
		Joins("inner join post on post.p_id = liked.to_like_post_id").
		Where("liked.from_user_id= ? AND liked.state= ?", uid, true).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Order("liked.create_time desc").Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询发布过帖子列表 每个帖子返回封面、pid、点赞数、评论数
func QueryHomePostListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseHomePost) {
	var homePostList []_response.ResponseHomePost
	var count int64
	global.MysqlDB.Model(&db.Post{}).
		Select("p_id,title_photo_url,comment_number,like_number").
		Where("uid = ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Order("post.create_time desc").Scan(&homePostList)

	return count, homePostList
}
