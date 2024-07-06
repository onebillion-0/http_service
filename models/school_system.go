package models

type SchoolMemberLoginInfo struct {
	AppID    int    `json:"app_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
