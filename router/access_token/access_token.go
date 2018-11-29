package access_token

import (
	"github.com/gin-gonic/gin"
	"open_api_token/libs/logger"
)

func Get(c *gin.Context) {
	logger.Info("test")
}

func Refresh(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
