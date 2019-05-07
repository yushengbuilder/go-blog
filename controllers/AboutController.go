package controllers

import (
	"html/template"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s, _ := template.ParseFiles("./views/about.html")
	s.Execute(w, nil)
}