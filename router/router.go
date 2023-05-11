package router

import (
	"subjectInformation/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// r.Use(middleware.Error)
	apiRouter := r.Group("/api")
	{
		publiccontroller := controller.PublicDatabaseController{}
		apiRouter.POST("/publiccommonDatabase", publiccontroller.Add)
	}
}
