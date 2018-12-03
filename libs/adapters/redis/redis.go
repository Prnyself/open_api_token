package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"open_api_token/libs/encrypt"
	"open_api_token/settings"
	"strconv"
	"time"
)

var client *redis.Client

func init() {
	var (
		err                  error
		host, port, password string
		redisDb              int
	)
	timeoutChan := make(chan int)

	sec, err := settings.Cfg.GetSection("cache_redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'cache_redis': %v", err)
	}

	host = sec.Key("host").String()
	port = sec.Key("port").String()
	password = sec.Key("password").String()
	redisDb = sec.Key("db").MustInt(0)

	go func() {
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password, // no password set
			DB:       redisDb,  // use default DB
		})
		timeoutChan <- 1
	}()

	// 设置5秒时长, 超时则连接失败
	select {
	case <-timeoutChan:
		return
	case <-time.After(time.Duration(5) * time.Second):
		log.Fatal("Fail to connection redis: ", errors.New("timeout"))
	}
}

func IsConnected() bool {
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("redis not connected: ", err)
		return false
	}
	return pong == "PONG"
}

func GenerateToken(key, secret string) string {
	notSecond := time.Now().Unix()
	signatureCode := fmt.Sprintf("%s%d%s", key, notSecond, secret)
	return encrypt.Sha1(signatureCode)
}

func GetToken(appId int) (string, error) {
	appKey := getAppKey(appId)
	res, err := client.Get(appKey).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		log.Println("redis not connected, err: ", err)
		return "", err
	}
	return res, nil
}

func SetToken(appId int, token string) error {
	tokenKey := getTokenKey(token)
	appKey := getAppKey(appId)
	if err := client.Set(appKey, token, time.Duration(7200)*time.Second).Err(); err != nil {
		return err
	}
	if err := client.Set(tokenKey, appId, time.Duration(7200)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func TokenExpireIn(token string) time.Duration {
	tokenKey := getTokenKey(token)
	res, err := client.TTL(tokenKey).Result()
	if err != nil {
		return 0
	}
	return res
}

func GetAppId(token string) (int, error) {
	tokenKey := getTokenKey(token)
	res, err := client.Get(tokenKey).Result()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func getTokenKey(token string) string {
	return fmt.Sprintf("access_token:%s:app", token)
}

func getAppKey(appId int) string {
	return fmt.Sprintf("app:%d:access_token", appId)
}
