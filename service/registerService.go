package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/db"
	"picture_community/response"
	"time"
)

func RegisterService(param _request.RegisterUser) response.ResStruct {
	//todo 生成ID
	UID, _ := user.QueryIDAndPasswordByUsername(param.Username)
	if UID > 0 {
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    "Register fail: duplicate username",
			Data:       nil,
		}
	}
	newUser := db.User{
		UID:       produceID(),
		Username:  param.Username,
		Password:  param.Password,
		Telephone: 0,
		Email:     "",
		Status:    0,
	}
	err := user.InsertUser(newUser)
	if err != nil {
		fmt.Println(err)
		return response.ResStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Message:    "Register fail",
			Data:       gin.H{"err": err},
		}
	} else {
		return response.ResStruct{
			HttpStatus: http.StatusOK,
			Code:       response.SuccessCode,
			Message:    "Register success",
			Data:       nil,
		}
	}
}

func IsUnique(param _request.IsUniqueInfo) response.ResStruct {
	//todo 接口修改
	ID, _ := user.QueryIDAndPasswordByUsername(param.Username)

	if ID > 0 {
		return response.ResStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Message:    "duplicate username",
			Data:       nil,
		}
	} else {
		return response.ResStruct{
			HttpStatus: http.StatusOK,
			Code:       response.SuccessCode,
			Message:    "unique username",
			Data:       nil,
		}
	}
}

func produceID() uint {
	rand.Seed(time.Now().Unix())
	data := rand.Int63n(100)
	return uint(data)
}
