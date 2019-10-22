package main

import (
	"github.com/guapo-organizations/sso/user"
	"log"
	"net/http"
)

func validateTicket(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {

	//登录
	http.HandleFunc("/user/login", user.LoginHandler)
	//注册
	http.HandleFunc("/user/register", user.RegisterHandler)

	//认证
	http.HandleFunc("/validateTicket", validateTicket)

	//开启服务
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
