package post

import (
	"fmt"
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryCollectionByUserID(u_id uint, post_id uint) (uint, bool, error) {
	var collection db.Collection
	collection.PID = post_id
	collection.UID = u_id
	fmt.Println("uid=", u_id, "post_id=", post_id)
	err := global.MysqlDB.Find(&collection, "p_id=? AND uid=?", post_id, u_id).Error
	fmt.Println("state:", collection.State)
	return collection.ID, collection.State, err
}

func InsertCollectionByUserID(newCollection db.Collection) (int64, error) {

	err := global.MysqlDB.Create(&newCollection).Error

	var post db.Post
	global.MysqlDB.First(&post, newCollection.PID)
	num := post.CollectionNumber
	global.MysqlDB.Model(&post).Where("post_id=?", newCollection.PID).Update("collection_number", num+1) //帖子点赞数➕1

	return int64(newCollection.UID), err

}
func UpdateCollectionByUserID(id uint, state bool) (bool, error) {
	var collection db.Collection
	collection.ID = id
	fmt.Println(id)
	err := global.MysqlDB.Model(&collection).Where("id=?", id).Update("state", state).Error
	post_id := global.MysqlDB.Model(&collection).Select("p_id")
	var post db.Post
	global.MysqlDB.First(&post, post_id)
	num := post.CollectionNumber
	if state {
		global.MysqlDB.Model(&post).Update("collection_number", num+1) //帖子点赞数➕1
		return collection.State, err
	}
	global.MysqlDB.Model(&post).Update("collection_number", num-1) //帖子点赞数➖1
	return collection.State, err
}
