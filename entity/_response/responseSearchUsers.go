package _response

type ResponseSearchUsers struct {
	Uid      uint   `json:"uid"`
	Username string `json:"username"`
	Profile  string `json:"profile"`
	Nickname string `json:"nickname"`
	Motto    string `json:"motto"`
}
