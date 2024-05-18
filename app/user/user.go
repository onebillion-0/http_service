package user

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	sdk "github.com/oneliuliu/user_sdk"
	"github.com/oneliuliu/user_sdk/biz/interface/sdk/dto/request"
	"service/models"
	"service/utils"
	"strconv"
)

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
