package router

import (
	"net/http"
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
		g.POST("/register", controller.RegisterController)
		g.GET("/queryUsernameIsUnique", controller.UsernameIsUniqueController)
		g.GET("/queryEmailIsUnique", controller.EmailIsUniqueController)
		g.GET("/queryTelephoneIsUnique", controller.TelephoneIsUniqueController)
		g.POST("/updateUserDetail", middleware.AuthMiddleware, controller.UpdateUserController)
	}
	p := r.Group("/post")
	{ //p.Use(middleware.AuthMiddleware())
		p.POST("/create", middleware.AuthMiddleware, controller.CreatePostController)
	}
	f := r.Group("/forward")
	{
		f.POST("/new", middleware.AuthMiddleware, controller.NewForwardController)
	}
	l := r.Group("/like")
	{
		l.POST("/new", middleware.AuthMiddleware, controller.CreateLikeController)
		l.GET("/query", middleware.AuthMiddleware, controller.QueryLikeController)
		l.POST("/cancel", middleware.AuthMiddleware, controller.CancelLikeController)
	}

	u := r.Group("/list")
	{
		u.Use(middleware.AuthMiddleware)
		//u.GET("/:id", controller.UserHome)
		u.GET("/post", controller.UserPost)
		u.GET("/follow", controller.UserFollow)
		u.GET("/fans", controller.UserFans)
		u.GET("/like", controller.UserPostLike)     //给本用户的帖子点赞过的用户列表
		u.GET("/likepost", controller.UserLikePost) //本用户点赞过的帖子
		u.GET("/collection", controller.UserCollection)
	}

	r.StaticFS("/upload/pictures", http.Dir("./storage"))

	r.GET("/token", controller.GetToken)

}
