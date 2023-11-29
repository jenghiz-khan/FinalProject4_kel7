package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jenghiz-khan/FinalProject4_kel7/controllers"
	"github.com/jenghiz-khan/FinalProject4_kel7/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
		userRouter.PATCH("/update-account", middlewares.Authentication(), controllers.TopupUser)
		userRouter.GET("/get", controllers.GetDataUser)
	}

	categoryRouter := r.Group("/categories")
	{
		categoryRouter.Use(middlewares.Authentication(), middlewares.RoleAuthorization())
		categoryRouter.POST("/", controllers.PostCategory)
		categoryRouter.GET("/", controllers.GetCategory)
		categoryRouter.PATCH("/:categoryId", controllers.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", controllers.DeleteCategory)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", middlewares.RoleAuthorization(), controllers.PostProducts)
		productRouter.GET("/", controllers.GetProducts)
		productRouter.PUT("/:id", middlewares.RoleAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:id", middlewares.RoleAuthorization(), controllers.DeleteProduct)
	}

	transactionRouter := r.Group("/transactions")
	{
		transactionRouter.Use(middlewares.Authentication())
		transactionRouter.POST("/", controllers.PostTransactions)
		transactionRouter.GET("/my-transactions", controllers.GetMyTransactions)
		transactionRouter.GET("/user-transactions", middlewares.RoleAuthorization(), controllers.GetUserTransactions)
	}

	return r
}