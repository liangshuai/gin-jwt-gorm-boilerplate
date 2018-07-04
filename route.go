package main

import (
	"github.com/liangshuai/retu/system"
	"github.com/liangshuai/retu/model"
	"github.com/gin-gonic/gin"
	"github.com/liangshuai/retu/controllers"
)

func InitRoutes(engine *gin.Engine) {
	userAuthMiddleware := system.NewAuthMiddleware(model.UserRoleSignedIn)
	adminAuthMiddleware := system.NewAuthMiddleware(model.UserRoleAdmin)
	superAdminAuthMiddleware := system.NewAuthMiddleware(model.UserRoleSuperAdmin)
	api := engine.Group("/api/")
	{
		api.POST("/signin", userAuthMiddleware.LoginHandler)
		api.GET("/signup", controllers.Signup)

		home := api.Group("/home")
		home.Use(userAuthMiddleware.MiddlewareFunc())
		{
			home.GET("", controllers.HomeGet)
		}
		admin := api.Group("/admin")
		admin.Use(adminAuthMiddleware.MiddlewareFunc())
		{
			admin.GET("/", controllers.HomeGetAdmin)
		}
		sudo := api.Group("/sudo")
		sudo.Use(superAdminAuthMiddleware.MiddlewareFunc())
		{
			sudo.GET("/", controllers.HomeGetSudo)
		}
	}
}