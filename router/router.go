package router

import (
	"subjectInformation/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// r.Use(middleware.Error)
	apiRouter := r.Group("/api")
	{
		projectController := controller.ProjectController{}
		apiRouter.POST("/commonDatabase/project", projectController.Add)
		// apiRouter.PUT("/commonDatabase/:id", projectController.Change)
	}
}
