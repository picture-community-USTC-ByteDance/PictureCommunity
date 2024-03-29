package post

import (
	"fmt"
	"gorm.io/gorm"
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

//查询一级评论
func QueryFirstCommentDAO(pagesize int, page int, postid uint) (error, []_response.QueryCommentBack, int64) {
	var coms []_response.QueryCommentBack
	var count int64
	//err := global.MysqlDB.Debug().
	//	Model(&db.Comment{}).
	//	Where("post_id = ? AND parent_id = 0", postid).Count(&count).
	//	Offset((page - 1) * pagesize).Limit(pagesize).Find(&coms).Error
	err := global.MysqlDB.Table("comment").
		Select("comment.c_id,comment.user_id,comment.content,comment.update_time,comment.like_number,comment.child_number,user.username").
		Joins("inner join user on user.uid = comment.user_id").Debug().
		Where("post_id = ? AND parent_id = 0", postid).Count(&count).
		Offset((page - 1) * pagesize).Limit(pagesize).Find(&coms).Error
	fmt.Println(coms)
	fmt.Println(err)
	return err, coms, count
}

//查询二级评论
func QuerySecondCommentDAO(pagesize int, page int, postid uint, parent_id uint) (error, []_response.QueryCommentBack2, int64) {
	var coms []_response.QueryCommentBack2
	var count int64
	err := global.MysqlDB.Table("comment").
		Select("comment.c_id,comment.user_id,comment.content,comment.update_time,comment.like_number,comment.child_number,comment.parent_id,user.username").
		Joins("inner join user on user.uid = comment.user_id").Debug().
		Where("post_id = ? AND parent_id = ?", postid, parent_id).Count(&count).
		Offset((page - 1) * pagesize).Limit(pagesize).Find(&coms).Error
	fmt.Println(coms)
	fmt.Println(err)
	return err, coms, count
}
func QueryNicknameAndProfile(userid uint) (nickname string, profile string) {
	var t struct {
		Nickname string
		Profile  string
	}
	global.MysqlDB.Model(db.UserDetail{}).First(&t, userid)
	return t.Nickname, t.Profile
}

//新添加一个评论，返回添加的这项数据
func CreateFirstLevelCommentDAO(userid uint, postid uint, content string) (error, db.Comment) {
	var com db.Comment
	var re db.Comment

	//先根据帖子id获取帖子，检查帖子是否存在
	var posttemp db.Post
	err := global.MysqlDB.Debug().Where("p_id = ?", postid).First(&posttemp).Error
	if err != nil {
		return err, re
	}

	//插入新的信息到数据库
	//自动生成comment的id
	com.CID = global.CommentIDGenerator.NewID()
	com.Content = content
	com.ParentID = 0
	com.PostID = postid
	com.UserID = userid
	com.ChildNumber = 0
	com.LikeNumber = 0
	com.DeleteStatus = false

	result := global.MysqlDB.Create(&com)
	err = result.Error

	if err == nil {
		//更新帖子的评论数
		err = global.MysqlDB.Model(&posttemp).Update("comment_number", gorm.Expr("comment_number+1")).Error
	}
	global.MysqlDB.First(&re, com.CID)
	return err, re
}

//新添加一个二级评论，返回添加的这项数据
func CreateSecondLevelCommentDAO(userid uint, postid uint, parentid uint, content string) (error, db.Comment) {
	var com db.Comment

	//获取帖子
	var re db.Comment
	var posttemp db.Post
	err := global.MysqlDB.Debug().Where("p_id = ?", postid).First(&posttemp).Error
	if err != nil {
		return err, re
	}
	//获取父评论
	var commenttemp db.Comment
	err = global.MysqlDB.Debug().Where("c_id = ?", parentid).First(&commenttemp).Error
	if err != nil {
		return err, re
	}
	//往数据库里写数据
	//自动生成comment的id
	com.CID = global.CommentIDGenerator.NewID()
	com.Content = content
	com.ParentID = parentid
	com.PostID = postid
	com.UserID = userid
	com.ChildNumber = 0
	com.LikeNumber = 0
	com.DeleteStatus = false

	result := global.MysqlDB.Debug().Create(&com)
	err = result.Error

	if err == nil {
		//更新帖子的评论数
		err = global.MysqlDB.Debug().Model(&posttemp).Update("comment_number", gorm.Expr("comment_number+1")).Error
		if err != nil {
			return err, re
		}
		//更新以及评论的子评论数
		err = global.MysqlDB.Debug().Model(&commenttemp).Update("child_number", gorm.Expr("child_number+1")).Error
	}
	global.MysqlDB.First(&re, com.CID)
	return err, re
}
func DeleteCommentDAO(commentid uint) (error, int) {
	var comment db.Comment
	err := global.MysqlDB.Debug().First(&comment, commentid).Error
	if err != nil {
		return err, -2
	}
	if comment.DeleteStatus == true {
		return err, -1
	}
	comment.Content = "该评论已被删除"
	comment.DeleteStatus = true
	err = global.MysqlDB.Debug().Save(&comment).Error
	return err, 1

}
