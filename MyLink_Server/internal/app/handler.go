package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Successful Connection!")
}

func Login(c *gin.Context){

}