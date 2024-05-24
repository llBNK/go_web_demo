package routes

import (
	"net/http"
	"web_demo/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinLogger())

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})

	return r

}
