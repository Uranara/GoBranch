package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(30)"`
	Punches []Punches
}

// add user  punches nil

func AddUser(context *gin.Context) {
	var user User
	data := context.PostForm("username")
	avatar, _ := context.FormFile("avatar")
	if avatar != nil {
		avatar.Filename = data + ".jpg"
		err := context.SaveUploadedFile(avatar, "./assets/avatars/"+avatar.Filename)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}
	user = User{
		Model:   gorm.Model{},
		Name:    data,
		Punches: nil,
	}
	DB.Create(&user)

	context.Redirect(http.StatusFound, "/")
}

// search user  where id = ?

func SearchUser(context *gin.Context) {
	var user []User
	UserId := context.Query("id")
	DB.Where("id = ?", UserId).Find(&user)

	context.JSON(200, gin.H{
		"user": user[0].Name,
	})

}

// ShowUser show  all users
func ShowUser(context *gin.Context) {
	var user []User
	DB.Find(&user)
	context.JSON(200, gin.H{
		"user": user,
	})

}
