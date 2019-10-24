package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "my-aplication/sso/lib/database"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	http.ListenAndServe(":8083",r)
}