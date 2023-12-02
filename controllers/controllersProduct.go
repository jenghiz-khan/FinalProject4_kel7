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

func PostProducts(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	//create data
	err := db.Create(&product).Error

	if err != nil {
		errr := error_utils.NewBadRequest("failed to create product")
		c.JSON(errr.Status(), errr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          product.ID,
		"title":       product.Title,
		"price":       product.Price,
		"stock":       product.Stock,
		"category_id": product.CategoryID,
		"created_at":  product.CreatedAt,
	})
}

func GetProducts(c *gin.Context) {
	db := database.GetDB()
	var product []models.Product

	err := db.Find(&product).Error

	if err != nil {
		newErr := error_utils.NewBadRequest("error")
		c.JSON(newErr.Status(), err)
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		err := error_utils.NewBadRequest("invalid parameter")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	product := models.Product{}

	product.ID = uint(productId)

	err = db.First(&product, productId).Error

	if err != nil {
		err := error_utils.NewNotFoundError("category not found")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var updateData struct {
		Title      string `json:"title"`
		Price      int    `json:"price"`
		Stock      int    `json:"stock"`
		CategoryID uint   `json:"category_id"`
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&updateData)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	// Update dan cek apakah category_id sesuai dengan database
	check := db.Model(&product).Updates(models.Product{
		Title:      updateData.Title,
		Price:      updateData.Price,
		Stock:      updateData.Stock,
		CategoryID: updateData.CategoryID,
	}).Error
	
	if check != nil {
		errr := error_utils.NewBadRequest("failed to update product)")
		c.JSON(errr.Status(), errr)
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func DeleteProduct(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	id, _ := strconv.Atoi(c.Param("id"))

	productID := uint(userData["id"].(float64))

	product.ID = productID
	product.ID = uint(id)

	if err := db.First(&product, id).Error; err != nil {
		err := error_utils.NewNotFoundError("task not found")
		c.JSON(err.Status(), err)
		return
	}

	err := db.Delete(&product).Error
	if err != nil {
		err := error_utils.NewBadRequest("Failed to delete product")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product has been succesfully deleted",
	})
}
