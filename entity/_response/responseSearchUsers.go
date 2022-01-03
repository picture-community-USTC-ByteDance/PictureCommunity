package _response

type ResponseSearchUsers struct {
	Profile  string `json:"profile"`
	Uid      uint   `json:"uid"`
	Nickname string `json:"nickname"`
	Motto    string `json:"motto"`
}
