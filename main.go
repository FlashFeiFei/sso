package main

import (
	"github.com/guapo-organizations/sso/user"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func validateTicket(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {

	//登录
	http.HandleFunc("/login", loginHandler)
	//注册
	http.HandleFunc("register", user.RegisterHandler)

	//认证
	http.HandleFunc("/validateTicket", loginHandler)

	//开启服务
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
