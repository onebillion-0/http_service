package user

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	sdk "github.com/oneliuliu/user_sdk"
	"github.com/oneliuliu/user_sdk/domain/model"
	"service/models"
	"service/utils"
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
	info := &model.UserInfo{
		Uid:      userCommonInfo.Uid,
		NickName: userCommonInfo.NickName,
	}
	err = sdk.CreateUser(info)
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
