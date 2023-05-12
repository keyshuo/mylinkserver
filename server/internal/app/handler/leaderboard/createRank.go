package leaderboard

import (
	"MyLink_Server/server/internal/app/handler"
	sqloperate "MyLink_Server/server/internal/app/handler/sqloperate"

	"github.com/gin-gonic/gin"
)

func CreateRank(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		handler.WriteFailed(c, "please login")
		return
	}

	account := c.Value("account")
	time := c.Query("time")
	date := c.Query("date")
	msg := "insert into ranktable value ( ?, ?, ?);"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	defer db.Close()
	errmsg = db.Exec(account, time, date)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	handler.WriteOK(c, "")
}
