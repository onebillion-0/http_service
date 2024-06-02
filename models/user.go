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

// 微信小程序，凭证校验后返回的JSON数据包模型
type AppCode2SessionJson struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// 微信小程序，用户登录凭证校验模型
type AppCode2Session struct {
	Code      string
	AppId     string
	AppSecret string
}
