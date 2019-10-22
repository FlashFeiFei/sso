package user

import (
	"github.com/guapo-organizations/sso/help"
	"net/http"
	"strings"
)


func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	//不是post方法
	if !strings.EqualFold(req.Method, "POST") {
		w.Write([]byte("方法不是post"))
		help.Response400(w, http.StatusMethodNotAllowed, "方法不是post", nil)
		return
	}

	username := req.PostFormValue("username")
	password := req.PostFormValue("password")

	_, ok := users.Load(username)

	if ok {
		help.Response400(w, http.StatusNotAcceptable, "用户已存在", nil)
		return
	}

	users.Store("username", password)

	help.Response200(w, http.StatusCreated, "注册成功", nil)
	return
}
