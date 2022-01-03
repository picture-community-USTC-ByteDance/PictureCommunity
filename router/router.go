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
	{
		p.POST("/create", middleware.AuthMiddleware(), controller.CreatePostController)
	}
	r.StaticFS("/upload/pictures", http.Dir("./storage"))
}
