package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s, err := template.ParseFiles("./views/about.html")
	if err != nil {
		fmt.Println(err)
	}
	s.Execute(w, nil)
}