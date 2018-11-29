package settings

import "time"

func loadApp() {
	RunMode = Cfg.Section("").Key("app_mode").MustString("debug")
	HTTPPort = Cfg.Section("app").Key("port").MustInt(3309)
	ReadTimeout = time.Duration(Cfg.Section("app").Key("read_timeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(Cfg.Section("app").Key("write_timeout").MustInt(60)) * time.Second
}
