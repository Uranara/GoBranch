package api

import "github.com/gin-gonic/gin"

// home page

func Index(context *gin.Context) {

	context.HTML(200, "index.html", nil)
}

// admin login

func Login(context *gin.Context) {

	context.HTML(200, "login.html", nil)
}

// admin add user page

func AddPage(context *gin.Context) {

	context.HTML(200, "addUser.html", nil)
}

// create webserver  databases  users and punches tables

func CreateData(context *gin.Context) {

	// INNODB  "gorm:table_options", "ENGINE=INNODB"

	err := DB.Set("gorm:table_options", "ENGINE=INNODB").AutoMigrate(&Punches{}, &User{})
	if err != nil {
		return
	}

	context.HTML(200, "index.html", nil)
}

// delete webserver  databases  users and punches tables

func DeleteData(context *gin.Context) {

	DB.Exec("drop table users")
	DB.Exec("drop table punches")

	context.HTML(200, "index.html", nil)
}
