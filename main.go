package main

import (
	"subjectInformation/config"
	"subjectInformation/middleware"
	"subjectInformation/model"
	"subjectInformation/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Config.AppMode)
	r := gin.Default()

	config.InitSession(r)
	model.InitModel()
	store := cookie.NewStore([]byte("1145411919810"))
	r.Use(sessions.Sessions("UserInfo", store))
	r.Use(middleware.Error)
	router.InitRouter(r)

	r.Run(":8080")
}
