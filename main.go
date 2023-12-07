package main

import (
	"Go/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//f, _ := os.Create("./logs/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()

	r.Static("../assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	router.InitRouter(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}

}
