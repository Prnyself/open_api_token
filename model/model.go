package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"open_api_token/settings"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
}

func init() {
	var (
		err                                        error
		dbType, dbName, user, password, host, port string
	)
	timeoutChan := make(chan int)

	sec, err := settings.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("type").String()
	dbName = sec.Key("database").String()
	user = sec.Key("user").String()
	password = sec.Key("password").String()
	host = sec.Key("host").String()
	port = sec.Key("port").String()

	go func() {
		connConf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			port,
			dbName)

		db, err = gorm.Open(dbType, connConf)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		timeoutChan <- 1
	}()

	// 设置5秒时长, 超时则连接失败
	select {
	case <-timeoutChan:
		return
	case <-time.After(time.Duration(5) * time.Second):
		log.Fatal("Fail to connection database: ", errors.New("timeout"))
	}
}
