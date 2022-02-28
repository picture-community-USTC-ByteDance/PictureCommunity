package service

import (
	"github.com/gin-gonic/gin"
	"picture_community/dao/firstpage"
	"picture_community/response"
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

/*根据上面获得的id取得对应的帖子详细信息*/
func GetPostDetail(uid uint, page int, pageSize int) (bool, []response.ResPost) {
	isOk, postIdList := GetPostIdList(uid, page, pageSize)
	if isOk == false {
		return false, nil
	}
	var res []response.ResPost = make([]response.ResPost, len(postIdList))
	index := 0
	for _, value := range postIdList {
		temp := firstpage.QueryDetailById(value)
		temp.PID = value
		temp.Username = firstpage.QueryUsernameById(temp.UID)
		//commentIdList := firstpage.QueryCommentIdById(value)
		//var com []response.ResComment = make([]response.ResComment, len(commentIdList))
		//cindex := 0
		//for _, val := range commentIdList {
		//	comment := firstpage.QueryCommentById(val)
		//	com[cindex] = comment
		//	cindex++
		//}
		//temp.Comment = com
		res[index] = temp
		index++
	}
	return true, res
}

func GetSingleDetail(pid uint, c *gin.Context) response.ResSinglePost {
	post := firstpage.QuerySingleDetailById(pid)
	post.PID = pid
	myUid, _ := c.Get("uid")
	post.Photos = firstpage.QueryAllPhotos(pid)
	post.Is_like = firstpage.QueryIsLiked(pid, myUid.(uint))
	return post
}
