package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"web/models"
)

type Notify struct {
	Message string
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./views/login.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		var user models.User
		user, err := user.GetUserByNameAndPassword(r.Form.Get("name"), r.Form.Get("password"))
		if err != nil {
			message := "账号密码不匹配"
			t, err := template.ParseFiles("./views/login.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, message)
			return
		}
		http.Redirect(w, r, "index", 301)
	}
}
