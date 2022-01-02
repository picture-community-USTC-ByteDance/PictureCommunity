package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			//抛弃请求
			ctx.Abort()
			return
		}
		//"Bearer "占了7位
		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			//抛弃请求
			ctx.Abort()
			return
		}

		//验证通过后，token获取userID
		userId := claims.ID
		//var user db.User
		//global.MysqlDB.First(&user, userId)

		//没找到用户
		if userId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			//抛弃请求
			ctx.Abort()
			return
		}

		//如果用户存在，将用户信息写入上下文,其他需要登录才能使用的api可用user,_:=ctx.Get("uid")
		ctx.Set("uid", userId)
		//ctx.Set("user", user)
		ctx.Next()
	}
}
