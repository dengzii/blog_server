package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server/models"
	"server/tools"
)

var Mysql *gorm.DB

func Init() {

	var err error
	Mysql, err = gorm.Open("mysql", "root:root@/blog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("db init failed!" + err.Error())
	}
	models.Init()
}

func Insert(value interface{}) {

	Mysql.NewRecord(value)
	Mysql.Create(value)
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
