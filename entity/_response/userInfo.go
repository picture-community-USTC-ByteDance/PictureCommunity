package _response

// 个人主页展示的的用户基本信息
type UserData struct {
	Uid              uint   `json:"uid"`
	Nickname         string `json:"nickname"`
	Username         string `json:"username"`
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
	UID           uint   `json:"uid"`
	Nickname      string `json:"nickname"`
	Sex           bool   `json:"sex"`
	Birthday      string `json:"birthday"`
	Address       string `json:"address"`
	Motto         string `json:"motto"`
	Profile       string `json:"profile"`
	OriginProfile string `json:"origin_profile"`
	Email         string `json:"email"`
	Telephone     uint   `json:"telephone"`
}
