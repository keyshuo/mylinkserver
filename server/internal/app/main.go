package app

import (
	"MyLink_Server/server/internal/app/handler"
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
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

	serv.server.GET("/ping", handler.Ping)
}

func (serv *Server) Run() {
	klog.InitFlags(nil)
	defer klog.Flush()
	flag.Set("logtostderr", "false")
	flag.Set("alsologtostderr", "false")
	flag.Parse()

	if err := serv.server.Run(":8118"); err != nil {
		klog.Error(err, "gin run failed")
		return
	}
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
