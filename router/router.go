package router

import (
	"picture_community/controller/login"
	"picture_community/controller/post"
	"picture_community/global"
)

func SetRouter() {
	r := global.GinEngin
	g := r.Group("/user")
	{
		g.POST("/login", login.LoginController)
	}
	p := r.Group("/post")
	{
		p.POST("/create", post.CreatePostController)
	}
}
