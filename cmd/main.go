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

	log.Println("server is running in :8080 ...")

	r.Run(":8080")
}
