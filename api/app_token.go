package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"service/models"
)

func GetWXToken() (string, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&js_code=%s&grant_type=client_credential"
	path := fmt.Sprintf(url, models.AppID, models.AppSecret)
	resp, err := http.DefaultClient.Get(path)
	if err != nil {
		return "", err
	}

	var tokenJson models.AppTokenToSessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenJson); err != nil {
		fmt.Println("decoder error...")
		return "", err
	}
	return tokenJson.AccessToken, nil
}

func WXUserPhoneNumber(code string) (string, error) {
	token, err := GetWXToken()
	if err != nil {
		return "", err
	}
	path := "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
	path = fmt.Sprintf(path, token)

	info := make(map[string]interface{})
	info["code"] = code
	bts, _ := json.Marshal(info)

	nr := bytes.NewReader(bts)

	resp, err := http.DefaultClient.Post(path, "\"application/json\"", nr)
	if err != nil {
		return "", err
	}

	var tokenJson models.AppPhoneNumberResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenJson); err != nil {
		fmt.Println("decoder error...")
		return "", err
	}
	return tokenJson.PhoneNumber.PhoneNumber, nil
}
