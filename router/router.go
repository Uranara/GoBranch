package router

import (
	"Go/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	v1 := r.Group("")
	{
		v1.GET("/", api.Index)
		v1.GET("/get", api.Find)
		v1.GET("/delete", api.Delete)
		v1.GET("/show", api.Show)

		v1.POST("/form", api.FormPost)
		v1.POST("/add", api.AddUser)

	}

}
