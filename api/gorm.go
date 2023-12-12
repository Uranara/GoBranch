package api

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB
var err error

func init() {

	// connect database
	dsn := "root:200046@tcp(127.0.0.1:3306)/webserver?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		// close transaction
		SkipDefaultTransaction: false,
		// print log
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("db.connected error", err)
	}

}
