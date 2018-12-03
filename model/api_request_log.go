package model

type ApiRequestLog struct {
	Model
	IsFail      int `json:"is_fail"`
	ApiId       int `json:"api_id"`
	CampusAppId int `json:"campus_app_id"`
	UserId      int `json:"user_id"`
}

func GetAllLogs(limit int) []ApiRequestLog {
	var logs []ApiRequestLog
	db.Limit(limit).Find(&logs)
	return logs
}
