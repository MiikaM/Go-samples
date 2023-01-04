package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDb() {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_USERPASSWORD"), os.Getenv("MYSQL_CONNECTION"), os.Getenv("MYSQL_ENDPOINT"))
	fmt.Print(connectionStr)
	d, err := gorm.Open("mysql", connectionStr)

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
