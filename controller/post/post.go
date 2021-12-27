package post

import (
	"fmt"
	"net/http"
	"picture_community/entity"
	"picture_community/service/post"

	"github.com/gin-gonic/gin"
)

func CreatePostController(c *gin.Context) {
	var u entity.CreatePost
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": err.Error(),
			"post_id": "",
		})
		return
	}
	fmt.Println("controller: ", u.Content, u.ImageUrl, u.ID)
	status, message, postid := post.CreatePost(u)
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"post_id": postid,
	})
}
