package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server/conf"
	"server/tools"
)

var Mysql *gorm.DB

func Init() {

	var err error
	dbConf := conf.Get().Db
	dbUrl := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		dbConf.Password, dbConf.Username, dbConf.Database, dbConf.Charset)
	fmt.Println(dbUrl)
	Mysql, err = gorm.Open("mysql", dbUrl)

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
