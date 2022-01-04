package main

import (
	"picture_community/entity/db"
	"picture_community/global"
)

/*通过用户id获取关注的人列表*/
func QueryFollowListByUID(uid uint, page int, pageSize int) (int64, []uint) {
	var searchUsers []uint
	var count int64 //FollowedID
	global.MysqlDB.Model(db.Follow{}).
		Select("followed_id").
		Where("follow.uid = ?", uid).Count(&count).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&searchUsers)
	return count, searchUsers
}

/*通过关注的人的id获取他们的最新帖子id*/
func QueryNewPostId(uid uint) uint {
	var pid uint
	result := global.MysqlDB.Model(db.Post{}).
		Select("p_id").
		Where("post.uid = ?", uid).
		Order("update_time desc").Take(&pid)
	if result.RowsAffected == 0 {
		return 0
	} else {
		return pid
	}
}

/*通过关注的人的id获取他们的转发帖子id*/
