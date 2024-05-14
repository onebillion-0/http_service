package models

type UserInfo struct {
	BaseInfo *BaseInfo `json:"commonInfo"`
}

type BaseInfo struct {
	NickName    string `json:"nick_name"`
	Uid         string `json:"uid"`
	Avatar      string `json:"avatar"`
	Sex         string `json:"sex"`
	PhoneNumber string `json:"phone_number"`
	IdCard      string `json:"id_card"`
	Age         int64  `json:"age"`
}
