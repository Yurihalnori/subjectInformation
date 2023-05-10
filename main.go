package main

import (
	"subjectInformation/config"
	"subjectInformation/middleware"
	"subjectInformation/model"
	"subjectInformation/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Config.AppMode)
	r := gin.Default()

	config.InitSession(r)
	model.InitModel()

	r.Use(middleware.Error)
	router.InitRouter(r)

	r.Run(":8080")
}
