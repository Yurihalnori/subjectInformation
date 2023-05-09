package router

import (
	"subjectInformation/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// r.Use(middleware.Error)
	apiRouter := r.Group("/api")
	{
		helloController := controller.HelloController{}
		apiRouter.GET("", helloController.Hello)
		apiRouter.GET("/time", helloController.HelloTime)

		publiccontroller := controller.PublicDatabaseController{}
		apiRouter.POST("/publiccommonDatabase", publiccontroller.Search)
	}
}
