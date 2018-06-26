package main

import (
	"github.com/liangshuai/retu/system"
	"github.com/gin-gonic/gin"
	"github.com/liangshuai/retu/controllers"
)

func InitRoutes(engine *gin.Engine) {
	api := engine.Group("/api/")
	{
		api.POST("/signin", system.JWT.LoginHandler)
		// TODO Signup

		home := api.Group("/home")
		home.Use(system.JWT.MiddlewareFunc())
		{
			home.GET("", controllers.HomeGet)
		}
	}
}