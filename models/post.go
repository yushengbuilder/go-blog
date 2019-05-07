package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type post struct {
	Id      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

func GetAll() []post {

	var posts []post
	posts = readFakeDataFile("./fakeData/posts.json")
	return posts
}

func GetInfo(id int) interface{} {
	var posts []post
	posts = readFakeDataFile("./fakeData/posts.json")
	for _, value := range posts {
		if id == value.Id {
			return value
		} else {
			continue
		}
	}
	return nil
}

func readFakeDataFile(path string) []post {
	fi, _ := os.Open(path)
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)
	var posts []post
	json.Unmarshal([] byte (string(fd)), &posts)
	return posts
}
