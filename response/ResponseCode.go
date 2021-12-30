package response

const (
	SuccessCode     = 200
	FailCode        = 300
	CheckFailCode   = 400
	ServerErrorCode = 500
	UnAuthorized    = 600
)

type ResStruct struct {
	HttpStatus int         //http状态
	Code       int         //状态码
	Message    string      //信息
	Data       interface{} //数据
}
