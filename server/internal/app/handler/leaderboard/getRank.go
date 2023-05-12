package leaderboard

import (
	"MyLink_Server/server/internal/app/handler"
	sqloperate "MyLink_Server/server/internal/app/handler/sqloperate"

	"github.com/gin-gonic/gin"
)

//low difficulty
func GetRankLow(c *gin.Context) {

	msg := "select user.username,ranklow.date,ranklow.score from ranklow join user on user.account=ranklow.account order by score desc limit 50;"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	defer db.Close()

	result, errmsg := db.SearchRows(&UserRank{})
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	handler.WriteOK(c, result)
}

//medium difficulty
func GetRankMedium(c *gin.Context) {
	msg := "select user.username,rankmedium.date,rankmedium.score from rankmedium join user on user.account=rankmedium.account order by score desc limit 50;"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	defer db.Close()

	result, errmsg := db.SearchRows(&UserRank{})
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	handler.WriteOK(c, result)
}

//high difficulty
func GetRankHigh(c *gin.Context) {
	msg := "select user.username,rankhigh.date,rankhigh.score from rankhigh join user on user.account=rankhigh.account order by score desc limit 50;"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	defer db.Close()

	result, errmsg := db.SearchRows(&UserRank{})
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	handler.WriteOK(c, result)
}
