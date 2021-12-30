package router

import (
	"picture_community/controller"
	"picture_community/global"
	"picture_community/middleware"
)

func SetRouter() {
	r := global.GinEngine

	r.GET("/search", controller.Search)
	g := r.Group("/user")
	{
		g.POST("/login", controller.LoginController)
	}
	p := r.Group("/post")
	{
		p.Use(middleware.AuthMiddleware()) //认证是否登录,

		p.POST("/create", controller.CreatePostController)
	}
}
