package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	// "github.com/claudiu/gocron"
	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/liangshuai/retu/system"
	"github.com/liangshuai/retu/model"
	// "github.com/liangshuai/retu/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	initLogger()
	system.LoadConfig()
	config := system.GetConfig()
	model.InitDB(system.GetConnectionString())
	model.AutoMigrate()
	system.Init(
		config.Realm,
		config.Secret,
		config.TTL,
	)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(system.CORSMiddleware())
	InitRoutes(engine)
	StartServer(engine)
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

// func initSessions(router *gin.Engine) {
// 	config := system.GetConfig()
// 	store := cookie.NewStore([]byte(config.SessionSecret))
// 	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7* 86400})
// 	router.Use(sessions.Sessions("gin-session", store))
// 	router.Use(controllers.ContextData())
// 	router.GET("/", controllers.HomeGet)
// 	auth := router.Group("/auth")
// 	auth.Use()
// 	auth.POST("/signin", )
// 	router.Run(":8080")
// }