package firstpage

import (
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"
)

/*根据帖子Id获得帖子详细信息*/
func QueryDetailById(pid uint) response.ResPost {
	var res response.ResPost
	global.MysqlDB.Table("user_detail a").Debug().
		Select("a.uid,a.nickname,a.`profile`,b.title_photo_url,b.content,b.like_number,b.comment_number,b.create_time").
		Joins("INNER JOIN `post` b ON b.p_id  = ?  and a.uid = b.uid", pid).
		Scan(&res)
	return res
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
func QueryCommentById(cid uint) response.ResComment {
	var res response.ResComment
	global.MysqlDB.Table("user_detail a").Debug().
		Select("a.nickname,a.`profile`,b.like_number,b.content").
		Joins("INNER JOIN `comment` b ON b.c_id  = ? and b.parent_id = 0 and a.uid = b.user_id", cid).
		Scan(&res)
	return res
}

func QuerySingleDetailById(pid uint) response.ResSinglePost {
	var res response.ResSinglePost
	global.MysqlDB.Table("user_detail a").Debug().
		Select("a.nickname,a.`profile`,b.title_photo_url,b.content,b.like_number,b.create_time").
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
