package post

import (
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

func InsertPostByUserID(newPost db.Post) (int64, error) {
	err := global.MysqlDB.Create(&newPost).Error
	return int64(newPost.UID), err
}

//个人主页里面 根据用户UID查询他的收藏帖子列表 每个帖子返回封面、pid
func QueryCollectionListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponsePost) {
	var responsePost []_response.ResponsePost
	var count int64
	global.MysqlDB.Model(db.Collection{}).
		Select("post.p_id,title_photo_url").
		Joins("inner join post on post.p_id = collection.p_id").
		Where("collection.uid= ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&responsePost)
	return count, responsePost
}

//个人主页里面 根据用户UID查询他的收藏帖子列表 每个帖子返回封面、pid、点赞数、评论数
func QueryHomeCollectionListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseHomePost) {
	var responsePost []_response.ResponseHomePost
	var count int64
	global.MysqlDB.Model(db.Collection{}).
		Select("post.p_id,title_photo_url,comment_number,like_number").
		Joins("inner join post on post.p_id = collection.p_id").
		Where("collection.uid= ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&responsePost)
	return count, responsePost
}

//个人主页里面 根据用户UID查询自己点赞的帖子列表 每个帖子返回封面、pid
func QueryLikeList2ByUID(uid uint, page int, pageSize int) (int64, []_response.ResponsePost) {
	var searchUsers []_response.ResponsePost
	var count int64
	global.MysqlDB.Model(db.Liked{}).
		Select("post.p_id,title_photo_url").
		Joins("inner join post on post.p_id = liked.to_like_post_id").
		Where("liked.from_user_id= ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询自己点赞的帖子列表 每个帖子返回封面、pid、点赞数、评论数
func QueryLikePostByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseHomePost) {
	var searchUsers []_response.ResponseHomePost
	var count int64
	global.MysqlDB.Model(db.Liked{}).
		Select("post.p_id,title_photo_url,comment_number,like_number").
		Joins("inner join post on post.p_id = liked.to_like_post_id").
		Where("liked.from_user_id= ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询发布过帖子列表 每个帖子返回封面、pid
func QueryPostListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponsePost) {
	var postList []_response.ResponsePost
	var count int64
	global.MysqlDB.Model(&db.Post{}).
		Select("p_id,title_photo_url").
		Where("uid = ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&postList)

	return count, postList
}

//个人主页里面 根据用户UID查询发布过帖子列表 每个帖子返回封面、pid、点赞数、评论数
func QueryHomePostListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponseHomePost) {
	var homePostList []_response.ResponseHomePost
	var count int64
	global.MysqlDB.Model(&db.Post{}).
		Select("p_id,title_photo_url,comment_number,like_number").
		Where("uid = ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&homePostList)

	return count, homePostList
}

//个人主页里面 根据用户UID查询他的收藏帖子列表 每个帖子返回封面、pid
func QueryCollectionListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponsePost) {
	var responsePost []_response.ResponsePost
	var count int64
	global.MysqlDB.Model(db.Collection{}).
		Select("post.p_id,title_photo_url").
		Joins("inner join post on post.p_id = collection.p_id").
		Where("collection.uid= ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&responsePost)
	return count, responsePost
}

//个人主页里面 根据用户UID查询自己点赞的帖子列表 每个帖子返回封面、pid
func QueryLikeList2ByUID(uid uint, page int, pageSize int) (int64, []_response.ResponsePost) {
	var searchUsers []_response.ResponsePost
	var count int64
	global.MysqlDB.Model(db.Liked{}).
		Select("post.p_id,title_photo_url").
		Joins("inner join post on post.p_id = liked.to_like_post_id").
		Where("liked.from_user_id= ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

//个人主页里面 根据用户UID查询发布过帖子列表 每个帖子返回封面、pid
func QueryPostListByUID(uid uint, page int, pageSize int) (int64, []_response.ResponsePost) {
	var postList []_response.ResponsePost
	var count int64
	global.MysqlDB.Model(&db.Post{}).
		Select("p_id,title_photo_url").
		Where("uid=?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&postList)

	return count, postList
}
