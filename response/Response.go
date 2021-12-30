package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, SuccessCode, data, msg)
}

func Fail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusBadRequest, FailCode, data, msg)
}

func CheckFail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusUnprocessableEntity, CheckFailCode, data, msg)
}

func ServerError(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusInternalServerError, ServerErrorCode, data, msg)
}

func UnAuthorization(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusInternalServerError, UnAuthorized, data, msg)
}

/*********************************************************
** 函数功能: 处理返回信息
**********************************************************/
func HandleResponse(ctx *gin.Context, res ResStruct) {
	ctx.JSON(res.HttpStatus, gin.H{"code": res.Code, "data": res.Data, "message": res.Message})
}
