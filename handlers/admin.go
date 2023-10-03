package handlers

import (
	"fmt"
	"github.com/xueyuanjun/chitchat/models"
	"net/http"
	"net/url"
	"strconv"
)

//	func Admin(writer http.ResponseWriter, request *http.Request) {
//		generateHTML(writer, nil, "auth.layout", "navbar", "admin")
//	}
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var sess, err = session(writer, request)
	user, err := sess.User()
	username := user.Name
	if err != nil || username != "admin" {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err := request.ParseForm()
		if request.PostFormValue("Id") == "1" {
			danger("Cannot delete admin")
			http.Redirect(writer, request, "/admin", 302)
			return
		}
		id, err := strconv.Atoi(request.PostFormValue("id"))
		err = models.DeleteUser(id)
		if err != nil {
			danger(err, "Cannot delete user")
		}
		err = models.DeleteSession(id)
		http.Redirect(writer, request, "/admin", 302)
	}
}

func DeleteTopic(writer http.ResponseWriter, request *http.Request) {
	var sess, err = session(writer, request)
	user, err := sess.User()
	username := user.Name
	if err != nil || username != "admin" {
		http.Redirect(writer, request, "/login", 302)
	} else {

		err := request.ParseForm()
		id, err := strconv.Atoi(request.PostFormValue("id"))
		err = models.DeleteTopic(id)
		if err != nil {
			danger(err, "Cannot delete topic")
		}
		err = models.DeletePost(id)
		if err != nil {
			danger(err, "Cannot delete topic")
		}
		http.Redirect(writer, request, "/admin", 302)
	}
}

func Select_Thread(writer http.ResponseWriter, request *http.Request) {
	result, err := models.Union()
	sess, err := session(writer, request)
	if err == nil && sess.Email == "admin@admin.com" {
		generateHTML(writer, result, "layout", "ad.navbar", "admin")
	} else {
		errorMessage := "对不起，您不是管理员。"
		redirectURL := fmt.Sprintf("/?error=%s", url.QueryEscape(errorMessage))
		http.Redirect(writer, request, redirectURL, http.StatusSeeOther)
	}
}
