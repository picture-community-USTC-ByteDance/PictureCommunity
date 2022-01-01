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
	{ //p.Use(middleware.AuthMiddleware())
		p.POST("/create", middleware.AuthMiddleware(), controller.CreatePostController)
	}
	l := r.Group("/like")
	{
		l.POST("/new", controller.CreateLikeController)
		l.GET("/query", controller.QueryLikeController)
		l.POST("/cancel", controller.CancelLikeController)
	}
}
