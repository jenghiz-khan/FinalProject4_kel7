package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jenghiz-khan/FinalProject4_kel7/database"
	"github.com/jenghiz-khan/FinalProject4_kel7/helpers"
	"github.com/jenghiz-khan/FinalProject4_kel7/models"
	"github.com/jenghiz-khan/FinalProject4_kel7/utils/error_utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifToken, err := helpers.VerifToken(c)
		_ = verifToken

		if err != nil {
			errr := error_utils.NewUnauthorized("unauthorized")
			c.AbortWithStatusJSON(errr.Status(), errr)
			return
		}
		c.Set("userData", verifToken)
		c.Next()
	}
}

func RoleAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			err := error_utils.NewNotFoundError("not found")
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}
		if user.Role != "admin" {
			err := error_utils.NewUnauthorized("hanya admin yang dapat akses")
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}