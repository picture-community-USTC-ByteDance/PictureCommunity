package firstpage

import (
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

/*根据帖子Id获得帖子详细信息*/
func QueryDetailById(pid uint) _response.ResPost {
	var res _response.ResPost
	global.MysqlDB.Table("user_detail a").Debug().
		Select("a.uid,a.nickname,a.`profile`,b.title_photo_url,b.content,b.like_number,b.comment_number,b.create_time").
		Joins("INNER JOIN `post` b ON b.p_id  = ?  and a.uid = b.uid", pid).
		Scan(&res)
	return res
}

/*根据uid获取我所关注的人的所有帖子和转发内容*/
func QueryAllPostByUID(uid uint, page int, pageSize int) []_response.ResPost {
	var result []_response.ResPost
	offset := (page - 1) * pageSize
	global.MysqlDB.Raw("select p1.uid,p1.p_id,u.nickname,u.`profile`,p1.title_photo_url,p1.content,p1.like_number,p1.comment_number,p1.create_time from post p1,user_detail u where p1.uid in (select followed_id from follow where uid = ?) "+
		"and p1.uid = u.uid \nUNION ALL\nselect p2.uid,p2.p_id,u.nickname,u.`profile`,p2.title_photo_url,p2.content,p2.like_number,p2.comment_number,p2.create_time from post p2 ,user_detail u where p_id in (select to_forward_post_id from forward "+
		"where author_user_id\nin (select followed_id from follow where uid = ?)) and p2.uid = u.uid\nORDER BY create_time DESC  limit ? offset ?", uid, uid, pageSize, offset).
		Scan(&result)
	return result
}
func QueryUsernameById(uid uint) string {
	var res db.User
	global.MysqlDB.Select("username").Where("uid = ?", uid).First(&db.User{}).Scan(&res)
	return res.Username
}

/*通过帖子Id获得这个帖子所有的一级评论id*/
func QueryCommentIdById(pid uint) []uint {
	var res []uint
	global.MysqlDB.Model(db.Comment{}).Debug().
		Select("c_id").Where("parent_id = 0 and post_id = ?", pid).Scan(&res)
	return res
}

/*通过评论id获得评论信息(parent_id为0)*/
func QueryCommentById(cid uint) _response.ResComment {
	var res _response.ResComment
	global.MysqlDB.Table("user_detail a").Debug().
		Select("a.nickname,a.`profile`,b.like_number,b.content").
		Joins("INNER JOIN `comment` b ON b.c_id  = ? and b.parent_id = 0 and a.uid = b.user_id", cid).
		Scan(&res)
	return res
}

func QuerySingleDetailById(pid uint) _response.ResSinglePost {
	var res _response.ResSinglePost
	global.MysqlDB.Table("user_detail a").Debug().
		Select("a.uid,a.nickname,a.`profile`,b.title_photo_url,b.content,b.like_number,b.create_time").
		Joins("INNER JOIN `post` b ON b.p_id  = ?  and a.uid = b.uid", pid).
		Scan(&res)
	return res
}

func QueryIsLiked(pid uint, uid uint) bool {
	var res db.Liked
	result := global.MysqlDB.Debug().Select("state").
		Where("from_user_id = ? AND to_like_post_id = ?", uid, pid).
		Find(&db.Liked{}).Scan(&res)
	if result.RowsAffected > 0 && res.State == true {
		return true
	}
	return false
}
func QueryIsFollowed(myUid uint, uid uint) bool {
	var res db.Follow
	result := global.MysqlDB.Debug().Select("state").
		Where("uid = ? AND followed_id = ?", myUid, uid).
		Find(&db.Follow{}).Scan(&res)
	if result.RowsAffected > 0 && res.State == true {
		return true
	}
	return false
}
func QueryAllPhotos(pid uint) []string {
	var res []string
	global.MysqlDB.Model(db.PostPhoto{}).Debug().
		Select("url").Where("p_id = ?", pid).Scan(&res)
	return res
}

func QueryUidByPid(pid uint) uint {
	var res db.Post
	global.MysqlDB.Debug().Select("uid").
		Where("p_id = ?", pid).
		Find(&db.Post{}).Scan(&res)
	return res.UID
}
func QueryPageNum(uid uint) int {
	var res []_response.ResPost
	global.MysqlDB.Raw("select p1.uid from post p1,user_detail u where p1.uid in (select followed_id from follow where uid = ?) "+
		"and p1.uid = u.uid \nUNION ALL\nselect p2.uid from post p2 ,user_detail u where p_id in (select to_forward_post_id from forward "+
		"where author_user_id\nin (select followed_id from follow where uid = ?)) and p2.uid = u.uid", uid, uid).
		Scan(&res)
	return len(res)
}
