package handlers

import (
	"fmt"
	"github.com/xueyuanjun/chitchat/models"
	"net/http"
	"net/url"
)

func Change(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "change")
}

func Change_Auth(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}

	newpassword := request.FormValue("newpassword")
	sfz := request.FormValue("sfz")
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Sfz == sfz {
		user.Password = models.Encrypt(newpassword)
		if err := user.Update(); err != nil {
			danger(err, "Update")
		}
		http.Redirect(writer, request, "/login", 302)
	} else {
		// 显示错误消息并重新渲染页面
		errorMessage := "密钥错误，请重新尝试。"
		redirectURL := fmt.Sprintf("/change?error=%s", url.QueryEscape(errorMessage))
		http.Redirect(writer, request, redirectURL, http.StatusSeeOther)
	}
}
