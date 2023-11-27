package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jenghiz-khan/FinalProject4_kel7/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifToken, err := helpers.VerifToken(c)
		_ = verifToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unautenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifToken)
		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}

}