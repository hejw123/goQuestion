package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Setup() {
	var err error
	var connectString = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"127.0.0.1:3306",
		"cardsystem",
	)
	DB, err = gorm.Open("mysql", connectString)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		time.Sleep(10 * time.Second)
		DB, err = gorm.Open("mysql", connectString)
		if err != nil {
			panic(err.Error())
		}
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(100)
	DB.DB().SetMaxOpenConns(500)
}
