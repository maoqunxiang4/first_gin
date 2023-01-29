package main

import (
	"github.com/gin-gonic/gin"
	"goPro/controller"
	"goPro/middleware"
)

func collectRoutes(r *gin.Engine) *gin.Engine {
	ug := r.Group("/user")
	{
		ug.POST("/register", controller.Register())
		ug.GET("/login", controller.Login())
		ug.GET("/info", middleware.AuthMiddleware(), controller.Info)
	}
	return r
}
