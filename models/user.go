package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type User struct {
	Id       int    `json:id`
	Name     string `json:name`
	Password string `json:password`
}

func (u *User) GetUsers() []User {
	var posts []User
	posts = u.readFakeDataFile("./fakeData/users.json") //获取文件中的假数据，转换成json，之后替换成数据库查询
	return posts
}

func (u *User) GetUser(id int) User {
	var users []User
	var user User
	users = u.readFakeDataFile("./fakeData/users.json") //获取文件中的假数据，转换成json，之后替换成数据库查询
	for _, value := range users { //循环筛选文章
		if id == value.Id {
			user = value
		} else {
			continue
		}
	}
	return user
}

func (u *User) GetUserByNameAndPassword(n, p string) (User, error) {
	var users []User
	var user User
	users = u.readFakeDataFile("./fakeData/users.json") //获取文件中的假数据，转换成json，之后替换成数据库查询
	for _, value := range users { //循环筛选文章
		if n == value.Name && p == value.Password {
			return value, nil
		} else {
			continue
		}
	}
	return user, errors.New("账号密码不匹配")
}

//此处之后替换成数据库查询
func (u *User) readFakeDataFile(path string) []User {
	fi, _ := os.Open(path)      //打开文件
	defer fi.Close()            //关闭文件
	fd, _ := ioutil.ReadAll(fi) //读取文件中所有内容
	var users []User
	json.Unmarshal([] byte (string(fd)), &users) //将文件中的内容转换成json
	return users
}
