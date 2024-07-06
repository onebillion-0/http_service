package route

import (
	"github.com/gin-gonic/gin"
	"service/app/school_system"
	"service/app/user"
)

func IntUserRoute(r *gin.Engine) {
	// 微信小程序用户接口
	app := r.Group("/app")
	{
		app.POST("/login", user.Login)
	}
	r.POST("/create_user", user.CreateNewUser)

	school := r.Group("/school_system")
	{
		school.POST("login", school_system.Login)
		school.POST("register", school_system.RegisterMember)
		school.POST("register_appid", school_system.RegisterAppid)
		school.GET("get_appid", school_system.GetAppid)
	}

}
