package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s, _ := template.ParseFiles("./views/index.html")
	s.Execute(w, nil)
}

func Info(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form["id"]) // id
	s, _ := template.ParseFiles("./views/post.html")
	s.Execute(w, nil)
}