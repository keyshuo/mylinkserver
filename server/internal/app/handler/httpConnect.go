package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//connect and visit successfully
func WriteOK(c *gin.Context, msg interface{}) {
	if msg != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

//connect and visit failed
func WriteFailed(c *gin.Context, err interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": err,
	})
}
