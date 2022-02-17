package post

import (
	"fmt"
	"gorm.io/gorm"
	"picture_community/entity/db"
	"picture_community/global"
	"time"
)

func QueryLikeByUserID(u_id uint, post_id uint) (uint, bool, error) {
	var like db.Liked
	like.ToLikePostID = post_id
	like.FromUserID = u_id
	fmt.Println("uid=", u_id, "post_id=", post_id)
	//err := global.MysqlDB.First(&like).Error
	err := global.MysqlDB.Find(&like, "to_like_post_id=? AND from_user_id=?", post_id, u_id).Error
	fmt.Println("err=", err)
	//查询是否存在这个帖子
	var post db.Post
	err_p := global.MysqlDB.First(&post, post_id).Error
	if err_p == gorm.ErrRecordNotFound { //不存在该帖子
		//fmt.Println("err_p:", err_p)
		return like.ID, like.State, err_p
	}
	return like.ID, like.State, err
}

func InsertLikeByUserID(u_id uint, post_id uint) (int64, error) {
	newLike := db.Liked{
		ToLikePostID: post_id,
		FromUserID:   u_id,
		State:        true,
		UpdateTime:   time.Time{},
		CreateTime:   time.Time{},
	}
	err := global.MysqlDB.Create(&newLike).Error
	//result := global.MysqlDB.Create(&newLike)
	//fmt.Println(result.RowsAffected)
	var post db.Post
	global.MysqlDB.First(&post, post_id)
	num := post.LikeNumber
	global.MysqlDB.Model(&post).Update("like_number", num+1) //帖子点赞数➕1

	return int64(newLike.FromUserID), err

}

func UpdateLikeByUserID(id uint, state bool) (bool, error) {
	var like db.Liked
	like.ID = id
	err := global.MysqlDB.Model(&like).Where("id=?", id).Update("state", state).Error
	global.MysqlDB.First(&like)
	post_id := like.ToLikePostID

	var post db.Post
	global.MysqlDB.First(&post, post_id)
	num := post.LikeNumber
	if state {
		global.MysqlDB.Model(&post).Update("like_number", num+1) //帖子点赞数➕1
		return like.State, err
	}
	if num != 0 {
		global.MysqlDB.Model(&post).Update("like_number", num-1) //帖子点赞数➕1
	}

	return like.State, err
}
