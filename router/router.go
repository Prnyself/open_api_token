package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"open_api_token/libs/result_handler"
	"open_api_token/model"
	"open_api_token/router/access_token"
	"open_api_token/settings"
	"os"
	"strconv"
)

func InitRouter() *gin.Engine {
	f, _ := os.Create(fmt.Sprintf("%s%s.%s", settings.LogSavePath, settings.LogSaveName, settings.LogFileExt))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	openApi := r.Group("/open_api")
	auth := openApi.Group("/authentication")
	auth.Use()
	{
		auth.GET("/get_access_token", access_token.Get)
		auth.GET("/refresh_access_token", access_token.Refresh)
		auth.DELETE("/delete_access_token", access_token.Delete)
	}

	openApi.GET("/test/say_hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "hello world",
		})
	})

	openApi.GET("/request_log", func(context *gin.Context) {
		limit, _ := strconv.Atoi(context.Query("per_page"))
		logs := model.GetAllLogs(limit)
		result := result_handler.OkResult(map[string]interface{}{
			"result": logs,
		})
		context.JSON(http.StatusOK, result)
	})

	return r
}
