package user

import (
	"MyLink_Server/server/internal/app/handler"
	app "MyLink_Server/server/internal/app/handler/sqloperate"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var inputUser User
	inputUser.Account = c.Query("account")
	inputUser.Username = c.Query("username")
	inputUser.Password = c.Query("password")

	msg := "select count(*) from user where account = ? ;"
	db, errmsg := app.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	defer db.Close()

	//这里可能存在BUG，count语句使用string数组保存
	result, err := db.Search(inputUser.Account)

	if err == "" {
		if result[0] >= "1" {
			handler.WriteFailed(c, "account have existed")
			return
		}
	} else {
		handler.WriteFailed(c, errmsg)
		return
	}

	msg = "select count(*) from user where username = ? ;"
	errmsg = db.UpdateMysql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	result1, err := db.Search(inputUser.Username)

	if err == "" {
		if result1[0] >= "1" {
			handler.WriteFailed(c, "username have existed")
			return
		} else {
			msg = "insert into user value ( ?, ?, ?);"
			errmsg = db.UpdateMysql(msg)
			if err != "" {
				handler.WriteFailed(c, errmsg)
				return
			}
			errmsg = db.Exec(inputUser.Account, inputUser.Username, inputUser.Password)
			if err != "" {
				handler.WriteFailed(c, errmsg)
				return
			}
			handler.WriteOK(c, "")
			return
		}
	} else {
		handler.WriteFailed(c, errmsg)
		return
	}
}
