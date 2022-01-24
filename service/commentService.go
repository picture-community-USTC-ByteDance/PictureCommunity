package service

import (
	"picture_community/dao/post"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
)

func QueryCommentService(comment _request.QueryComment) (error, []_response.QueryCommentBackTemp) {
	//查询comment表,得到该post下的所有一级评论
	//var backtemp []_response.QueryCommentBackTemp
	err, back := post.QueryFirstCommentDAO(comment.PageSize, comment.Page, comment.PID)
	//根据用户id查询user_detail表，返回用户的昵称、头像
	//back.NickName=" "
	//back.Profile="nil"
	return err, back
}
func CreateFirstLevelCommentService(comment _request.CreateFirstLevelComment) (err error, back _response.CreateFirstLevelCommentBack) {
	err, back = post.CreateFirstLevelCommentDAO(comment.UID, comment.PID, comment.Content)
	return
}
func CreateSecondLevelCommentService(comment _request.CreateSecondLevelComment) (err error, back _response.CreateSecondLevelCommentBack) {
	err, back = post.CreateSecondLevelCommentDAO(comment.UID, comment.PID, comment.ParentId, comment.Content)
	return
}
func DeleteCommentService(comment _request.DeleteComment) (err error) {
	err = post.DeleteCommentDAO(comment.CID)
	return
}
