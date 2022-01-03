package service

import (
	"picture_community/dao/post"
	"picture_community/entity/db"
	"time"
)

func CreateLike(u_id uint, post_id uint) bool {
	newLike := db.Liked{
		ToLikePostID: u_id,
		FromUserID:   post_id,
		State:        true,
		UpdateTime:   time.Time{},
		CreateTime:   time.Time{},
	}
	id, state, err := post.QueryLikeByUserID(u_id, post_id)
	if err != nil {
		_, err := post.InsertLikeByUserID(newLike)
		if err != nil {
			return false
		}
		return true
	}
	if state != true {
		_, err := post.UpdateLikeByUserID(id, true)
		if err != nil {
			return false
		}
		return true
	}
	return true
}

func QueryLike(u_id uint, post_id uint) bool {
	_, state, err := post.QueryLikeByUserID(u_id, post_id)
	if err != nil {
		return false
	}
	return state
}
func CancelLike(u_id uint, post_id uint) bool {
	id, state, _ := post.QueryLikeByUserID(u_id, post_id)
	if state == true {
		_, err := post.UpdateLikeByUserID(id, false)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
