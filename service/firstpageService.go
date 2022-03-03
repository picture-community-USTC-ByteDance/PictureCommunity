package service

import (
	"github.com/gin-gonic/gin"
	"picture_community/dao/firstpage"
	"picture_community/entity/_response"
)

/*获得关注的人对应的帖子Id和转发Id*/
func GetPostIdList(uid uint, page int, pageSize int) (bool, []uint) {
	var postIdList []uint = make([]uint, pageSize*2)
	count, followUserList := firstpage.QueryFollowListByUID(uid, page, pageSize)
	if count == 0 {
		return false, nil
	}
	index := 0
	for i := 0; i < len(followUserList); i++ {
		prow, postId := firstpage.QueryNewPostId(followUserList[i])
		frow, forwardId := firstpage.QueryNewForwardId(followUserList[i])
		if prow != 0 {
			postIdList[index] = postId
			index++
		}
		if frow != 0 {
			postIdList[index] = forwardId
			index++
		}
	}
	return true, postIdList[0:index]
}

/*根据自己的uid取得关注的人的帖子详细信息*/
func GetPostDetail(uid uint, page int, pageSize int) (bool, _response.TotalRes) {
	var realRes _response.TotalRes
	pageNum := firstpage.QueryPageNum(uid)
	if pageNum%pageSize == 0 {
		realRes.Page = pageNum / pageSize
	} else {
		realRes.Page = pageNum/pageSize + 1
	}
	var res = firstpage.QueryAllPostByUID(uid, page, pageSize)
	for i := 0; i < len(res); i++ {
		res[i].Username = firstpage.QueryUsernameById(res[i].UID)
	}
	realRes.TotalPost = res
	return true, realRes
}

func GetSingleDetail(pid uint, c *gin.Context) _response.ResSinglePost {
	post := firstpage.QuerySingleDetailById(pid)
	UidOfPost := firstpage.QueryUidByPid(pid)
	post.Username = firstpage.QueryUsernameById(UidOfPost)
	post.PID = pid
	myUid, _ := c.Get("uid")
	post.Photos = firstpage.QueryAllPhotos(pid)
	post.Is_follow = firstpage.QueryIsFollowed(myUid.(uint), UidOfPost)
	post.Is_like = firstpage.QueryIsLiked(pid, myUid.(uint))
	return post
}
