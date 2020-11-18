package main

import (
	"fmt"
	"net/http"
)


func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>临兵斗者，皆阵列前行。前方高能预警</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系"+"<a href=\"mailto:jinhesui@163.com\">jinhesui@163.com</a>")
	} else {
		fmt.Fprint(w, "<h1>请求页面未找到：（</h1>"+"<p>如有疑惑，请联系我们。</p>")
	}
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}