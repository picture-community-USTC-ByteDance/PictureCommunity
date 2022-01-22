package service

import (
	"math"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
	"strings"
)

func SearchService(u _request.SearchUsers) (int64, int, []_response.ResponseSearchUsers) {
	/*支持多个关键词搜索，关键词用空格隔开*/
	keywords := "%" + strings.Replace(u.NickName, " ", "%", -1) + "%"

	count, searchUsers := user.QueryUserListByNickname(keywords, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, searchUsers
}
