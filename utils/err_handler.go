package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/models"
)

func SetErrInformation(ctx *gin.Context, code int, message string) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    code,
			"message": message,
		})
}

func SetSuccessInformation(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    0,
			"message": models.StatusSuccessMessage,
		})
}
