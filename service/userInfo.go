package service

import (
	"net/http"
	"picture_community/entity/_response"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"
)

func GetUserData(uid int) response.ResStruct {
	var userdata _response.UserData

	// 在db.UserData{}数据库中查询，查询结果保存到结构体变量userdata中
	global.MysqlDB.Model(db.UserData{}).
		Select("nickname,sex,profile,motto,followers_number,fans_number,posts_number,collection_number,forward_number").
		Joins("inner join user_detail on user_data.uid = user_detail.uid").
		Where("uid=?", uid).Scan(&userdata)

	// 查找成功，返回用户数据信息
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       userdata,
	}
}

func GetUserPosts(uid int) response.ResStruct {
	var userposts []_response.UserPosts

	// 在db.Post{}数据库中查询，查询结果保存到结构体变量userposts中
	global.MysqlDB.Model(db.Post{}).
		Select("pid,title_photo_url,like_number,comment_number,forward_number").
		Where("uid=?", uid).Scan(userposts)

	// 查找成功，返回用户帖子数组
	return response.ResStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Message:    "ok",
		Data:       userposts,
	}
}
