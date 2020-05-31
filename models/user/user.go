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
	Github    string `json:"github"`
	Views     int32  `json:"views"`
	PassWd    string `json:"-"`
}

func GetUser(name string, passwd string) *User {
	var user User
	tab := []interface{}{&User{}}
	db.CreateTable(tab)
	db.Mysql.Where("name = ? AND pass_wd = ?", name, passwd).Find(&user)
	return &user
}

func AddUser(name string, passwd string, email string) (success bool) {
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
	success = db.Insert(&user)
	return
}

func GetUserInfo(name string) interface{} {
	var user User
	result := db.Mysql.Where("name = ?", name).Find(&user).RowsAffected
	if result == 0 {
		return nil
	}
	return user
}

func init() {

}
