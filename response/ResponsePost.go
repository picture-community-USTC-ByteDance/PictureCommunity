package response

/*帖子*/
type ResPost struct {
	Nickname        string
	Profile         string
	Title_photo_url string
	Content         string
	Like_number     int
	Comment_number  int
	Comment         []ResComment
}
