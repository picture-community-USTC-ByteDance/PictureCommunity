package service

import (
	"fmt"
	"picture_community/dao/post"
)

func CreateLike(u_id uint, post_id uint) bool {

	id, state, err := post.QueryLikeByUserID(u_id, post_id)
	if id == 0 { //没找到时 err仍然等于nil，id=0，state = false
		fmt.Println(err)
		_, err := post.InsertLikeByUserID(u_id, post_id)
		if err != nil {
			return false
		}
		return true
	}
	if state != true {
		_, err := post.UpdateLikeByUserID(id, true)
		if err != nil {
			//fmt.Println(err)
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
