package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"service/models"
	"strings"
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
	path := " https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
	path = fmt.Sprintf(path, token)
	val := url.Values{}
	val.Set("code", code)
	body := strings.NewReader(val.Encode())
	resp, err := http.DefaultClient.Post(path, "\"application/x-www-form-urlencoded\"", body)
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
