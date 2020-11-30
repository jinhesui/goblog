package main

import (
	"database/sql"
	"fmt"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router
var db *sql.DB

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>临兵斗者，皆阵列前行。</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系"+"<a href=\"mailto:jinhesui@163.com\">jinhesui@163.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到：（</h1><p>如有疑惑，请联系我们。</p>")
}

// Article 对应一条文章数据
type Article struct {
	Title, Body string
	ID          int64
}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
