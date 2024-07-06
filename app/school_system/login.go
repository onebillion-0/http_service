package school_system

import (
	"github.com/gin-gonic/gin"
	"service/models"
)

func Login(c *gin.Context) {
	var acsJson models.SchoolMemberLoginInfo
	err := c.BindJSON(&acsJson)

}
