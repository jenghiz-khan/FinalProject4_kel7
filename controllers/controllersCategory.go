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

	err := db.Preload("Product").Find(&category).Error

	if err != nil {
		newErr := error_utils.NewBadRequest("Bad request")
		c.JSON(newErr.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func UpdateCategory(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	categoryId, err := strconv.Atoi(c.Param("categoryId"))

	if err != nil {
		err := error_utils.NewBadRequest("invalid parameter")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	Category := models.Category{}

	Category.ID = uint(categoryId)

	err = db.First(&Category, categoryId).Error

	if err != nil {
		err := error_utils.NewNotFoundError("category not found")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var updateData struct {
		Type string `json:"type"`
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&updateData)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	db.Model(&Category).Update("type", updateData.Type)
	c.JSON(http.StatusOK, gin.H{
		"id": 					Category.ID,
		"type": 				Category.Type,
		"sold_product_amount":	Category.Sold_product_amount,
		"updated_at": 			Category.UpdatedAt,
	})

}

func DeleteCategory(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	categoryId, err := strconv.Atoi(c.Param("categoryId"))

	if err != nil {
		err := error_utils.NewBadRequest("invalid parameter")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	Category := models.Category{}

	userID := uint(userData["id"].(float64))

	Category.ID = userID
	Category.ID = uint(categoryId)

	if err := db.First(&Category, categoryId).Error; err != nil {
		err := error_utils.NewNotFoundError("category not found")
		c.JSON(err.Status(), err)
		return
	}

	err = db.Delete(&Category, categoryId).Error
	if err != nil {
		err := error_utils.NewBadRequest("failed to delete category")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Category has been succesfully deleted",
	})
}