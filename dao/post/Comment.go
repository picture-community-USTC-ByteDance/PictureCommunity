package post

import (
	"fmt"
	"gorm.io/gorm"
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryFirstCommentDAO(pagesize int, page int, postid int) (error, []_response.QueryCommentBackTemp) {
	var coms []_response.QueryCommentBackTemp
	err := global.MysqlDB.Debug().Model(&db.Comment{}).Where("post_id = ? AND parent_id = 0", postid).Offset(page - 1).Limit(pagesize).Find(&coms).Error
	fmt.Println(coms)
	return err, coms
}
func QuerySecondCommentDAO(pagesize int, page int, postid int, parent_id int) (error, []_response.QueryCommentBackTemp2) {
	var coms []_response.QueryCommentBackTemp2
	err := global.MysqlDB.Debug().Model(&db.Comment{}).Where("post_id = ? AND parent_id = ?", postid, parent_id).Offset(page - 1).Limit(pagesize).Find(&coms).Error
	fmt.Println(coms)
	return err, coms
}
func CreateFirstLevelCommentDAO(userid uint, postid uint, content string) (error, _response.CreateFirstLevelCommentBack) {
	var com db.Comment
	//自动生成comment的id
	//com.CID = 100
	com.Content = content
	com.ParentID = 0
	com.PostID = postid
	com.UserID = userid
	com.ChildNumber = 0
	com.LikeNumber = 0
	com.DeleteStatus = false
	result := global.MysqlDB.Debug().Create(&com)
	err := result.Error
	var re _response.CreateFirstLevelCommentBack
	re.PID = postid
	re.Content = content
	if err == nil {
		//更新帖子的评论数
		var posttemp db.Post
		err = global.MysqlDB.Debug().Where("p_id = ?", postid).First(&posttemp).Error
		if err != nil {
			return err, re
		}
		err = global.MysqlDB.Debug().Model(&posttemp).Update("comment_number", gorm.Expr("comment_number+1")).Error
	}
	return err, re
}
func CreateSecondLevelCommentDAO(userid uint, postid uint, parentid uint, content string) (error, _response.CreateSecondLevelCommentBack) {
	var com db.Comment
	//自动生成comment的id
	//com.CID=0
	com.Content = content
	com.ParentID = parentid
	com.PostID = postid
	com.UserID = userid
	com.ChildNumber = 0
	com.LikeNumber = 0
	com.DeleteStatus = false
	result := global.MysqlDB.Debug().Create(&com)
	err := result.Error
	var re _response.CreateSecondLevelCommentBack
	re.PID = postid
	re.ParentID = parentid
	re.Content = content
	if err == nil {
		//更新帖子的评论数
		var posttemp db.Post
		err = global.MysqlDB.Debug().Where("p_id = ?", postid).First(&posttemp).Error
		if err != nil {
			return err, re
		}
		err = global.MysqlDB.Debug().Model(&posttemp).Update("comment_number", gorm.Expr("comment_number+1")).Error
		if err != nil {
			return err, re
		}
		//更新以及评论的子评论数
		var commenttemp db.Comment
		err = global.MysqlDB.Debug().Where("c_id = ?", parentid).First(&commenttemp).Error
		if err != nil {
			return err, re
		}
		err = global.MysqlDB.Debug().Model(&commenttemp).Update("child_number", gorm.Expr("child_number+1")).Error
	}
	return err, re
}
func DeleteCommentDAO(commentid uint) error {
	var comment db.Comment
	err := global.MysqlDB.Debug().Where("c_id = ?", commentid).First(&comment).Error
	if err != nil {
		return err
	}
	err = global.MysqlDB.Debug().Model(&comment).Update("content", "该评论已被删除").Error
	return err
}
