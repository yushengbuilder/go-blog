package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"web/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	posts := models.GetPostAll()
	s, err := template.ParseFiles("./views/index.html")
	if err != nil {
		fmt.Println(err)
	}
	s.Execute(w, posts)
}

func Info(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil {
		fmt.Println(err)
	}
	post :=models.GetPostInfo(id)
	if err != nil {
		fmt.Println(err)
	}
	s, err := template.ParseFiles("./views/post.html")
	if err != nil {
		fmt.Println(err)
	}
	s.Execute(w, post)
}