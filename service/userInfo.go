package service

import (
	"fmt"
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

// 根据username查其他用户的信息
func GetUserDataByUsername(username string, uid uint) (error, _response.UserData, bool, bool) {
	// 查指定用户的个人信息
	err, res := user.QueryUserDataByUsername(username)

	// 查指定用户与当前用户的关注与被关注信息
	countFollow, countFan := user.QueryIsFollowAndFan(username, uid)
	var isFollow, isFan bool
	if countFollow > 0 { // 查找成功，当前用户 的关注者包括 指定用户
		isFollow = true
		fmt.Println("isFollow = true")
	} else {
		isFollow = false
		fmt.Println("isFollow = false")
	}
	if countFan > 0 { // 查找成功，当前用户 的粉丝包括 指定用户
		isFan = true
	} else {
		isFan = false
	}
	return err, res, isFollow, isFan
}

// 查找个人主页的用户帖子数组信息
func GetUserPosts(uid uint, u _request.UserPosts) (int64, int, []_response.UserPosts) {
	count, userPosts := user.QueryUserPostsByUserID(uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	// 查找成功，返回用户帖子数组
	return count, totalPage, userPosts

}
