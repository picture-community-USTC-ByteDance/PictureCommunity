package post

import (
	"fmt"
	"gorm.io/gorm"
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)


//查询一级评论
func QueryFirstCommentDAO(pagesize int, page int, postid uint) (error, []_response.QueryCommentBackTemp, int64) {
	var coms []_response.QueryCommentBackTemp
	var count int64
	err := global.MysqlDB.Debug().
		Model(&db.Comment{}).
		Where("post_id = ? AND parent_id = 0", postid).Count(&count).
		Offset((page - 1) * pagesize).Limit(pagesize).Find(&coms).Error
	fmt.Println(coms)
	return err, coms, count
}


//查询二级评论
func QuerySecondCommentDAO(pagesize int, page int, postid uint, parent_id uint) (error, []_response.QueryCommentBackTemp2, int64) {

	var coms []_response.QueryCommentBackTemp2
	var count int64
	err := global.MysqlDB.Debug().
		Model(&db.Comment{}).
		Where("post_id = ? AND parent_id = ?", postid, parent_id).Count(&count).
		Offset((page - 1) * pagesize).Limit(pagesize).Find(&coms).Error
	fmt.Println(coms)
	return err, coms, count
}

func CreateFirstLevelCommentDAO(userid uint, postid uint, content string) (error, _response.CreateFirstLevelCommentBack) {
	var tmp db.UserDetail
	var com db.Comment
	//根据id获取该用户的一些其他信息：昵称、头像
	err := global.MysqlDB.First(&tmp, userid).Error
	if err != nil {
		return err, com
	}

	var re _response.CreateFirstLevelCommentBack
	//先根据帖子id获取帖子，检查帖子是否存在
	var posttemp db.Post
	err = global.MysqlDB.Debug().Where("p_id = ?", postid).First(&posttemp).Error
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
	com.Profile = tmp.Profile
	com.NickName = tmp.Nickname

	result := global.MysqlDB.Create(&com)
	err = result.Error

	if err == nil {
		//更新帖子的评论数
		err = global.MysqlDB.Model(&posttemp).Update("comment_number", gorm.Expr("comment_number+1")).Error
	}
	global.MysqlDB.First(&re, com.CID)
	return err, re
}

func CreateSecondLevelCommentDAO(userid uint, postid uint, parentid uint, content string) (error, _response.CreateSecondLevelCommentBack) {
	var tmp db.UserDetail
	var com db.Comment
	//根据id获取该用户的一些其他信息：昵称、头像
	err := global.MysqlDB.First(&tmp, userid).Error
	if err != nil {
		return err, com
	}
	//获取帖子
	var re _response.CreateSecondLevelCommentBack
	var posttemp db.Post
	err = global.MysqlDB.Debug().Where("p_id = ?", postid).First(&posttemp).Error
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
	com.Profile = tmp.Profile
	com.NickName = tmp.Nickname

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
