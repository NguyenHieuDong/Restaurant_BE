package middleware

import (
	"fmt"
	helper "golang-restaurant-management/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"err": fmt.Sprint("No authorizion header provided")})
			c.Abort()
			return
		}
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("firstName", claims.First_name)
		c.Set("lastName", claims.Last_name)
		c.Set("uid", claims.Uid)

		c.Next()

	}
}
