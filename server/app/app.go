package app

import (
	"fmt"
	"net/http"

	"github.com/aman-lf/sales-server/config"
	"github.com/aman-lf/sales-server/controller"
	"github.com/aman-lf/sales-server/database"
	"github.com/aman-lf/sales-server/middleware"
	"github.com/gin-gonic/gin"
)

func Start() {
	database.ConnectDB()

	gin.SetMode(config.Cfg.Mode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.JSONLoggerMiddleware())
	router.Use(middleware.CORSMiddleware())

	addAPIRoutes(router)

	ProcessFiles()

	router.Run(fmt.Sprintf("%s:%s", config.Cfg.Host, config.Cfg.Port))
}

func addAPIRoutes(router *gin.Engine) {
	router.GET("/", welcomeFunc)

	controller.SetupProductRoute(router)
	controller.SetupsaleRoute(router)
	controller.SetupNotificationRoute(router)
}

func welcomeFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Sales Dashboard server!",
	})
}
