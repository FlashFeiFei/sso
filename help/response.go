package help

import (
	"encoding/json"
	"net/http"
)

//200~300的响应
func Response200(w http.ResponseWriter, status int, msg string, data map[interface{}]interface{}) {

	if status >= 200 || status < 300 {
		w.WriteHeader(status)
	} else {
		panic("状态需要在200 ~ 300之间")
	}

	responseBase(w, msg, data)
}

//400之间的响应
func Response400(w http.ResponseWriter, status int, msg string, data map[interface{}]interface{}) {

	if status >= 400 || status < 500 {
		w.WriteHeader(status)
	} else {
		panic("状态需要在200 ~ 300之间")
	}

	responseBase(w, msg, data)
}

func responseBase(w http.ResponseWriter, msg string, data map[interface{}]interface{}) {
	result := make(map[string]interface{})
	result["msg"] = msg

	if data != nil {
		result["data"] = data
	}

	responseMsg, _ := json.Marshal(result)

	w.Write(responseMsg)
}
