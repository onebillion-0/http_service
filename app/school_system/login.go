package school_system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	sdk "github.com/onebillion-0/user_sdk"
	"github.com/onebillion-0/user_sdk/biz/constants"
	"net/http"

	"service/models"
	"service/utils"
	"strconv"
)

func Login(c *gin.Context) {
	var acsJson models.SchoolMemberLoginInfo
	err := c.BindJSON(&acsJson)
	if err != nil {
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}

	uid, err := strconv.ParseInt(acsJson.Uid, 10, 64)
	if err != nil {
		fmt.Println("parse uid error", err)
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}
	token, err := sdk.SchoolMemberLogin(c, uid, acsJson.Password)
	if err != nil {
		if errors.Is(err, constants.ERROR_INVALID_USERNAME_OR_PASSWORD) {
			utils.SetErrInformation(c, models.StatusInvalidUserOrPassword, models.StatusInvalidUserOrPasswordMessage)
			return
		}
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	utils.SetSuccessInformation(c)
	return
}

func GetAppid(c *gin.Context) {
	appidList, err := sdk.GetAppIDList(c)
	if err != nil {
		fmt.Println("get appid list error", err)
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
		return
	}
	c.JSON(http.StatusOK, appidList)
}
