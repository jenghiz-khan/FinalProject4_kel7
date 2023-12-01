package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jenghiz-khan/FinalProject4_kel7/database"
	"github.com/jenghiz-khan/FinalProject4_kel7/helpers"
	"github.com/jenghiz-khan/FinalProject4_kel7/models"
	"github.com/jenghiz-khan/FinalProject4_kel7/utils/error_utils"
)

func PostTransactions(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))

	transaction := models.Transaction_History{}
	Product := models.Product{}
	User := models.User{}
	Category := models.Category{}

	if contentType == appJSON {
		c.ShouldBindJSON(&transaction)
	} else {
		newErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(newErr.Status(), newErr)
		return
	}

	transaction.UserID = userID
	Category.ID = Product.CategoryID

	if err := db.First(&Product, transaction.ProductID).Error; err != nil {
		newErr := error_utils.NewNotFoundError("product not found")
		c.AbortWithStatusJSON(newErr.Status(), newErr)
	}

	if transaction.Quantity > Product.Stock {
		err := error_utils.NewBadRequest("exceeds the product stock limit")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	if err := db.First(&User, transaction.UserID).Error; err != nil {
		newErr := error_utils.NewNotFoundError("product not found")
		c.AbortWithStatusJSON(newErr.Status(), newErr)
	}

	if User.Balance < transaction.Quantity*Product.Price {
		err := error_utils.NewBadRequest("the balance is insufficient")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	//mengurangi product product stock
	Product.Stock -= transaction.Quantity

	// Mengurangi saldo pengguna
	User.Balance -= transaction.Quantity * Product.Price

	transaction.Total_price = transaction.Quantity * Product.Price

	db.Create(&transaction)
	db.Save(&User)
	db.Save(&Product)

	db.First(&Category, Product.CategoryID)

	//menambahkan sold product dari quantity
	Category.Sold_product_amount += transaction.Quantity

	db.Save(&Category)

	c.JSON(http.StatusCreated, gin.H{
		"message":       "You have successfully purchased the product",
		"total_price":   transaction.Total_price,
		"quantity":      transaction.Quantity,
		"product_title": Product.Title,
	})
}

func GetMyTransactions(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var transaction []models.Transaction_History

	Transaction := models.Transaction_History{}
	userID := uint(userData["id"].(float64))

	Transaction.UserID = userID

	err := db.Preload("Product").Find(&transaction, Transaction).Error

	if err != nil {
		errr := error_utils.NewNotFoundError("task not found")
		c.JSON(errr.Status(), errr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func GetUserTransactions(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		err := error_utils.NewBadRequest("invalid id")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var transaction []models.Transaction_History

	Transaction := models.Transaction_History{}

	Transaction.ID = uint(id)

	err = db.Preload("Product").Preload("User").Where("id = ?", id).Find(&transaction, Transaction).Error
	if err != nil {
		errr := error_utils.NewNotFoundError("transactions not found")
		c.JSON(errr.Status(), errr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}
