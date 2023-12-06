package api

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:200046@tcp(127.0.0.1:3306)/webserver?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true,
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {

		log.Fatalln("db.connected error", err)
	}

	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})
	//if err != nil {
	//	log.Fatalln("db.connected error", err)
	//
	//}

	//
	//err = db.AutoMigrate(&api.User{})
	//if err != nil {
	//	return
	//}

	fmt.Println("连接成功")
}
