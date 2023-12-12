package router

import (
	"Go/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	v1 := r.Group("")
	{
		// Home
		v1.GET("/", api.Index)
		v1.GET("/createData", api.CreateData)
		v1.GET("/deleteData", api.DeleteData)
		v1.GET("/login", api.Login)
		v1.GET("/addPage", api.AddPage)

		// User
		v1.POST("/addUser", api.AddUser)
		v1.GET("/searchUser", api.SearchUser)
		v1.GET("/showUser", api.ShowUser)

		// Punch
		v1.GET("/showPunch", api.ShowPunch)
		v1.POST("/formPost", api.FormPost)

	}

}
