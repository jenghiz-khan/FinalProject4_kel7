package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jenghiz-khan/FinalProject4_kel7/database"
	"github.com/jenghiz-khan/FinalProject4_kel7/helpers"
	"github.com/jenghiz-khan/FinalProject4_kel7/models"
	"github.com/jenghiz-khan/FinalProject4_kel7/utils/error_utils"
)

func PostTransactions(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	transaction := models.Transaction_History{}

	if contentType == appJSON {
		c.ShouldBindJSON(&transaction)
	} else {
		newErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(newErr.Status(), newErr)
		return
	}
	err := db.Debug().Create(&transaction).Error

	if err != nil {
		err := error_utils.NewBadRequest("bad request")
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "You have successfully purchased the product",
		"transaction_bill": transaction,
	})
}

func GetMyTransactions(c *gin.Context) {
	db := database.GetDB()
	var transaction []models.Transaction_History

	err := db.Find(&transaction).Error

	if err != nil {
		newErr := error_utils.NewBadRequest("error")
		c.JSON(newErr.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,

	})
}

func GetUserTransactions(c *gin.Context) {
	
}