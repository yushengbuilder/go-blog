package routes

import (
	"net/http"
	"web/controllers"
)

func Route () {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/post", controllers.Info)
	http.HandleFunc("/about", controllers.About)
}
