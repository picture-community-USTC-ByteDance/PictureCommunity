package middleware

import (
	"github.com/gin-gonic/gin"
	"picture_community/response"
	"picture_community/utils"
	"strings"
)

/*********************************************************
 函数功能: 用户认证中间件
 用户登录才能用的接口，前端要在请求头带上jwt字符串。
前端请求头格式：Authorization: Bearer <token> 注意berare与token有一个空格。
认证成功后，需要登录的api接口使用 user,_:=ctx.Get(user)获取本次请求的用户信息。

**********************************************************/
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization，header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.UnAuthorization(ctx, nil, "请先登录")
			//抛弃请求
			ctx.Abort()
			return
		}
		//"Bearer "占了7位
		tokenString = tokenString[7:]
		claims, err := utils.ParserToken(tokenString)
		if err != nil {
			response.UnAuthorization(ctx, nil, err.Error())
			//抛弃请求
			ctx.Abort()
			return
		}
		userId := claims.ID
		ctx.Set("uid", userId)
		ctx.Next()
	}
}
