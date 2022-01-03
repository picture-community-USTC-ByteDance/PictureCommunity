package _request

type CreatePost struct {
	Content string `form:"content" json:"content"`
}
