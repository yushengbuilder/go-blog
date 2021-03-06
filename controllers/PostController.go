package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"web/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                                       //解析请求参数
	var postModel models.Post                           //获取文章model
	posts := postModel.GetPosts()                       //获取所有文章
	s, err := template.ParseFiles("./views/index.html") //解析模版
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Execute(w, posts) //合并模版
}

func Info(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                            //解析请求参数
	id, err := strconv.Atoi(r.Form.Get("id")) //转换请求中id为数字
	if err != nil {
		fmt.Println(err)
		return
	}
	var postModel models.Post          //获取文章model
	post, err := postModel.GetPost(id) //获取文章详情
	if err != nil {
		fmt.Println(err)
		t, _ := template.ParseFiles("./views/404.html")
		t.Execute(w, nil)                            //合并模版
		return
	}

	s, err := template.ParseFiles("./views/post.html") //解析模版
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Execute(w, post) //合并模版
}
