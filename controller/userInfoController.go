package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
	"reflect"
	"strings"
	"time"
)

const TIME_LAYOUT = "2006-01-02"

var loc *time.Location

func init() {
	var err error
	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
}

func strFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		vv[0] = vv[0] - 'a' + 'A'
		upperStr += string(vv)
	}
	return upperStr
}

func UpdateUserDetailController(c *gin.Context) {
	var u _request.UpdateUserDetailInfo
	var birthdayTime time.Time
	updateData := make(map[string]interface{})
	if err := c.BindJSON(&updateData); err != nil {
		fmt.Println("bind error")
		fmt.Println(err)
		response.CheckFail(c, nil, "更新用户信息参数错误")
		return
	}
	fmt.Printf("%v \n", &updateData)
	checkTemp := reflect.TypeOf(u)
	for k := range updateData {
		fieldName := strFirstToUpper(k)
		fmt.Println(fieldName)
		if _, ok := checkTemp.FieldByName(fieldName); !ok {
			response.CheckFail(c, nil, "更新用户信息参数错误")
			return
		}
	}
	if birVal, ok := updateData["birthday"]; ok {
		var err error
		birthdayTime, err = time.ParseInLocation(TIME_LAYOUT, birVal.(string), loc)
		if err != nil {
			fmt.Println(err)
			response.CheckFail(c, nil, "更新用户信息参数错误")
			return
		}
		updateData["birthday"] = birthdayTime
	}
	uid, _ := c.Get("uid")
	isOK, message := service.UpdateUserDetailService(updateData, uid.(uint))
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
}

func EmailIsUniqueController(c *gin.Context) {
	var u _request.EmailIsUniqueInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "检查email参数错误")
		return
	}

	isOK, message := service.EmailIsUniqueService(u)
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}

func TelephoneIsUniqueController(c *gin.Context) {
	var u _request.TelephoneIsUniqueInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "检查电话号参数错误")
		return
	}

	isOK, message := service.TelephoneIsUniqueService(u)
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
}

func UpdateUserEmailController(c *gin.Context) {
	var u _request.UpdateUserEmailInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "更新email参数错误")
		return
	}
	uid, _ := c.Get("uid")
	isOK, message := service.UpdateUserEmailService(u, uid.(uint))
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}

func UpdateUserTelephoneController(c *gin.Context) {
	var u _request.UpdateUserTelephoneInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "更新电话号码参数错误")
		return
	}

	uid, _ := c.Get("uid")
	isOK, message := service.UpdateUserTelephoneService(u, uid.(uint))
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}

func QueryMyDetailController(c *gin.Context) {
	uid, _ := c.Get("uid")
	isOK, message, detail := service.QueryMyDetailService(uid.(uint))
	if isOK {
		response.Success(c, detail, message)
	} else {
		response.Fail(c, nil, message)
	}
}
