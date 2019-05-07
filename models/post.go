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
	posts = readFakeDataFile("./fakeData/posts.json")//获取文件中的假数据，转换成json，之后替换成数据库查询
	return posts
}

func (p *Post) GetPost(id int) Post {
	var posts []Post
	var post Post
	posts = readFakeDataFile("./fakeData/posts.json")//获取文件中的假数据，转换成json，之后替换成数据库查询
	for _, value := range posts {//循环筛选文章
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
	fi, _ := os.Open(path)//打开文件
	defer fi.Close()//关闭文件
	fd, _ := ioutil.ReadAll(fi)//读取文件中所有内容
	var posts []Post
	json.Unmarshal([] byte (string(fd)), &posts)//将文件中的内容转换成json
	return posts
}
