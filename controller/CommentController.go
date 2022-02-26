package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func AddFirstLevelCommentController(c *gin.Context) {
	var u _request.CreateFirstLevelComment
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "参数出错")
		return
	}
	uid, _ := c.Get("uid")
	err, res := service.CreateFirstLevelCommentService(u, uid.(uint))
	if err != nil {
		response.Fail(c, nil, "插入失败")
	} else {
		response.Success(c, res, "一级评论插入成功")
	}

}

func AddSecondLevelCommentController(c *gin.Context) {
	var u _request.CreateSecondLevelComment
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "参数出错")
		return
	}
	uid, _ := c.Get("uid")
	err, res := service.CreateSecondLevelCommentService(u, uid.(uint))
	if err != nil {
		response.Fail(c, nil, "插入失败")
	} else {
		response.Success(c, res, "二级评论插入成功")
	}
}

//查询一级评论
func QueryCommentController(c *gin.Context) {
	var u _request.QueryComment
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "参数错误")
		return
	}
	if u.Page < 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码错误")
		return
	}
	err, res, totalpage := service.QueryCommentService(u)
	if err != nil {
		response.Fail(c, nil, "评论返回失败")
	} else {
		response.Success(c, gin.H{
			"comments":  res,
			"totalpage": totalpage,
		}, "返回一级评论成功")
	}
}

//查询二级评论
func QueryCommentController2(c *gin.Context) {
	var u _request.QueryComment2
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "参数错误")
		return
	}
	if u.Page < 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码错误")
		return
	}
	err, res, totalpage := service.QueryCommentService2(u)
	if err != nil {
		response.Fail(c, nil, "评论返回失败")
	} else {
		response.Success(c, gin.H{
			"comments":  res,
			"totalpage": totalpage,
		}, "返回二级评论成功")

	}
}

func DeleteCommentController(c *gin.Context) {
	var u _request.DeleteComment
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "参数错误")
		return
	}
	err, errtype := service.DeleteCommentService(u)
	if err != nil {
		response.Fail(c, nil, "评论删除出错")
	} else {
		if errtype == -1 {
			response.Fail(c, nil, "该评论已不存在")
		} else if errtype == 1 {
			response.Success(c, nil, "删除评论成功")
		}
	}
}
