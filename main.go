package main

import (
	"log"
	"net/http"
	"web/routes"
)


func main() {

	routes.Route()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}
}





