package router

import (
	gin "github.com/gin-gonic/gin"
)

func Initialize(){
	printMe()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
