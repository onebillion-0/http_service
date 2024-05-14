package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetErrInformation(ctx *gin.Context, code int, message string) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    code,
			"message": message,
		})
}
