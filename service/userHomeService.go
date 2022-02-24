package service

import (
	"math"
	"picture_community/dao/post"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
)

func UserPost(u _request.UserHomeRequest) (int64, int, []_response.ResponseHomePost) {

	count, followList := post.QueryHomePostListByUID(u.Uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, followList
}

func UserFollow(u _request.UserHomeRequest) (int64, int, []_response.ResponseSearchUsers) {

	count, followList := user.QueryFollowListByUID(u.Uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, followList
}

func UserFans(u _request.UserHomeRequest) (int64, int, []_response.ResponseSearchUsers) {

	count, fansList := user.QueryFansListByUID(u.Uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, fansList
}

func UserPostLike(u _request.UserHomeRequest) (int64, int, []_response.ResponseSearchUsers) {

	count, likeList := user.QueryLikeList1ByUID(u.Uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, likeList
}

func UserLikePost(u _request.UserHomeRequest) (int64, int, []_response.ResponseHomePost) {

	count, likePostList := post.QueryLikePostByUID(u.Uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, likePostList
}

func UserCollection(u _request.UserHomeRequest) (int64, int, []_response.ResponseHomePost) {

	count, collectionList := post.QueryHomeCollectionListByUID(u.Uid, u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, collectionList
}
