package model

type CampusApp struct {
	Id        int
	AppKey    string
	AppSecret string
}

func GetAppByKeySecret(key, secret string) CampusApp {
	var app CampusApp
	db.First(&app, &CampusApp{AppKey: key, AppSecret: secret})
	return app
}

func GetAppById(id int) CampusApp {
	var app CampusApp
	db.First(&app, id)
	return app
}
