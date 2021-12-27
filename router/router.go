package router

import (
	"picture_community/controller"
	"picture_community/controller/login"
	"picture_community/global"
)

func SetRouter() {
	r := global.GinEngin

	r.GET("/search", controller.Search)
	g := r.Group("/user")
	{
		g.POST("/login", login.LoginController)
		g.POST("/update", controller.UpdateUserInfo)
	}
	p := r.Group("/post")
	{
		p.POST("/create", controller.CreatePostController)
	}
}
