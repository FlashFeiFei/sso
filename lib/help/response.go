package help

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//200~300的响应
func Response200(c *gin.Context, status int, msg string, data map[interface{}]interface{}) {

	if status >= 200 || status < 300 {
		c.Writer.WriteHeader(status)
	} else {
		panic("状态需要在200 ~ 300之间")
	}

	responseBase(c, msg, data)
}

//400之间的响应
func Response400(c *gin.Context, status int, msg string, data map[interface{}]interface{}) {

	if status >= 400 || status < 500 {
		c.Writer.WriteHeader(status)
	} else {
		panic("状态需要在200 ~ 300之间")
	}

	responseBase(c, msg, data)
}

func responseBase(c *gin.Context, msg string, data map[interface{}]interface{}) {
	result := make(map[string]interface{})

	result["msg"] = msg

	if data != nil {
		result["data"] = data
	}

	responseMsg, _ := json.Marshal(result)
	c.Writer.Write(responseMsg)
}
