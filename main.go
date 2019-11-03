package main

import (
	"github.com/gin-gonic/gin"
	rc "my-aplication/sso/controller/register/controller"
	"net/http"
)

func main() {
	r := gin.Default()

	//用户管理模块，增删改查，curd
	register_group := r.Group("/register/user")
	{
		//手机号注册
		register_group.POST("/phone", rc.RegisterUserByPhone)
		//邮箱注册
		register_group.POST("/email", rc.RegisterUserByEmaill)

	}

	//验证码模块
	sms_group := r.Group("/sms")
	{

	}
	http.ListenAndServe(":8083", r)

}
