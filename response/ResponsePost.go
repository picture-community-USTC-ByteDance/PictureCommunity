package response

import "time"

/*帖子*/
type ResPost struct {
	PID             uint
	Nickname        string
	Profile         string
	Title_photo_url string
	Content         string
	Like_number     int
	Comment_number  int
	Create_time     time.Time
	Comment         []ResComment
}
