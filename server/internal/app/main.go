package app

import (
	"MyLink_Server/server/internal/app/handler"
	community "MyLink_Server/server/internal/app/handler/community"
	"MyLink_Server/server/internal/app/handler/leaderboard"
	usr "MyLink_Server/server/internal/app/handler/user"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	serv.server.POST("/login", usr.Login)

	serv.server.POST("/register", usr.Register)

	comment := serv.server.Group("/comment")
	{

		comment.GET("/getComment", community.GetComment)

		my := comment.Group("/my")
		{
			my.Use(AuthMiddleware)

			my.GET("/getMyComment", community.GetMyComment)

			my.GET("/createComment", community.CreateComment)

			my.GET("/deleteComment", community.DeleteComment)
		}
	}

	ranking := serv.server.Group("/rank")
	{

		my := ranking.Group("/my")
		{
			my.Use(AuthMiddleware)

			my.GET("getRankLow", leaderboard.GetRankLow)

			my.GET("getRankMedium", leaderboard.GetRankMedium)

			my.GET("getRankHigh", leaderboard.GetRankHigh)

			my.GET("createRank", leaderboard.CreateRank)

		}

	}

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
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, access-control-allow-origin, Origin, X-Requested-With, Content-Type, Accept, Content-Length, Accept-Encoding, Content-Range, Content-Disposition, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Set("content-type", "application/json")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}

func AuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		fmt.Print(tokenString)
		c.Status(http.StatusUnauthorized)
		return
	}
	if len(tokenString) == 0 {
		fmt.Print(tokenString)
		c.Status(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte("my_secret_key"), nil //后面需要改为文件读取本地的公钥
	})
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	if !token.Valid {
		c.Status(http.StatusUnauthorized)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		klog.Error("Get user's message occur error!")
	}
	if expire, ok := claims["expire"].(float64); ok {
		if time.Unix(int64(expire), 0).Before(time.Now()) {
			c.Status(http.StatusUnauthorized)
			return
		}
	}
	c.Set("account", claims["account"].(string))
	c.Set("status", claims["status"].(string))
	c.Set("username", claims["username"].(string))
	c.Next()
}

func ExceptionHandlerMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			klog.Error(r)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "somthing went wrong",
			})
		}
	}()
}
