package _response

type UserData struct {
	FollowersNumber  int `json:"followers_number"`
	FansNumber       int `json:"fans_number"`
	PostsNumber      int `json:"posts_number"`
	CollectionNumber int `json:"collection_number"`
	ForwardNumber    int `json:"forward_number"`
}
