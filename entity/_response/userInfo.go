package _response

// 个人主页展示的的用户基本信息
type UserData struct {
	Nickname         string `json:"nickname"`
	Sex              bool   `json:"sex"`
	Profile          string `json:"profile"`
	Motto            string `json:"motto"`
	FollowersNumber  int    `json:"followers_number"`
	FansNumber       int    `json:"fans_number"`
	PostsNumber      int    `json:"posts_number"`
	CollectionNumber int    `json:"collection_number"`
	ForwardNumber    int    `json:"forward_number"`
}

type UserDetail struct {
	UID           uint
	Nickname      string
	Sex           bool //false为女  true为男
	Birthday      string
	Address       string
	Motto         string
	Profile       string //略缩图
	OriginProfile string //详细头像url
}
