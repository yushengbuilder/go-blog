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
	var postModel models.Post
	posts := postModel.GetPosts()
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
	var postModel models.Post
	post := postModel.GetPost(id)
	if err != nil {
		fmt.Println(err)
	}
	s, err := template.ParseFiles("./views/post.html")
	if err != nil {
		fmt.Println(err)
	}
	s.Execute(w, post)
}