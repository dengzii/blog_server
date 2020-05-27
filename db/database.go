package db

import (
	"fmt"
	"github.com/dengzii/blog_server/conf"
	"github.com/dengzii/blog_server/tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Mysql *gorm.DB

func Init() {

	var err error
	dbConf := conf.Get().Db
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database, dbConf.Charset)
	fmt.Println(dbUrl)
	Mysql, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		panic("db init failed!" + err.Error())
	}
	Mysql.LogMode(dbConf.LogEnable)
}

func Insert(value interface{}) (success bool) {

	success = Mysql.NewRecord(value)
	Mysql.Create(value)
	return
}

func CreateTable(tables []interface{}) {

	for _, table := range tables {
		if Mysql.HasTable(table) {
			continue
		}
		tools.Log("Create table ", table)
		Mysql.CreateTable(table)
	}
}
