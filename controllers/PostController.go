package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"web/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()//解析请求参数
	var postModel models.Post//获取文章model
	posts := postModel.GetPosts()//获取所有文章
	s, err := template.ParseFiles("./views/index.html")//解析模版
	if err != nil {
		fmt.Println(err)
	}
	s.Execute(w, posts)//合并模版
}

func Info(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()//解析请求参数
	id, err := strconv.Atoi(r.Form["id"][0])//转换请求中id为数字
	if err != nil {
		fmt.Println(err)
	}
	var postModel models.Post//获取文章model
	post := postModel.GetPost(id)//获取文章详情
	s, err := template.ParseFiles("./views/post.html")//解析模版
	if err != nil {
		fmt.Println(err)
	}
	s.Execute(w, post)//合并模版
}