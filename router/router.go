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
		g.POST("/register", controller.RegisterController)
		g.GET("/queryUsernameIsUnique", controller.UsernameIsUniqueController)
		g.GET("/queryEmailIsUnique", controller.EmailIsUniqueController)
		g.GET("/queryTelephoneIsUnique", controller.TelephoneIsUniqueController)
	}
	p := r.Group("/post")
	{ //p.Use(middleware.AuthMiddleware())
		p.POST("/create", middleware.AuthMiddleware(), controller.CreatePostController)
	}
}
