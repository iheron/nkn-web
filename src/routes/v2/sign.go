package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(rg *gin.RouterGroup) gin.HandlerFunc {
	return func(c *gin.Context) {

		rg.GET("/sign", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message":"/v1/sign"})
		})
		rg.GET("/signin", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message":"/v1/signin"})
		})
		rg.GET("/signout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message":"/v1/signout"})
		})
	}
}
