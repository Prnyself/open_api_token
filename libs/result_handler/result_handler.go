package result_handler

type ResultHandler struct {
	Code        int
	Message     string
	Description string
	Result      interface{}
}

func OkResult(result interface{}) map[string]interface{} {
	res := &ResultHandler{
		Code:        10000,
		Message:     "ok",
		Description: "请求成功",
		Result:      result,
	}
	return res.toHash()
}

func (handler *ResultHandler) toHash() map[string]interface{} {
	res := make(map[string]interface{}, 4)
	res["code"] = handler.Code
	res["message"] = handler.Message
	res["description"] = handler.Description
	res["result"] = handler.Result
	return res
}
