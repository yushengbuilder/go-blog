package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Post struct {
	Id      int    `json:id`
	Detail  string `json:detail`
	Title   string `json:title`
	Content string `json:content`
}

func (p *Post) GetPosts() []Post {
	var posts []Post
	posts = readFakeDataFile("./fakeData/posts.json")
	return posts
}

func (p *Post) GetPost(id int) Post {
	var posts []Post
	var post Post
	posts = readFakeDataFile("./fakeData/posts.json")
	for _, value := range posts {
		if id == value.Id {
			post = value
		} else {
			continue
		}
	}
	return post
}

//此处之后替换成数据库查询
func readFakeDataFile(path string) []Post {
	fi, _ := os.Open(path)
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)
	var posts []Post
	json.Unmarshal([] byte (string(fd)), &posts)
	return posts
}
