package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteOK(c *gin.Context, msg string) {
	if msg != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}
