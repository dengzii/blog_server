package user

import "server/db"

type User struct {
	ID         uint
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Bio        string `json:"bio"`
	Links      string `json:"links"`
	Likes      int32  `json:"likes"`
	Follower   int32  `json:"follower"`
	Following  int32  `json:"following"`
	PassWd     string
	CreateDate string
}

type UserJson struct {
	Name    string `json:"name"`
	Passwd  string `json:"passwd"`
	Email   string `json:"email"`
	Captcha string `json:"captcha"`
}

func GetUser(name string, passwd string) *User {
	var user User
	tab := []interface{}{&User{}}
	db.CreateTable(tab)
	db.Mysql.Where("name = ? AND pass_wd = ?", name, passwd).Find(&user)
	return &user
}

func AddUser(name string, passwd string, email string) *User {
	tab := []interface{}{&User{}}
	db.CreateTable(tab)
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
