package handlers

import (
	"github.com/xueyuanjun/chitchat/models"
	"net/http"
)

func SCI(writer http.ResponseWriter, request *http.Request) {
	threads, _ := models.Select("sci")
	sess, _ := session(writer, request)
	if sess.Email == "admin@admin.com" {
		generateHTML(writer, threads, "layout", "ad.navbar", "sci")
	} else {
		generateHTML(writer, threads, "layout", "navbar", "sci")
	}
}

func XinDe(writer http.ResponseWriter, request *http.Request) {
	threads, _ := models.Select("xinde")
	sess, _ := session(writer, request)
	if sess.Email == "admin@admin.com" {
		generateHTML(writer, threads, "layout", "ad.navbar", "xinde")
	} else {
		generateHTML(writer, threads, "layout", "navbar", "xinde")
	}
}

func Meeting(writer http.ResponseWriter, request *http.Request) {
	threads, _ := models.Select("meeting")
	sess, _ := session(writer, request)
	if sess.Email == "admin@admin.com" {
		generateHTML(writer, threads, "layout", "ad.navbar", "meeting")
	} else {
		generateHTML(writer, threads, "layout", "navbar", "meeting")
	}
}

func ChiHe(writer http.ResponseWriter, request *http.Request) {
	threads, _ := models.Select("joy")
	sess, _ := session(writer, request)
	if sess.Email == "admin@admin.com" {
		generateHTML(writer, threads, "layout", "ad.navbar", "chihe")
	} else {
		generateHTML(writer, threads, "layout", "navbar", "chihe")
	}
}
