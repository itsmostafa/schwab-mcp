package main

import (
	"simple-login/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/register", controllers.RegisterHandler)
	router.POST("/login", controllers.LoginHandler)
	router.GET("/profile", controllers.ProfileHandler)

	router.Run()
}
