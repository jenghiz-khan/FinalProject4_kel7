package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jenghiz-khan/FinalProject4_kel7/database"
	"github.com/jenghiz-khan/FinalProject4_kel7/helpers"
	"github.com/jenghiz-khan/FinalProject4_kel7/models"
	"github.com/jenghiz-khan/FinalProject4_kel7/utils/error_utils"
)

func PostCategory(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	category := models.Category{}

	if contentType == appJSON {
		c.ShouldBindJSON(&category)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	err := db.Debug().Create(&category).Error

	if err != nil {
		err := error_utils.NewBadRequest("bad request")
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":						category.ID,
		"type":    					category.Type,
		"sold_product_amount":		category.Sold_product_amount,
		"created_at":				category.CreatedAt,
	})

}

func GetCategory(c *gin.Context) {
	db := database.GetDB()
	var category []models.Category

	err := db.Find(&category).Error

	if err != nil {
		newErr := error_utils.NewBadRequest("error")
		c.JSON(newErr.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func UpdateCategory(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_ = contentType
	id := c.Param("id")

	category := models.Category{}
	categoryID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&category)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	category.ID = categoryID

	err := db.Model(&category).Where("id = ?", id).Updates(models.Category{
		Type: category.Type,
	}).First(&category).Error

	if err != nil {
		err := error_utils.NewBadRequest("Failed to edit user")
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": 					category.ID,
		"type": 				category.Type,
		"sold_product_amount":	category.Sold_product_amount,
		"updated_at": 			category.UpdatedAt,
	})

}

func DeleteCategory(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	category := models.Category{}
	id := c.Param("id")

	categoryID := uint(userData["id"].(float64))

	category.ID = categoryID

	err := db.Delete(&category, id).Error
	if err != nil {
		err := error_utils.NewUnauthorized("Unauthorized")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been  succesfully deleted",
	})
}