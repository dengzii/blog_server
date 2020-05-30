package user

import "github.com/dengzii/blog_server/db"

type User struct {
	Name      string `json:"name",gorm:"unique;not null VARCHAR(191)"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Links     string `json:"links"`
	Likes     int32  `json:"likes"`
	Follower  int32  `json:"follower"`
	Following int32  `json:"following"`
	PassWd    string
}

type PageInfo struct {
	Name   string `json:"name",gorm:"unique;not null VARCHAR(191)"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Bio    string `json:"bio"`
}

func GetUser(name string, passwd string) *User {
	var user User
	tab := []interface{}{&User{}}
	db.CreateTable(tab)
	db.Mysql.Where("name = ? AND pass_wd = ?", name, passwd).Find(&user)
	return &user
}

func AddUser(name string, passwd string, email string) *User {
	user := &User{
		Name:      name,
		Email:     email,
		Avatar:    "",
		Bio:       "",
		Links:     "",
		Likes:     0,
		Follower:  0,
		Following: 0,
		PassWd:    passwd,
	}
	db.Insert(user)
	return user
}

func init() {

}
