package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Login (w http.ResponseWriter, r *http.Request) {
	if r.Method =="GET" {
		t, err := template.ParseFiles("./views/login.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println(r.Form.Get("name"))
		fmt.Println(r.Form.Get("password"))
	}
}