package user

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	sdk "github.com/oneliuliu/user_sdk"
	"github.com/oneliuliu/user_sdk/biz/interface/sdk/dto/request"
	"net/http"
	api2 "service/api"
	"service/models"
	"service/utils"
	"strconv"
)

func Login(c *gin.Context) {
	var acsJson models.AppCode2SessionJson
	code := c.PostForm("code")
	acs := models.AppCode2Session{
		Code:      code,
		AppId:     models.AppID,
		AppSecret: models.AppSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	path := fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code)
	res, err := http.DefaultClient.Get(path)
	if err != nil {
		fmt.Println("微信登录凭证校验接口请求错误")
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}
	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
		fmt.Println("decoder error...")
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}

	phoneNumber, err := api2.WXUserPhoneNumber(code)
	if err != nil {
		fmt.Println("获取用户手机号失败")
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}
	exist, err := sdk.CheckUserRegisteredByPhoneNumber(phoneNumber)
	if err != nil {
		fmt.Println("查询用户状态失败")
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}
	if !exist {
		uid, err := utils.GSnowFlake.NextId()
		if err != nil {
			fmt.Println("uid err:%v", err)
			utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
			return
		}
		req := request.CreateUserRequest{
			PhoneNumber: phoneNumber,
			Uid:         strconv.FormatInt(uid, 10),
		}
		err = sdk.CreateUser(&req)
		if err != nil {
			utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
			return
		}

		utils.SetErrInformation(c, models.StatusSuccess, models.StatusSuccessMessage)
	}
}

func CreateNewUser(ctx *gin.Context) {
	userCommonInfoJson, exit := ctx.GetPostForm("common_info")
	if !exit {
		utils.SetErrInformation(ctx, models.StatusInvalidParam, models.StatusInvalidParamMessage)
		return
	}
	userCommonInfo := &models.BaseInfo{}
	err := sonic.UnmarshalString(userCommonInfoJson, userCommonInfo)
	if err != nil {
		utils.SetErrInformation(ctx, models.StatusUnmarshalErr, models.StatusUnmarshalErrMessage)
		return
	}
	//TODO: 校验入参
	//userInfo := &models.UserInfo{}
	//userInfo.BaseInfo = userCommonInfo
	req := request.CreateUserRequest{
		Uid:         userCommonInfo.Uid,
		NickName:    userCommonInfo.NickName,
		Avatar:      userCommonInfo.Avatar,
		Sex:         userCommonInfo.Sex,
		PhoneNumber: userCommonInfo.PhoneNumber,
		IdCard:      userCommonInfo.IdCard,
		Age:         strconv.Itoa(int(userCommonInfo.Age)),
	}
	err = sdk.CreateUser(&req)
	if err != nil {
		utils.SetErrInformation(ctx, models.StatusFail, models.StatusFailMessage)
		return
	}
	utils.SetErrInformation(ctx, models.StatusSuccess, models.StatusSuccessMessage)
	//userName, exist := ctx.GetQuery("name")
	//if !exist {
	//	utils.SetErrInformation(ctx, models.StatusInvalidParam, models.StatusInvalidParamMessage)
	//}
	//userId, exist := ctx.GetQuery("uid")
	//if !exist {
	//	utils.SetErrInformation(ctx, models.StatusInvalidParam, models.StatusInvalidParamMessage)
	//}
	//avatar, exist := ctx.GetQuery("avatar")
	//if !exist {
	//	Avatar = "default.url"
	//}
}
