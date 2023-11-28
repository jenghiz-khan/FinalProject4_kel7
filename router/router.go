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
		categoryRouter.Use(middlewares.Authentication())
		categoryRouter.POST("/", middlewares.RoleAuthorization(), controllers.PostCategory)
		categoryRouter.GET("/", middlewares.RoleAuthorization(), controllers.GetCategory)
		categoryRouter.PATCH("/:id", middlewares.RoleAuthorization(), controllers.UpdateCategory)
		categoryRouter.DELETE("/:id", middlewares.RoleAuthorization(), controllers.DeleteCategory)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.PostProducts)
		productRouter.GET("/", controllers.GetProducts)
		productRouter.PUT("/:id", controllers.UpdateProduct)
		productRouter.DELETE("/:id", controllers.DeleteProduct)
	}

	transactionRouter := r.Group("/transactions")
	{
		transactionRouter.Use(middlewares.Authentication())
		transactionRouter.POST("/", controllers.PostTransactions)
		transactionRouter.GET("/my-transactions", controllers.GetMyTransactions)
		transactionRouter.GET("/user-transactions", controllers.GetUserTransactions)
	}

	return r
}