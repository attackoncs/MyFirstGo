package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"yl/ginessential/model"
)

var DB *gorm.DB

//初始化数据库
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "rootyl"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset) //拼接数据库的配置

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{}) //自动创建数据表
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
