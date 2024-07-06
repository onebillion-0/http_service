package school_system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sdk "github.com/onebillion-0/user_sdk"
	"github.com/onebillion-0/user_sdk/biz/application/command"
	"service/models"
	"service/utils"
	"strconv"
)

func Register(c *gin.Context) {
	var acsJson models.SchoolSystemRegister
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
	cmd := &command.SchoolMemberCommand{
		Uid:      uid,
		NickName: acsJson.NickName,
		Gender:   acsJson.Gender,
		Age:      int64(acsJson.Age),
		Password: acsJson.Password,
		Appid:    int64(acsJson.Appid),
		Role:     acsJson.Role,
	}
	err = sdk.SchoolMemberRegister(c, []*command.SchoolMemberCommand{cmd})
	if err != nil {
		utils.SetErrInformation(c, models.StatusFail, models.StatusFailMessage)
	}
	utils.SetSuccessInformation(c)
	return
}
