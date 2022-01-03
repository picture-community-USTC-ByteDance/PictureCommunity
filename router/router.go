package router

import (
	"net/http"
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
		p.POST("/create", middleware.AuthMiddleware, controller.CreatePostController)
	}
	l := r.Group("/like")
	{
		l.POST("/new", middleware.AuthMiddleware, controller.CreateLikeController)
		l.GET("/query", middleware.AuthMiddleware, controller.QueryLikeController)
		l.POST("/cancel", middleware.AuthMiddleware, controller.CancelLikeController)
	}
	r.StaticFS("/upload/pictures", http.Dir("./storage"))

	r.GET("/token", controller.GetToken)

}
