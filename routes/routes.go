package routes

import (
	"net/http"
	"web/controllers"
)

func Route() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) //处理静态文件

	http.HandleFunc("/", controllers.Index)      //首页
	http.HandleFunc("/post", controllers.Info)   //详情页
	http.HandleFunc("/about", controllers.About) //关于页
}
