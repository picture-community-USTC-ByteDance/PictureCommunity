package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func CreateCollectionController(c *gin.Context) {
	//var u db.Collection
	var u _request.CollectionRequest
	err := c.ShouldBind(&u)
	uid, _ := c.Get("uid")
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.CreateCollection(uid.(uint), u.PostID)
	if res {
		response.Success(c, nil, "ok")
		return
	}
	response.Fail(c, nil, "creat error")
}
func QueryCollectionController(c *gin.Context) {
	var u _request.CollectionRequest
	err := c.ShouldBind(&u)
	uid, _ := c.Get("uid")
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.QueryCollection(uid.(uint), u.PostID)
	//res := service.QueryCollection(u.UID, u.PostID)
	response.Success(c, res, "ok")
}

func CancelCollectionController(c *gin.Context) {
	//var u db.Collectiond
	var u _request.CollectionRequest
	err := c.ShouldBind(&u)
	uid, _ := c.Get("uid")
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.CancelCollection(uid.(uint), u.PostID)
	if res {
		response.Success(c, nil, "ok")
		return
	}
	response.Fail(c, nil, "cancel collection error")
}
