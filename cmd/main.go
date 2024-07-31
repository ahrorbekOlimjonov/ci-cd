package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, gin.H{
			"message": "Hello from CI/CD test project",
		})
	})

	r.GET("/hi", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, gin.H{
			"message": "Hi from CI/CD test project",
		})
	})

	log.Println("server is running in :6565 ...")

	r.Run(":6565")
}
