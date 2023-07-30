package router

import (
	"subjectInformation/controller"
	"subjectInformation/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// r.Use(middleware.Error)
	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/admin/change", middleware.LoginCheck, middleware.AdminCheck, controller.UserController{}.ChangeUserInfo)
		userRouter := apiRouter.Group("/user")
		{
			userController := controller.UserController{}
			userRouter.POST("/register", userController.CreateUser)
			userRouter.POST("/login", userController.Login)
			userRouter.DELETE("/logout", userController.Logout)
			userRouter.GET("/:id", userController.CheckUserStatus)
		}

		newsRouter := apiRouter.Group("/news")
		{
			newsController := controller.NewsController{}
			newsRouter.POST("/", middleware.LoginCheck, middleware.AdminCheck, newsController.AddNews)
			newsRouter.GET("/", newsController.GetNews)
			newsRouter.GET("/:id", middleware.LoginCheck, newsController.ApplyOneNews)
			newsRouter.PUT("/:id", middleware.LoginCheck, middleware.AdminCheck, newsController.EditOneNews)
			newsRouter.DELETE("/:id", middleware.LoginCheck, middleware.AdminCheck, newsController.DeleteOneNews)
			newsRouter.POST("/search", newsController.SearchNews)
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

		projectRouter := apiRouter.Group("/common/project")
		{
			projectController := controller.ProjectController{}
			projectRouter.POST("", projectController.Add)
			projectRouter.PUT("/:id", projectController.Change)
			projectRouter.DELETE("/:id", projectController.Delete)
		}

		InstituteRouter := apiRouter.Group("/common/institute")
		{
			instituteController := controller.InstituteController{}
			InstituteRouter.POST("", instituteController.Add)
			InstituteRouter.PUT("/:id", instituteController.Change)
			InstituteRouter.DELETE("/:id", instituteController.Delete)
		}

		BookRouter := apiRouter.Group("/common/book")
		{
			BookController := controller.BookController{}
			BookRouter.POST("", BookController.Add)
			BookRouter.PUT("/:id", BookController.Change)
			BookRouter.DELETE("/:id", BookController.Delete)
		}

		DissertationRouter := apiRouter.Group("/common/dissertation")
		{
			DissertationController := controller.DissertationController{}
			DissertationRouter.POST("", DissertationController.Add)
			DissertationRouter.PUT("/:id", DissertationController.Change)
			DissertationRouter.DELETE("/:id", DissertationController.Delete)
		}

		ArticleRouter := apiRouter.Group("/common/article")
		{
			ArticleController := controller.ArticleController{}
			ArticleRouter.POST("", ArticleController.Add)
			ArticleRouter.PUT("/:id", ArticleController.Change)
			ArticleRouter.DELETE("/:id", ArticleController.Delete)
		}

		apiRouter.PUT("/search", controller.SearchController{}.SearchCommonDB)
	}

}
