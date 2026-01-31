package main

import (
	"log"

	"github.com/MouslyCode/sema/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .Env File")
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFile("/", "./web/index.html")
	router.Register(r)

	r.Run(":8080")
}
