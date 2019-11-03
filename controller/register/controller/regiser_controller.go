package controller

import (
	"github.com/gin-gonic/gin"
	"my-aplication/sso/controller/register"
	"my-aplication/sso/lib/help"
	"net/http"
)

//通过手机号注册
func RegisterUserByPhone(c *gin.Context) {
	var userInfo regster.UserRegisterInfoByPhone
	if c.Bind(&userInfo) != nil {
		//我没有用gin自带的返回响应，我自己封装了原生的响应
		help.Response400(c, http.StatusBadRequest, "解析注册数据失败", nil)
		return
	}

	if !help.VerifyMobileFormat(userInfo.Phone) {
		help.Response400(c, http.StatusBadRequest, "手机号格式不对", nil)
		return
	}

	//数据入库

	//反应响应
	help.Response200(c, http.StatusCreated, "注册用户成功", nil)
	return
}

//通过email注册
func RegisterUserByEmaill(c *gin.Context) {
	var userInfo regster.UserRegisterInfoByEmail
	if c.Bind(&userInfo) != nil {
		//我没有用gin自带的返回响应，我自己封装了原生的响应
		help.Response400(c, http.StatusBadRequest, "解析注册数据失败", nil)
		return
	}

	if !help.VerifyEmailFormat(userInfo.Email) {
		help.Response400(c, http.StatusBadRequest, "email格式不对", nil)
		return
	}

	//数据入库

	//反应响应
	help.Response200(c, http.StatusCreated, "注册用户成功", nil)
	return
}
