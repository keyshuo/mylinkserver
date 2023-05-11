package user

import (
	"fmt"

	"MyLink_Server/server/internal/app/handler"
	app "MyLink_Server/server/internal/app/handler/sqloperate"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func Register(c *gin.Context) {
	var inputUser User
	inputUser.Account = c.Query("account")
	inputUser.Username = c.Query("username")
	inputUser.Password = c.Query("password")

	db, err := app.NewMySql()
	if err != nil {
		klog.Error(err)
		handler.WriteFailed(c, "database connection failed")
		return
	}

	defer db.Close()
	err = db.Test()
	if err != nil {
		klog.Error(err)
		handler.WriteFailed(c, "database connection failed")
		return
	}
	//这里可能存在BUG，count语句使用string数组保存
	msg := fmt.Sprintf("select count(*) from user where account = %s ;", inputUser.Account)
	if err = db.Prepare(msg); err != nil {
		klog.Error(err)
		handler.WriteFailed(c, "account is illegal")
		return
	}

	result, err := db.Search(msg)

	if err == nil {
		if result[0] == "1" {
			klog.Error("Error: ", err)
			handler.WriteFailed(c, "account have existed")
			return
		} else {
			msg = fmt.Sprintf("insert into user value ( %s, %s, %s);", inputUser.Account, inputUser.Username, inputUser.Password)
			err = db.Prepare(msg)
			if err != nil {
				klog.Error(err)
				handler.WriteFailed(c, "account or password is illegal")
				return
			}
			err = db.Exec(msg)
			if err != nil {
				klog.Error("Error: ", err)
				handler.WriteFailed(c, "register failed")
				return
			}
			handler.WriteOK(c, "")
			return
		}
	} else {
		klog.Error("Error: ", err)
		handler.WriteFailed(c, "search error,please try again")
		return
	}

}
