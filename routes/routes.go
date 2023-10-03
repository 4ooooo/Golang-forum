package routes

import (
	"net/http"

	"github.com/xueyuanjun/chitchat/handlers"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"change",
		"GET",
		"/change",
		handlers.Change,
	},
	{
		"update",
		"POST",
		"/change_auth",
		handlers.Change_Auth,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"newThread",
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	{
		"createThread",
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},
	{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},

	{
		"select_thread",
		"GET",
		"/select_thread",
		handlers.Select_Thread,
	},
	{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
	{
		"admin",
		"GET",
		"/admin",
		handlers.Select_Thread,
	},
	{
		"deleteuser",
		"POST",
		"/deleteuser",
		handlers.DeleteUser,
	},
	{
		"deletetopic",
		"POST",
		"/deletetopic",
		handlers.DeleteTopic,
	},
	{
		"like",
		"POST",
		"/like",
		handlers.Like,
	},
	{
		"sci",
		"GET",
		"/sci",
		handlers.SCI,
	},
	{
		"xinde",
		"GET",
		"/xinde",
		handlers.XinDe,
	},
	{
		"meeting",
		"GET",
		"/meeting",
		handlers.Meeting,
	},
	{
		"chihe",
		"GET",
		"/chihe",
		handlers.ChiHe,
	},
}
