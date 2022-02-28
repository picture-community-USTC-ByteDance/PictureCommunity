package service

import (
	"math"
	"picture_community/dao/post"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
)

func QueryCommentService(comment _request.QueryComment) (error, []_response.QueryCommentBackTemp, int) {
	//查询comment表,得到该post下的所有一级评论
	err, back, count := post.QueryFirstCommentDAO(comment.PageSize, comment.Page, comment.PID)
	totalpage := int(math.Ceil(float64(count) / float64(comment.PageSize)))
	return err, back, totalpage
}
func QueryCommentService2(comment _request.QueryComment2) (error, []_response.QueryCommentBackTemp2, int) {
	//查询comment表,得到该post和该一级评论下的所有二级评论
	err, back, count := post.QuerySecondCommentDAO(comment.PageSize, comment.Page, comment.PID, comment.ParentId)
	totalpage := int(math.Ceil(float64(count) / float64(comment.PageSize)))
	return err, back, totalpage
}
func CreateFirstLevelCommentService(comment _request.CreateFirstLevelComment, uid uint) (err error, back _response.CreateFirstLevelCommentBack) {
	err, back = post.CreateFirstLevelCommentDAO(uid, comment.PID, comment.Content)
	return
}
func CreateSecondLevelCommentService(comment _request.CreateSecondLevelComment, uid uint) (err error, back _response.CreateSecondLevelCommentBack) {
	err, back = post.CreateSecondLevelCommentDAO(uid, comment.PID, comment.ParentId, comment.Content)
	return
}
func DeleteCommentService(comment _request.DeleteComment) (err error, err_type int) {
	err, err_type = post.DeleteCommentDAO(comment.CID)
	return
}
