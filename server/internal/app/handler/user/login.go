package user

import (
	"fmt"

	"MyLink_Server/server/internal/app/handler"
	sqloperate "MyLink_Server/server/internal/app/handler/sqloperate"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func Login(c *gin.Context) {
	var inputUser User
	inputUser.Account = c.Query("account")
	inputUser.Password = c.Query("password")

	db, err := sqloperate.NewMySql()
	if err != nil {
		klog.Error(err)
		handler.WriteFailed(c, "database connection failed")
		return
	}

	defer db.Close()
	msg := fmt.Sprintf("select count(*) from user where account= %s and password=%s;", inputUser.Account, inputUser.Password)
	err = db.Prepare(msg)
	if err != nil {
		klog.Error(err)
		handler.WriteFailed(c, "account or password is incorrect")
		return
	}
	result, err := db.Search(msg)
	if err != nil {
		klog.Error(err)
		handler.WriteFailed(c, "account or password is incorrect")
		return
	}
	if err == nil {
		if result[0] == "1" {
			tokenString, err := GenerateToken(inputUser)
			if err != nil {
				klog.Error("Error: ", err)
				handler.WriteFailed(c, "token generate failed")
				return
			}
			handler.WriteOK(c, tokenString)
		} else {
			klog.Error(err)
			handler.WriteFailed(c, "account or password is incorrect")
			return
		}
	}
}
