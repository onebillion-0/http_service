package route

import (
	"github.com/gin-gonic/gin"
	"service/app/user"
)

func IntUserRoute(r *gin.Engine) {
	// 微信小程序用户接口
	app := r.Group("/app")
	{
		app.POST("/login", user.Login)
	}

	r.POST("/create_user", user.CreateNewUser)
}
