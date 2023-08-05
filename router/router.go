package router

import (
	gin "github.com/gin-gonic/gin"
)

func Initialize(){
	//Initialize Server
	router := gin.Default()
	initializeRoutes(router)
	// Run the Server
	router.Run()
}
