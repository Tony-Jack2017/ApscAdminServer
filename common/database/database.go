package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var dns string

func InitDatabase() {
	dns = fmt.Sprintf("%s:%d@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local")
}

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect the database failed !!!, error: %s\n", err)
		panic("program is end")
	}
}
