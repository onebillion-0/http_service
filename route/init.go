package route

import "github.com/gin-gonic/gin"

func InitRoute() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	IntUserRoute(r)
	r.Run(":8080")
}
