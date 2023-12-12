package main

import (
	"Go/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {
	// create gin.log
	f, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	// default cross domain middleware
	r.Use(cors.Default())

	// load static directory
	r.Static("../assets", "./assets")

	// load static html templates
	r.LoadHTMLGlob("templates/*")

	// jump router
	router.InitRouter(r)

	// bind port  8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}

}
