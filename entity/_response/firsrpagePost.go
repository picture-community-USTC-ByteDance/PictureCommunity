package _response

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
type TotalRes struct {
	Page      int
	TotalPost []ResPost
}
type ResSinglePost struct {
	UID         uint
	PID         uint
	Username    string
	Nickname    string
	Profile     string
	Photos      []string
	Content     string
	Like_number int
	Create_time time.Time
	Is_like     bool
	Is_follow   bool
}

type ResPossibleFriends struct {
	UID      uint
	Username string
	Nickname string
	Profile  string
	Motto    string
}
