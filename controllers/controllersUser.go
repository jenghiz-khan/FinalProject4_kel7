package controllers

import (
	"fmt"
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
		err := error_utils.NewBadRequest("bad request")
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id"         :  User.ID,
		"full_name"  :  User.Fullname,
		"email"      :  User.Email,
		"balance"    :  User.Balance,
		"created_at" :  User.CreatedAt,
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

	User.ID = userID

	if err := db.First(&User).Error; err != nil {
		err := error_utils.NewNotFoundError("User not found")
		c.JSON(err.Status(), err)
		return
	}

	var input struct {
		Balance int `json:"balance"`
	}
	if contentType == appJSON {
		c.ShouldBindJSON(&input)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	// Menambahkan nilai tambahan ke Balance
	User.Balance += input.Balance

	// Simpan perubahan ke dalam database
	if err := db.Save(&User).Error; err != nil {
		err := error_utils.NewBadRequest("Maksimal balance 100000000")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Your balance has been successfully updated to Rp. %d", User.Balance),
	})

}

func GetDataUser(c *gin.Context) {
	db := database.GetDB()
	var user []models.User

	err := db.Find(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "error",
			"msg"   : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
