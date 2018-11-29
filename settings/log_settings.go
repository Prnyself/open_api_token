package settings

import "log"

func loadLogs() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get section 'log': %v", err)
	}

	LogSavePath = sec.Key("log_save_path").MustString("logs/")
	LogSaveName = sec.Key("log_save_name").MustString("gin")
	LogFileExt = sec.Key("log_file_ext").MustString("log")
	TimeFormat = sec.Key("time_format").MustString("2006-01-02")

	secDebug, _ := Cfg.GetSection("log.debug")
	if err != nil {
		log.Fatalf("Fail to get section 'log.debug': %v", err)
	}
	DebugLogName = secDebug.Key("log_save_name").MustString("debug")
}
