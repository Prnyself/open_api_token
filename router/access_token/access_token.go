package access_token

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"open_api_token/libs/adapters/redis"
	"open_api_token/libs/result_handler"
	"open_api_token/model"
)

func Get(c *gin.Context) {
	key := c.Query("key")
	secret := c.Query("secret")
	if key == "" || secret == "" {
		c.JSON(http.StatusOK, gin.H(model.GetMessageByCode(10008)))
		return
	}

	app := model.GetAppByKeySecret(key, secret)
	// 如果app不存在
	if app.Id == 0 {
		c.JSON(http.StatusOK, gin.H(model.GetMessageByCode(10021)))
		return
	}

	token, err := redis.GetToken(app.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H(model.GetMessageByCode(20014)))
		return
	}
	var expiresIn int
	if token != "" {
		expiresIn = int(redis.TokenExpireIn(token).Seconds())
	} else {
		token = redis.GenerateToken(app.AppKey, app.AppSecret)
		err := redis.SetToken(app.Id, token)
		log.Println(err)
		expiresIn = 7200
	}
	res := result_handler.OkResult(map[string]interface{}{
		"access_token": token,
		"expires_in":   expiresIn,
	})
	c.JSON(http.StatusOK, gin.H(res))
}

func Refresh(c *gin.Context) {
	var (
		token     string
		expiresIn int
	)
	key := c.Query("key")
	secret := c.Query("secret")
	if key == "" || secret == "" {
		c.JSON(http.StatusOK, gin.H(model.GetMessageByCode(10008)))
		return
	}

	app := model.GetAppByKeySecret(key, secret)
	// 如果app不存在
	if app.Id == 0 {
		c.JSON(http.StatusOK, gin.H(model.GetMessageByCode(10021)))
		return
	}

	if redis.IsConnected() {
		token = redis.GenerateToken(app.AppKey, app.AppSecret)
		err := redis.SetToken(app.Id, token)
		log.Println(err)
		expiresIn = 7200
	} else {
		c.JSON(http.StatusOK, gin.H(model.GetMessageByCode(20014)))
		return
	}

	res := result_handler.OkResult(map[string]interface{}{
		"access_token": token,
		"expires_in":   expiresIn,
	})
	c.JSON(http.StatusOK, gin.H(res))
}

func Delete(c *gin.Context) {

}
