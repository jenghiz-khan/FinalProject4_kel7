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

var (
	appJSON = "application/json"
)

func RegisterUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":			User.ID,
		"full_name":    User.Fullname,
		"email":		User.Email,
		"balance":		User.Balance,
		"created_at":	User.CreatedAt,
	})

}

func LoginUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	password = User.Password

	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		errr := error_utils.NewUnauthorized("invalid email / password")
		c.JSON(errr.Status(), errr)
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		errr := error_utils.NewUnauthorized("invalid email / password")
		c.JSON(errr.Status(), errr)
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func TopupUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_ = contentType

	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	User.ID = userID

	err := db.Model(&User).Updates(models.User{
		Balance: User.Balance,
	}).First(&User).Error

	if err != nil {
		err := error_utils.NewBadRequest("invalid update")
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your balance has been successfully updated to Rp. <current Balance>",
	})

}