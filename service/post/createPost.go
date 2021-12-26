package post

import (
	"fmt"
	daopost "picture_community/dao/post"
	"picture_community/entity"
)

func CreatePost(parm entity.CreatePost) (status int, message string, post_id int64) {
	fmt.Println("contents::", parm.ID, parm.ImageUrl, parm.Content)
	newPost := entity.Post{
		UID:              parm.ID,
		TitlePhotoUrl:    parm.ImageUrl,
		Content:          parm.Content,
		PhotoNumber:      1,
		CommentNumber:    0,
		LikeNumber:       0,
		ForwardNumber:    0,
		CollectionNumber: 0,
	}
	postID, err := daopost.InsertPostByUserID(newPost)
	if err != nil {
		return -1, err.Error(), 0
	}
	return 0, "create post success", postID
}
