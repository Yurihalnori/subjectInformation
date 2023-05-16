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

		uniqueRouter := apiRouter.Group("/uniqueDatabase")
		{
			uniqueController := controller.UniqueController{}
			uniqueRouter.POST("", uniqueController.Add)
		}

		userRouter := apiRouter.Group("/user")
		{
			userController := controller.UserController{}
			userRouter.POST("/register", userController.CreateUser)
			userRouter.POST("/login", userController.Login)
			userRouter.DELETE("/logout", userController.Logout)
		}
		newsRouter := r.Group("/news")
		{
			newsController := controller.NewsController{}
			newsRouter.POST("/", newsController.AddNews)
		}
	}
}
