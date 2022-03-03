package service

import (
	"math"
	"picture_community/dao/post"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
)

func QueryCommentService(comment _request.QueryComment) (error, []_response.QueryCommentBack, int) {
	//查询comment表,得到该post下的所有一级评论
	err, back, count := post.QueryFirstCommentDAO(comment.PageSize, comment.Page, comment.PID)
	if err != nil {
		return err, back, -1
	}
	for i, com := range back {
		back[i].NickName, back[i].Profile = post.QueryNicknameAndProfile(com.UserID)
	}
	totalpage := int(math.Ceil(float64(count) / float64(comment.PageSize)))
	return err, back, totalpage
}
func QueryCommentService2(comment _request.QueryComment2) (error, []_response.QueryCommentBack2, int) {
	//查询comment表,得到该post和该一级评论下的所有二级评论
	err, back, count := post.QuerySecondCommentDAO(comment.PageSize, comment.Page, comment.PID, comment.ParentId)
	if err != nil {
		return err, back, -1
	}
	for i, com := range back {
		back[i].NickName, back[i].Profile = post.QueryNicknameAndProfile(com.UserID)
	}
	totalpage := int(math.Ceil(float64(count) / float64(comment.PageSize)))
	return err, back, totalpage
}
func CreateFirstLevelCommentService(comment _request.CreateFirstLevelComment, uid uint) (error, _response.CreateCommentBack) {
	err, backtemp := post.CreateFirstLevelCommentDAO(uid, comment.PID, comment.Content)
	if err != nil {
		return err, _response.CreateCommentBack{}
	}
	back := _response.CreateCommentBack{
		CID:          backtemp.CID,
		PostID:       backtemp.PostID,
		ParentID:     backtemp.ParentID,
		ChildNumber:  backtemp.ChildNumber,
		Content:      backtemp.Content,
		UserID:       backtemp.UserID,
		LikeNumber:   backtemp.LikeNumber,
		DeleteStatus: backtemp.DeleteStatus,
		UpdateTime:   backtemp.UpdateTime,
		CreateTime:   backtemp.CreateTime,
	}
	back.NickName, back.Profile = post.QueryNicknameAndProfile(backtemp.UserID)
	return err, back
}
func CreateSecondLevelCommentService(comment _request.CreateSecondLevelComment, uid uint) (error, _response.CreateCommentBack) {
	err, backtemp := post.CreateSecondLevelCommentDAO(uid, comment.PID, comment.ParentId, comment.Content)
	if err != nil {
		return err, _response.CreateCommentBack{}
	}
	back := _response.CreateCommentBack{
		CID:          backtemp.CID,
		PostID:       backtemp.PostID,
		ParentID:     backtemp.ParentID,
		ChildNumber:  backtemp.ChildNumber,
		Content:      backtemp.Content,
		UserID:       backtemp.UserID,
		LikeNumber:   backtemp.LikeNumber,
		DeleteStatus: backtemp.DeleteStatus,
		UpdateTime:   backtemp.UpdateTime,
		CreateTime:   backtemp.CreateTime,
	}
	back.NickName, back.Profile = post.QueryNicknameAndProfile(backtemp.UserID)
	return err, back
}
func DeleteCommentService(comment _request.DeleteComment) (err error, err_type int) {
	err, err_type = post.DeleteCommentDAO(comment.CID)
	return
}
