package router

import "picture_community/global"
import "picture_community/controller/login"

func SetRouter() {
	r := global.GinEngin
	g := r.Group("/user")
	{
		g.POST("/login", login.LoginController)
	}
}
