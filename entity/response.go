package entity

type ResponseSearchUsers struct {
	Profile  string `json:"profile"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Motto    string `json:"motto"`
}
