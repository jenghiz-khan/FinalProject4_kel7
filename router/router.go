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
		categoryRouter.POST("/create", controllers.PostCategory)
		categoryRouter.GET("/get", controllers.GetCategory)
		categoryRouter.PATCH("/patch/:categoryId", controllers.UpdateCategory)
		categoryRouter.DELETE("/delete/:categoryId", controllers.DeleteCategory)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/create", middlewares.RoleAuthorization(), controllers.PostProducts)
		productRouter.GET("/get", controllers.GetProducts)
		productRouter.PUT("/put/:id", middlewares.RoleAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("delete/:id", middlewares.RoleAuthorization(), controllers.DeleteProduct)
	}

	transactionRouter := r.Group("/transactions")
	{
		transactionRouter.Use(middlewares.Authentication())
		transactionRouter.POST("/create", controllers.PostTransactions)
		transactionRouter.GET("/my-transactions", controllers.GetMyTransactions)
		transactionRouter.GET("/user-transactions/:id", middlewares.RoleAuthorization(), controllers.GetUserTransactions)
	}

	return r
}
