package middleware

import (
	"picture_community/response"
	"picture_community/utils"

	"github.com/gin-gonic/gin"
)

/*********************************************************
 函数功能: 用户认证中间件
 用户登录才能用的接口，前端要在请求头带上jwt字符串。
前端请求头格式：Authorization: Bearer <token> 注意berare与token有一个空格。
认证成功后，需要登录的api接口使用 user,_:=ctx.Get(user)获取本次请求的用户信息。

**********************************************************/
func AuthMiddleware(ctx *gin.Context) {
	//获取Authorization，header
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		response.UnAuthorization(ctx, nil, "请先登录")
		//抛弃请求
		ctx.Abort()
		return
	}
	claims, err := utils.ParserToken(tokenString)
	if err != nil {
		response.UnAuthorization(ctx, nil, err.Error())
		ctx.Abort()
		return
	}
	userId := claims.ID
	ctx.Set("uid", uint(userId))
	ctx.Next()
}
