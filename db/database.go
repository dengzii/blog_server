package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server/tools"
)

var Mysql *gorm.DB

func Init() {

	var err error
	Mysql, err = gorm.Open("mysql", "root:root@/blog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("db init failed!" + err.Error())
	}
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
