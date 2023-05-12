package community

import (
	"MyLink_Server/server/internal/app/handler"
	sqloperate "MyLink_Server/server/internal/app/handler/sqloperate"
	"fmt"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func GetComment(c *gin.Context) {
	page := c.Query("page")
	var page_int int

	if _, err := fmt.Sscan(page, &page_int); err == nil {
		klog.Error(err)

	}
	page = string(rune((page_int - 1) * 50))

	msg := "select comment.*,user.username from comment inner join on comment.userid order by comment.time desc limit 50 offset ?; "
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	defer db.Close()

	result, errmsg := db.Search(page)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	handler.WriteOK(c, result)
}

func GetMyComment(c *gin.Context) {
	account := c.Query("account")
	page := c.Query("page")
	var page_int int

	if _, err := fmt.Sscan(page, &page_int); err == nil {
		klog.Error(err)

	}
	page = string(rune((page_int - 1) * 50))

	msg := "select comment.*,user.username from comment inner join on comment.userid=user.userid where comment.account = ? order by comment.time desc limit 50 offset ?; "
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	defer db.Close()

	result, errmsg := db.Search(account, page)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	handler.WriteOK(c, result)
}

func CreateComment(c *gin.Context) {
	status := c.Query("status")
	if status == "false" {
		handler.WriteFailed(c, "please login")
		return
	}
	comment := c.Query("comment")
	time := c.Query("time")
	account := c.Query("account")
	msg := "insert into comment value (?,?,?);"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	defer db.Close()
	errmsg = db.Exec(account, time, comment)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	handler.WriteOK(c, "")
}

func DeleteComment(c *gin.Context) {
	status := c.Query("status")
	if status == "false" {
		handler.WriteFailed(c, "please login")
		return
	}
	time := c.Query("time")
	account := c.Query("account")
	msg := "delete from comment where account = ? and time = ?;"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	defer db.Close()
	errmsg = db.Exec(account, time)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	handler.WriteOK(c, "")
}
