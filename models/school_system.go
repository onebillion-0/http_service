package models

import "github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"

type SchoolMemberLoginInfo struct {
	AppID    int    `json:"app_id"`
	Uid      string `json:"uid"`
	Password string `json:"password"`
}

type SchoolSystemRegister struct {
	NickName string              `json:"nick_name"`
	Uid      string              `json:"uid"`
	Age      int                 `json:"age"`
	Password string              `json:"password"`
	Appid    int                 `json:"appid"`
	Gender   string              `json:"gender"`
	Role     school_members.Role `json:"role"`
}

type Appid struct {
	Appid int `json:"appid"`
}
