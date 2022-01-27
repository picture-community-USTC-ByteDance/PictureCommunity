package service

import (
	"math"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
)

// 查找个人主页的用户信息
func GetUserData(uid uint) (error, _response.UserData) {
	err, res := user.QueryUserDataByUserID(uid)
	return err, res
}

// 查找个人主页的用户帖子数组信息
func GetUserPosts(uid uint, u _request.UserPosts) (int64, int, []_response.UserPosts) {
	count, userPosts := user.QueryUserPostsByUserID(uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	// 查找成功，返回用户帖子数组
	return count, totalPage, userPosts

}
