package router

import (
	"picture_community/controller"
	"picture_community/global"
	"picture_community/middleware"
)

func SetRouter() {
	r := global.GinEngine

	r.GET("/search", controller.SearchUsers)
	g := r.Group("/user")
	{
		g.POST("/login", controller.LoginController)
	}
	p := r.Group("/post")
	{ //p.Use(middleware.AuthMiddleware())
		p.POST("/create", middleware.AuthMiddleware(), controller.CreatePostController)
	}
	u := r.Group("/list")
	{ //u.Use(middleware.AuthMiddleware())
		//u.GET("/:id", controller.UserHome)
		u.GET("/post", controller.UserPost)             //本用户发布的帖子
		u.GET("/follow", controller.UserFollow)         //本用户关注用户列表
		u.GET("/fans", controller.UserFans)             //本用户粉丝列表
		u.GET("/like", controller.UserPostLike)         //给本用户的帖子点赞过的用户列表
		u.GET("/likepost", controller.UserLikePost)     //本用户点赞过的帖子列表
		u.GET("/collection", controller.UserCollection) //本用户收藏帖子列表
	}
}
