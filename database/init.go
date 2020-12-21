package database

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qt-sc/server/model"
)

var db * gorm.DB

func init(){
	var DBNAME = "mydb"
	var DBUSERNAME = "root"
	var DBPASSWORD = "root"
	var DBADDRESS = "localhost"
	var DBPORT = "3306"
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DBUSERNAME, DBPASSWORD, DBADDRESS, DBPORT, DBNAME)
	var err error
	db, err = gorm.Open("mysql", url)
	if err != nil {
		log.Println(err)
		return
	}

	db.SingularTable(true)
	createTable()

	return
}

func createTable() {

	if(!db.HasTable(&model.User{})) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
	}

	if(!db.HasTable(&model.Article{})) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Article{})
	}

	if(!db.HasTable(&model.Reply{})) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Reply{})
	}

	if(!db.HasTable(&model.Tag{})) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Tag{})
	}

	return
}
