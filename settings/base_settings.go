package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	LogSavePath  string
	LogSaveName  string
	LogFileExt   string
	TimeFormat   string
	DebugLogName string
)

func init() {
	var err error
	//log.Println(os.)
	Cfg, err = ini.Load("./config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/1app.ini': %v", err)
	}
	loadApp()
	loadLogs()
}
