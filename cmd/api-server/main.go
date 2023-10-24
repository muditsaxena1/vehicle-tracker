package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/vehicle-tracker/internal"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Vehicle Tracker",
		})
	})
	r.POST("/create/user", internal.CreateUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
