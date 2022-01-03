package post

import (
	"fmt"
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryLikeByUserID(u_id uint, post_id uint) (uint, bool, error) {
	var like db.Liked
	like.ToLikePostID = post_id
	like.FromUserID = u_id
	fmt.Println(u_id, post_id)
	err := global.MysqlDB.First(&like).Error
	return like.ID, like.State, err
}

func InsertLikeByUserID(newLike db.Liked) (int64, error) {
	fmt.Println(newLike.ToLikePostID)
	err := global.MysqlDB.Create(&newLike).Error
	var post db.Post
	global.MysqlDB.First(&post, newLike.ToLikePostID)
	num := post.LikeNumber
	global.MysqlDB.Model(&post).Update("LikeNumber", num+1) //帖子点赞数➕1
	return int64(newLike.FromUserID), err

}
func UpdateLikeByUserID(id uint, state bool) (bool, error) {
	var like db.Liked
	like.ID = id
	fmt.Println(id)
	err := global.MysqlDB.Model(&like).Update("state", state).Error
	post_id := global.MysqlDB.Model(&like).Select("ToLikePostID")
	var post db.Post
	global.MysqlDB.First(&post, post_id)
	num := post.LikeNumber
	if state {
		global.MysqlDB.Model(&post).Update("LikeNumber", num+1) //帖子点赞数➕1
	}
	global.MysqlDB.Model(&post).Update("LikeNumber", num-1) //帖子点赞数➖1
	return like.State, err
}
