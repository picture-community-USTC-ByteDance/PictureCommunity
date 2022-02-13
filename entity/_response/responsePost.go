package _response

type ResponsePost struct {
	PID           uint   `json:"pid"`
	TitlePhotoUrl string `json:"title_photo_url"`
}

type ResponseHomePost struct {
	PID           uint   `json:"pid"`
	TitlePhotoUrl string `json:"title_photo_url"`
	CommentNumber int    `json:"comment_number"`
	LikeNumber    int    `json:"like_number"`
}

