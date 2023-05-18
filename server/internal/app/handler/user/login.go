package user

import (
	"MyLink_Server/server/internal/app/handler"
	sqloperate "MyLink_Server/server/internal/app/handler/sqloperate"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

//POST use body
func Login(c *gin.Context) {
	var inputUser User
	inputUser.Account = c.Query("account")
	inputUser.Password = c.Query("password")

	msg := "select count(*) from user where account= ? and password=?;"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	defer db.Close()
	result, errmsg := db.Search(inputUser.Account, inputUser.Password)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	if result[0] == "1" {
		tokenString, err := GenerateToken(inputUser)
		if err != nil {
			klog.Error("Error: ", err)
			handler.WriteFailed(c, "token generate failed")
			return
		}
		handler.WriteOK(c, tokenString)
	}
}
