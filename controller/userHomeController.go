package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func UserPost(c *gin.Context) {
	var u _request.UserHomeRequest

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	count, totalPage, postList := service.UserPost(u)

	if count == 0 {
		response.Success(c, nil, "没有发过帖子")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "postList": postList}, "ok")
	}

}

func UserFollow(c *gin.Context) {
	var u _request.UserHomeRequest

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	count, totalPage, followList := service.UserFollow(u)

	if count == 0 {
		response.Success(c, nil, "没有关注其他用户")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "followlist": followList}, "ok")
	}

}

func UserFans(c *gin.Context) {
	var u _request.UserHomeRequest

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	count, totalPage, fansList := service.UserFans(u)

	if count == 0 {
		response.Success(c, nil, "没有粉丝")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "fanslist": fansList}, "ok")
	}
}
func UserPostLike(c *gin.Context) {
	var u _request.UserHomeRequest

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	count, totalPage, userPostLikeList := service.UserPostLike(u)

	if count == 0 {
		response.Success(c, nil, "本用户帖子的没有点赞")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "userPostLikeList": userPostLikeList}, "ok")
	}
}

func UserLikePost(c *gin.Context) {
	var u _request.UserHomeRequest

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	count, totalPage, userLikePostList := service.UserLikePost(u)

	if count == 0 {
		response.Success(c, nil, "本用户没有点赞过其他帖子")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "userLikePostList": userLikePostList}, "ok")
	}
}
func UserCollection(c *gin.Context) {
	var u _request.UserHomeRequest

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	count, totalPage, collectionList := service.UserCollection(u)

	if count == 0 {
		response.Success(c, nil, "没有收藏帖子")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "collectionList": collectionList}, "ok")
	}
}
