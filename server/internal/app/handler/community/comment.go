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

	if _, err := fmt.Sscan(page, &page_int); err != nil {
		klog.Error(err)
	}
	page = string(rune((page_int - 1) * 50))

	msg := "select user.username,comment.date,comment.comment from comment join user on user.account=comment.account order by comment.date desc limit 50 offset ?; "
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	defer db.Close()

	result, errmsg := db.SearchRows(&UserComment{}, page)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}

	handler.WriteOK(c, result)
}

func GetMyComment(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		handler.WriteFailed(c, "please login")
		return
	}
	account := c.Value("account")
	// fmt.Println(account)
	page := c.Query("page")
	var page_int int

	if _, err := fmt.Sscan(page, &page_int); err != nil {
		klog.Error(err)
	}
	page = string(rune((page_int - 1) * 50))

	msg := "select user.username,comment.date,comment.comment from comment join user on user.account=comment.account where comment.account = ? order by comment.date desc limit 50 offset ? ;"
	db, errmsg := sqloperate.NewMySql(msg)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	defer db.Close()

	result, errmsg := db.SearchRows(&UserComment{}, account, page)
	if errmsg != "" {
		handler.WriteFailed(c, errmsg)
		return
	}
	handler.WriteOK(c, result)
}

func CreateComment(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		handler.WriteFailed(c, "please login")
		return
	}
	comment := c.Query("comment")
	time := c.Query("time")
	account := c.Value("account")
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
	status := c.Value("status")
	if status == "false" {
		handler.WriteFailed(c, "please login")
		return
	}
	time := c.Query("time")
	account := c.Value("account")
	msg := "delete from comment where account = ? and date = ?;"
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
