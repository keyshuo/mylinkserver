package community

import (
	"MyLink_Server/server/internal/app/handler"
	sqloperate "MyLink_Server/server/internal/app/handler/sqloperate"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func GetComment(c *gin.Context) {
	db, err := sqloperate.NewMySql()
	if err != nil {
		klog.Error(err)
		handler.WriteFailed(c, "server error")
	}

	defer db.Close()

}
