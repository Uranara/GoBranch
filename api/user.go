package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(30)"`
	Punches []Punches
}

type Punches struct {
	gorm.Model
	UserID uint
}

func FormPost(context *gin.Context) {

	formData := context.PostForm("punch")
	var user User
	DB.Where("name = ?", formData).Find(&user)

	punch := Punches{
		Model:  gorm.Model{},
		UserID: user.ID,
	}
	DB.Create(&punch)
	context.JSON(200, gin.H{
		"code": punch.UserID,
	})

}

func Index(context *gin.Context) {

	err := DB.Set("gorm:table_options", "ENGINE=INNODB").AutoMigrate(&Punches{}, &User{})
	if err != nil {
		return
	}

	context.HTML(200, "index.html", nil)
}

func AddUser(context *gin.Context) {
	var user User
	data := context.PostForm("username")
	user = User{
		Model:   gorm.Model{},
		Name:    data,
		Punches: nil,
	}
	DB.Create(&user)
	context.JSON(200, gin.H{
		"users": data,
	})

}
func Find(context *gin.Context) {
	var punches []Punches
	DB.Find(&punches)
	context.JSON(200, gin.H{
		"punches": punches,
	})
}
func Show(context *gin.Context) {
	var user []User
	UserId := context.Query("id")
	DB.Where("id = ?", UserId).Find(&user)

	context.JSON(200, gin.H{
		"user": user[0].Name,
	})

}
func Delete(context *gin.Context) {

	DB.Exec("drop table users")
	DB.Exec("drop table punches")

	//context.Redirect(http.StatusFound, "/")
	context.HTML(200, "index.html", nil)
}
