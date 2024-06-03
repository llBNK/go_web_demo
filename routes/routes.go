package routes

import (
	"net/http"
	"web_demo/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinLogger())

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})

	return r

}

type mapValue struct {
	Numbers []int `json:"numbers"`
}
