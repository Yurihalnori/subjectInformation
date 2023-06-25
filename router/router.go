package router

import (
	"subjectInformation/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// r.Use(middleware.Error)
	apiRouter := r.Group("/api")
	{
		userRouter := apiRouter.Group("/user")
		{
			userController := controller.UserController{}
			userRouter.POST("/register", userController.CreateUser)
			userRouter.POST("/login", userController.Login)
			userRouter.DELETE("/logout", userController.Logout)
		}

		newsRouter := apiRouter.Group("/news")
		{
			newsController := controller.NewsController{}
			newsRouter.POST("/", newsController.AddNews)
			newsRouter.GET("/", newsController.GetNews)
		}

		uniqueRouter := apiRouter.Group("/uniqueDatabase")
		{
			uniqueController := controller.UniqueController{}
			uniqueRouter.POST("", uniqueController.Add)
			uniqueRouter.PUT("/:id", uniqueController.Change)
			uniqueRouter.DELETE("/:id", uniqueController.Delete)
		}

		teamworkRouter := apiRouter.Group("/teamwork")
		{
			teamworkController := controller.TeamworkController{}
			teamworkRouter.POST("", teamworkController.Add)
			teamworkRouter.PUT("/:id", teamworkController.Change)
			teamworkRouter.DELETE("/:id", teamworkController.Delete)
		}

	}

}
