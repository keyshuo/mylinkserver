package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *gin.Engine
}

func NewServer() *Server {
	serv := &Server{
		server: gin.Default(),
	}
	serv.Init()
	return serv
}

func (serv *Server) Init() {
	serv.server.Use(CorsMiddleware)

	serv.server.GET("/ping", Ping)

	//serv.server.POST("/login", Login)

	serv.server.POST("/register")

}

func CorsMiddleware(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, access-control-allow-origin, Origin, X-Requested-With, Content-Type, Accept, Content-Length, Accept-Encoding, Content-Range, Content-Disposition, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Set("content-type", "application/json")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
