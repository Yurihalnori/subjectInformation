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
			userRouter.GET("/", userController.CheckUserStatus)
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
			uniqueRouter.GET("", uniqueController.GetInfo)
			uniqueRouter.POST("", uniqueController.Add)
			uniqueRouter.PUT("/:id", uniqueController.Change)
			uniqueRouter.DELETE("/:id", uniqueController.Delete)
		}

		teamworkRouter := apiRouter.Group("/teamwork")
		{
			teamworkController := controller.TeamworkController{}
			teamworkRouter.GET("", teamworkController.GetInfo)
			teamworkRouter.POST("", teamworkController.Add)
			teamworkRouter.PUT("/:id", teamworkController.Change)
			teamworkRouter.DELETE("/:id", teamworkController.Delete)
		}

		projectRouter := apiRouter.Group("/commonDatabase/project")
		{
			projectController := controller.ProjectController{}
			projectRouter.GET("", projectController.GetInfo)
			projectRouter.POST("", projectController.Add)
			projectRouter.PUT("/:id", projectController.Change)
			projectRouter.DELETE("/:id", projectController.Delete)
		}

		InstituteRouter := apiRouter.Group("/commonDatabase/institute")
		{
			instituteController := controller.InstituteController{}
			InstituteRouter.GET("", instituteController.GetInfo)
			InstituteRouter.POST("", instituteController.Add)
			InstituteRouter.PUT("/:id", instituteController.Change)
			InstituteRouter.DELETE("/:id", instituteController.Delete)
		}

		BookRouter := apiRouter.Group("/commonDatabase/book")
		{
			BookController := controller.BookController{}
			BookRouter.GET("", BookController.GetInfo)
			BookRouter.POST("", BookController.Add)
			BookRouter.PUT("/:id", BookController.Change)
			BookRouter.DELETE("/:id", BookController.Delete)
		}

		DissertationRouter := apiRouter.Group("/commonDatabase/dissertation")
		{
			DissertationController := controller.DissertationController{}
			DissertationRouter.GET("", DissertationController.GetInfo)
			DissertationRouter.POST("", DissertationController.Add)
			DissertationRouter.PUT("/:id", DissertationController.Change)
			DissertationRouter.DELETE("/:id", DissertationController.Delete)
		}

		ArticleRouter := apiRouter.Group("/commonDatabase/article")
		{
			ArticleController := controller.ArticleController{}
			ArticleRouter.GET("", ArticleController.GetInfo)
			ArticleRouter.POST("", ArticleController.Add)
			ArticleRouter.PUT("/:id", ArticleController.Change)
			ArticleRouter.DELETE("/:id", ArticleController.Delete)
		}

		TutorRouter := apiRouter.Group("/commonDatabase/tutor")
		{
			TutorController := controller.TutorController{}
			TutorRouter.GET("", TutorController.GetInfo)
			TutorRouter.POST("", TutorController.Add)
			TutorRouter.PUT("/:id", TutorController.Change)
			TutorRouter.DELETE("/:id", TutorController.Delete)
		}

		apiRouter.POST("/file", controller.FileController{}.Upload)
		apiRouter.PUT("/commonDatabase/search", controller.SearchController{}.SearchCommonDB)
		apiRouter.PUT("/uniqueDatabase/search", controller.SearchController{}.SearchUniqueDB)
		apiRouter.PUT("/teamwork/search", controller.SearchController{}.SearchTeamworks)
	}

}
