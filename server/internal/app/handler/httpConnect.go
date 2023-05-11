package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//connect and visit successfully
func WriteOK(c *gin.Context, msg string) {
	if msg != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

//connect and visit failed
func WriteFailed(c *gin.Context, err string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": err,
	})
}
