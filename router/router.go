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
		g.POST("/updateUserDetail", middleware.AuthMiddleware, controller.UpdateUserDetailController)
		g.POST("/updateUserEmail", middleware.AuthMiddleware, controller.UpdateUserEmailController)
		g.POST("/updateUserTelephone", middleware.AuthMiddleware, controller.UpdateUserTelephoneController)
		g.GET("/queryMyDetail", middleware.AuthMiddleware, controller.QueryMyDetailController)
	}
	p := r.Group("/post")
	{ //p.Use(middleware.AuthMiddleware())
		p.POST("/create", middleware.AuthMiddleware, controller.CreatePostController)
		p.POST("/delete", middleware.AuthMiddleware, controller.DeletePostController)
		c := p.Group("/comment")
		{
			c.GET("/query", controller.QueryCommentController)
			c.GET("/querySecond", controller.QueryCommentController2)
			c.POST("/new", controller.AddFirstLevelCommentController)
			c.DELETE("/delete", controller.DeleteCommentController)
			c.POST("/secondNew", controller.AddSecondLevelCommentController)
		}

	}
	q := r.Group("/query")
	{
		q.GET("/userData", middleware.AuthMiddleware, controller.QueryUserData)
		q.GET("/userDetail", middleware.AuthMiddleware, controller.QueryUserPosts)
	}
	f := r.Group("/forward")
	{
		f.POST("/new", middleware.AuthMiddleware, controller.NewForwardController)
		f.POST("/delete", middleware.AuthMiddleware, controller.DeleteForwardController)
	}
	l := r.Group("/like")
	{
		l.POST("/new", middleware.AuthMiddleware, controller.CreateLikeController)
		l.GET("/query", middleware.AuthMiddleware, controller.QueryLikeController)
		l.POST("/cancel", middleware.AuthMiddleware, controller.CancelLikeController)
	}
	c := r.Group("/collection")
	{
		c.POST("/new", middleware.AuthMiddleware, controller.CreateCollectionController)
		c.GET("/query", middleware.AuthMiddleware, controller.QueryCollectionController)
		c.POST("/cancel", middleware.AuthMiddleware, controller.CancelCollectionController)
	}
	r.GET("/list/post", controller.UserPost)
	r.GET("/list/collection", controller.UserCollection)
	r.GET("/list/likepost", controller.UserLikePost) //本用户点赞过的帖子
	u := r.Group("/list")
	{
		u.Use(middleware.AuthMiddleware)
		//u.GET("/:id", controller.UserHome)
		u.GET("/follow", controller.UserFollow)
		u.GET("/fans", controller.UserFans)
		u.GET("/like", controller.UserPostLike) //给本用户的帖子点赞过的用户列表
	}
	m := r.Group("/firstpage")
	{

		m.GET("/getIdList", controller.GetIdListController)
		m.GET("/getDetailList", controller.GetDetailController)
	}

	r.POST("/upload", middleware.AuthMiddleware, controller.FileUploadController)
	r.StaticFS("/upload/pictures", http.Dir("./storage"))

	r.GET("/token", controller.GetToken)

}
