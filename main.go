package main

import (
	"subjectInformation/config"
	"subjectInformation/middleware"
	"subjectInformation/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Config.AppMode)
	r := gin.Default()

	config.InitSession(r)

	r.Use(middleware.Error)
	router.InitRouter(r)

	r.Run(":8080")
}
