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
		userRouter := apiRouter.Group("/user")
		{
			userController := controller.UserController{}
			userRouter.POST("/register", userController.CreateUser)
			userRouter.POST("/login", userController.Login)
			userRouter.DELETE("/logout", userController.Logout)
		}
	}
}
