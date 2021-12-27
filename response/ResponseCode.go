package response

import "github.com/gin-gonic/gin"

const (
	SuccessCode     = 2000
	FailCode        = 3000
	CheckFailCode   = 4000
	ServerErrorCode = 5000
)

type ResStruct struct {
	HttpStatus int    //http状态
	Code       int    //状态码
	Message    string //信息
	Data       gin.H  //数据

}
