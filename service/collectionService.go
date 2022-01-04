package service

import (
	"picture_community/dao/post"
	"picture_community/entity/db"
)

func CreateCollection(u_id uint, post_id uint) bool {

	newCollection := db.Collection{
		UID:   u_id,
		PID:   post_id,
		State: true,
	}

	id, state, _ := post.QueryCollectionByUserID(u_id, post_id)
	if id == 0 {
		_, err := post.InsertCollectionByUserID(newCollection)
		if err != nil {
			return false
		}
		return true
	}
	if state != true {
		_, err := post.UpdateCollectionByUserID(id, true)
		if err != nil {
			return false
		}
		return true
	}
	return true
}

func QueryCollection(u_id uint, post_id uint) bool {
	_, state, err := post.QueryCollectionByUserID(u_id, post_id)
	if err != nil {
		return false
	}
	return state
}
func CancelCollection(u_id uint, post_id uint) bool {
	id, state, _ := post.QueryCollectionByUserID(u_id, post_id)
	if state == true {
		_, err := post.UpdateCollectionByUserID(id, false)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
