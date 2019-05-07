package main

import (
	"net/http"
	"web/routes"
)

func main() {
	//引入路由
	routes.Route()
	//监听端口
	http.ListenAndServe(":80", nil)

}





