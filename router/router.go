package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jenghiz-khan/FinalProject4_kel7/controllers"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
		userRouter.PATCH("/topup", controllers.TopupUser)
	}

	categoryRouter := r.Group("/categories")
	{
		categoryRouter.POST("/", controllers.PostCategory)
		categoryRouter.GET("/", controllers.GetCategory)
		categoryRouter.PATCH("/:categoryID", controllers.UpdateCategory)
		categoryRouter.DELETE("/:categoryID", controllers.DeleteCategory)
	}

	productRouter := r.Group("/products")
	{
		productRouter.POST("/", controllers.PostProducts)
		productRouter.GET("/", controllers.GetProducts)
		productRouter.PUT("/:id", controllers.UpdateProduct)
		productRouter.DELETE("/:id", controllers.DeleteProduct)
	}

	transactionRouter := r.Group("/transactions")
	{
		transactionRouter.POST("/", controllers.PostTransactions)
		transactionRouter.GET("/my-transactions", controllers.GetMyTransactions)
		transactionRouter.GET("/user-transactions", controllers.GetUserTransactions)
	}

	return r
}