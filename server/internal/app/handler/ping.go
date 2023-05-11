package handler

import (
	"github.com/gin-gonic/gin"
)

//test connection
func Ping(c *gin.Context) {
	WriteOK(c, "success")
}
