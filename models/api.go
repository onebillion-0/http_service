package models

var AppTokenSession = AppTokenToSession{
	GrantType: "client_credential",
	Appid:     AppID,
	Secret:    AppSecret,
}

type AppTokenToSession struct {
	GrantType string `json:"grant_type"`
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
}

type AppTokenToSessionResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type AppPhoneNumberResponse struct {
	ErrCode     int         `json:"errcode"`
	ErrMsg      string      `json:"errmsg"`
	PhoneNumber PhoneNumber `json:"phone_info"`
}

type PhoneNumber struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}
