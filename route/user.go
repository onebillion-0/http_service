package route

import (
	"github.com/gin-gonic/gin"
	"service/app/user"
)

func IntUserRoute(r *gin.Engine) {
	r.POST("/create_user", user.CreateNewUser)
}
