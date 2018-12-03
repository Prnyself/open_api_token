package model

type ApiMessage struct {
	Model
	Code        int
	Name        string
	Description string
}

func GetMessageByCode(code int) map[string]interface{} {
	msg := ApiMessage{}
	db.Where(&ApiMessage{Code: code}).First(&msg)
	return msg.toHash()
}

func (apiMessage *ApiMessage) toHash() map[string]interface{} {
	res := make(map[string]interface{}, 3)
	res["code"] = apiMessage.Code
	res["message"] = apiMessage.Name
	res["description"] = apiMessage.Description
	return res
}
