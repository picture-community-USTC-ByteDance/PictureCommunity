package response

import "time"

/*帖子*/
type ResPost struct {
	UID             uint
	Username        string
	PID             uint
	Nickname        string
	Profile         string
	Title_photo_url string
	Content         string
	Like_number     int
	Comment_number  int
	Create_time     time.Time
	//Comment         []ResComment
}

type ResSinglePost struct {
	PID         uint
	Nickname    string
	Profile     string
	Photos      []string
	Content     string
	Like_number int
	Create_time time.Time
	Is_like     bool
}
