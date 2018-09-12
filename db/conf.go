package db

import (
	// Trigger mysql.init to do the initializations
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
)

const (
	USER     = "root"
	PASSWORD = "123456"
	HOST     = "localhost:3360"
	DBNAME   = "blog"
)

type MySQL struct {
	DB *gorm.DB
}

var DBHelper *MySQL

func init() {

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD,HOST ,DBNAME))
	if err != nil {
		log.Print("connect to mysql failed!")
	}

	DBHelper = &MySQL{
		DB: db,
	}

}
